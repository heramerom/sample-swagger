package sample_swagger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	typeString = "string"
	typeInt    = "integer"
	typeNumber = "number"
	typeObject = "object"
	typeArray  = "array"
	typeMap    = "map"
)

var buildInTypes = map[string]string{
	"time.Time":  "string",
	"*time.Time": "string",
}

var definitions = make(map[string]model)

func parse() string {
	var swagger Swagger
	err := json.Unmarshal([]byte(generatorJson), &swagger)
	if err != nil {
		fmt.Printf("unmarshal error: %s\n", err.Error())
		return ""
	}

	for _, v := range generatorModels {
		rt := reflect.TypeOf(v)
		rv := reflect.ValueOf(v)
		parseDefines(&rv, rt)
	}

	for _, v := range definitions {
		if swagger.Definitions == nil {
			swagger.Definitions = make(map[string]*Definition)
		}
		name, definition := v.toDefinition(false)
		if definition != nil && name != "" {
			swagger.Definitions[name] = definition
		}
	}
	if swagger.Swagger == "" {
		swagger.Swagger = "2.0"
	}
	bs, err := json.Marshal(swagger)
	if err != nil {
		fmt.Printf("marshal error: %s\n", err.Error())
		return ""
	}
	return string(bs)
}

type model struct {
	Name   string
	Type   string
	Object *model   `json:"object"`
	Fields []*model `json:",omitempty"` // properties

	Anonymous bool
}

func (m *model) expandFields() []*model {
	var fds []*model
	for _, f := range m.Fields {
		if f.Anonymous && f.Object != nil {
			pm, ok := definitions[f.Object.Name]
			if ok {
				fds = append(fds, pm.expandFields()...)
			}
		} else {
			fds = append(fds, f)
		}
	}
	return fds
}

func (m *model) toDefinition(ref bool) (name string, definition *Definition) {

	if m == nil {
		return
	}

	var d Definition
	name = m.Name
	d.Type = m.Type

	switch m.Type {
	case typeObject:
		if ref {
			if m.Object != nil {
				if isBaseDefinitions(m.Object.Type) {
					return m.Name, &Definition{Type: m.Object.Type}
				}
				if !isNestedObject(m.Object.Name) {
					return m.Name, &Definition{Type: typeObject, Ref: "#/definitions/" + m.Object.Name}
				}
			} else {
				if isBaseDefinitions(m.Type) {
					return m.Name, &Definition{Type: m.Type}
				}
				if !isNestedObject(m.Name) {
					return m.Name, &Definition{Type: typeObject, Ref: "#/definitions/" + m.Name}
				}
			}
		}

		if m.Object != nil {
			_, d := m.Object.toDefinition(true)
			return m.Name, d
		} else {
			if len(m.Fields) > 0 {
				ps := make(map[string]*Definition, len(m.Fields))
				fds := m.expandFields()
				for _, v := range fds {
					_, ps[v.Name] = v.toDefinition(true)
				}
				d.Properties = ps
			}
		}

	case typeArray:
		if m.Object != nil {
			_, d := m.Object.toDefinition(true)
			return m.Name, &Definition{Type: typeArray, Items: d}
		}
	case typeMap:
		if m.Object != nil {
			_, d := m.Object.toDefinition(true)
			return m.Name, &Definition{Type: typeObject, AdditionalProperties: d}
		}
	}
	definition = &d
	return
}

func isBaseDefinitions(typ string) bool {
	switch typ {
	case typeObject, typeArray, typeMap:
		return false
	}
	return true
}

func isNestedObject(name string) bool {
	return strings.Contains(name, "struct {")
}

func parseField(value reflect.Value, typ reflect.Type, fd reflect.StructField) *model {
	if fd.Name[0] > 'Z' || fd.Name[0] < 'A' {
		return nil
	}
	var f model
	f.Name = strings.Split(fd.Tag.Get("json"), ",")[0]
	if f.Name == "-" {
		return nil
	}
	f.Anonymous = fd.Anonymous
	if f.Name == "" {
		f.Name = fd.Name
	}
	switch typ.Kind() {
	case reflect.String:
		f.Type = typeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f.Type = typeInt
	case reflect.Float32, reflect.Float64:
		f.Type = typeNumber
	case reflect.Struct:
		if buildIn, ok := buildInTypes[typ.String()]; ok {
			f.Type = buildIn
			break
		}
		f.Type = typeObject
		m := parseDefines(&value, typ)
		f.Object = &m
	case reflect.Ptr:
		v, t := indirectType(fd.Type)
		return parseField(v, t, fd)
	case reflect.Array, reflect.Slice:
		f.Type = typeArray
		fmt.Println("name-->", fd.Type.Elem())
		v, t := indirectType(fd.Type.Elem())
		fmt.Println("name-->", t)
		m := parseDefines(&v, t)
		f.Object = &m
	case reflect.Map:
		f.Type = typeMap
		v, t := indirectType(fd.Type.Elem())
		fmt.Println("typ:", t)
		vm := parseDefines(&v, t)
		f.Object = &vm
	}
	return &f
}

func nameOfType(t reflect.Type) string {
	return t.String()
}

func indirectType(t reflect.Type) (reflect.Value, reflect.Type) {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.Indirect(reflect.New(t.Elem())), t.Elem()
	}
	return reflect.Indirect(reflect.New(t)), t
}

func parseDefines(v *reflect.Value, t reflect.Type) model {
	if v == nil {
		return model{}
	}

	switch t.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			v, t := indirectType(t)
			return parseDefines(&v, t)
		}
	}

	if t.Kind() == reflect.Ptr {
		obj := reflect.Indirect(*v).Interface()
		v := reflect.ValueOf(obj)
		t := reflect.TypeOf(obj)
		return parseDefines(&v, t)
	}

	key := nameOfType(t)
	if v, ok := definitions[key]; ok {
		if v.Type == typeObject && !strings.Contains(v.Name, "struct { ") {
			return model{Name: v.Name, Type: v.Type}
		}
		return v
	}
	// block dead loop
	definitions[key] = model{Name: key, Type: typeObject}

	var m model
	switch t.Kind() {
	case reflect.String:
		m.Name = typeString
		m.Type = typeString
		return m
	case reflect.Int:
		m.Name = typeInt
		m.Type = typeInt
		return m
	case reflect.Struct:
		m.Name = key
		m.Type = typeObject
		var fields []*model
		for i := 0; i < v.NumField(); i++ {
			v := v.Field(i)
			f := parseField(v, t.Field(i).Type, t.Field(i))
			if f == nil {
				continue
			}
			fields = append(fields, f)
		}
		m.Fields = fields
	case reflect.Array, reflect.Slice:
		m.Type = typeArray
		fmt.Println("name->", t.Elem().Name())
		v, t := indirectType(t.Elem())
		mm := parseDefines(&v, t)
		m.Object = &mm
	case reflect.Map:
		m.Type = typeMap
		v, t := indirectType(t.Elem())
		fmt.Println("typ:", t)
		mm := parseDefines(&v, t)
		m.Object = &mm
	}
	definitions[key] = m
	if m.Type == typeObject && !isNestedObject(m.Name) {
		return model{Name: m.Name, Type: m.Type}
	}
	return m
}
