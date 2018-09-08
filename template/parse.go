package template

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
		name, definition := v.toDefinition()
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
	Name        string
	Type        string
	Description string  `json:",omitempty"`
	ItemObject  *model  `json:",omitempty"`
	ValueObject *model  `json:",omitempty"`
	Fields      []field `json:",omitempty"`
}

func (m *model) expandFields() []field {
	var fds []field
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

func (m *model) toDefinition() (name string, definition *Definition) {

	if isNestedObject(m.Name) {
		return
	}
	name = m.Name

	var d Definition
	d.Type = m.Type
	if len(m.Fields) > 0 {
		ps := make(map[string]Property, len(m.Fields))
		fds := m.expandFields()
		for _, v := range fds {
			//p := Property{}
			//p.Type = v.Type
			//
			//if v.Type == typeObject {
			//	if v.Object != nil {
			//		if isNestedObject(v.Object.Name) {
			//			// nested object
			//		} else {
			//			p.Ref = "#/definitions/" + v.Object.Name
			//		}
			//	}
			//}
			//
			//if v.Type == typeArray {
			//	//ref := "#/definitions/" + v.ItemObject.Name
			//	//if v.ItemObject.Name == "" {
			//	//	ref = ""
			//	//}
			//	//p.Items = &Property{
			//	//	Type: v.ItemObject.Type,
			//	//	Ref:  ref,
			//	//}
			//}
			//
			//if v.Type == typeMap {
			//	p.Type = typeObject
			//	//p.Properties = &NestedProperty{
			//	//	Id:   v.KeyObject.Type,
			//	//	Name: v.ValueObject.Type,
			//	//}
			//}
			ps[v.Name] = *v.toProperty()
		}
		d.Properties = ps
	}
	definition = &d

	return
}

func isNestedObject(name string) bool {
	return strings.Contains(name, "struct {")
}

type field struct {
	Name        string
	Type        string
	Desc        string `json:",omitempty"`
	ItemObject  *model `json:",omitempty"`
	KeyObject   *model `json:",omitempty"`
	ValueObject *model `json:",omitempty"`
	Object      *model `json:",omitempty"`

	Anonymous bool
}

func (f *field) toProperty() *Property {
	var p Property
	p.Type = f.Type

	switch f.Type {

	case typeObject:
		if f.Object != nil {
			if isNestedObject(f.Object.Name) {
				// nested object
				_, nested := f.Object.toDefinition()
				if nested != nil {
					_, p.Properties = f.Object.toDefinition()
				}
			} else {
				p.Ref = "#/definitions/" + f.Object.Name
			}
		}
	case typeArray:
		if f.ItemObject != nil {
			if isNestedObject(f.ItemObject.Name) {
				_, p.Items = f.ItemObject.toDefinition()
			} else {
				p.Items.Ref = "#/definitions/" + f.ItemObject.Name
			}
		}

	case typeMap:

	}

	return &p
}

func parseField(value reflect.Value, typ reflect.Type, fd reflect.StructField) *field {
	if fd.Name[0] > 'Z' || fd.Name[0] < 'A' {
		return nil
	}
	var f field
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
		v, t := indirectType(fd.Type)
		m := parseDefines(&v, t)
		f.ItemObject = &m
	case reflect.Map:
		f.Type = typeMap
		v, t := indirectType(typ.Key())
		km := parseDefines(&v, t)
		f.KeyObject = &km
		v, t = indirectType(typ.Elem())
		vm := parseDefines(&v, t)
		f.ValueObject = &vm
	}
	return &f
}

func nameOfType(t reflect.Type) string {
	return t.String()
}

func indirectType(t reflect.Type) (reflect.Value, reflect.Type) {
	switch t.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Slice, reflect.Map:
		return reflect.Indirect(reflect.New(t.Elem())), t.Elem()
	}
	return reflect.Indirect(reflect.New(t)), t
}

func parseDefines(v *reflect.Value, t reflect.Type) model {
	if v == nil {
		return model{}
	}
	switch t.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Array:
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
		var fields []field
		for i := 0; i < v.NumField(); i++ {
			v := v.Field(i)
			f := parseField(v, t.Field(i).Type, t.Field(i))
			if f == nil {
				continue
			}
			fields = append(fields, *f)
		}
		m.Fields = fields
	case reflect.Array, reflect.Slice:
		m.Type = typeArray
		v, t := indirectType(t.Elem())
		mm := parseDefines(&v, t)
		m.ItemObject = &mm
	case reflect.Map:
		m.Type = typeMap
		v, t := indirectType(t.Elem())
		mm := parseDefines(&v, t)
		m.ValueObject = &mm
	}
	definitions[key] = m
	if m.Type == typeObject && !strings.Contains(m.Name, "struct { ") {
		return model{Name: m.Name, Type: m.Type}
	}
	return m
}
