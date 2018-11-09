package main

var templateModel = "// +build sample_swagger\n" +
	"\n" +
	"package sample_swagger\n" +
	"\n" +
	"type Info struct {\n" +
	"	Description    string `json:\"description\"`\n" +
	"	Version        string `json:\"version\"`\n" +
	"	Title          string `json:\"title\"`\n" +
	"	TermsOfService string `json:\"termsOfService\"`\n" +
	"	Contact        struct {\n" +
	"		Email string `json:\"email\"`\n" +
	"	} `json:\"contact\"`\n" +
	"	License struct {\n" +
	"		Name string `json:\"name\"`\n" +
	"		URL  string `json:\"url\"`\n" +
	"	} `json:\"license\"`\n" +
	"}\n" +
	"\n" +
	"type Router struct {\n" +
	"	Tags        []string            `json:\"tags\"`\n" +
	"	Summary     string              `json:\"summary\"`\n" +
	"	Description string              `json:\"description\"`\n" +
	"	OperationID string              `json:\"operationId\"`\n" +
	"	Consumes    []string            `json:\"consumes\"`\n" +
	"	Produces    []string            `json:\"produces\"`\n" +
	"	Parameters  []Parameter         `json:\"parameters\"`\n" +
	"	Responses   map[string]Response `json:\"responses\"`\n" +
	"}\n" +
	"\n" +
	"type Schema struct {\n" +
	"	Ref string `json:\"$ref,omitempty\"`\n" +
	"}\n" +
	"\n" +
	"type Parameter struct {\n" +
	"	In          string  `json:\"in\"`\n" +
	"	Name        string  `json:\"name\"`\n" +
	"	Type        string  `json:\"type\"`\n" +
	"	Description string  `json:\"description\"`\n" +
	"	Required    bool    `json:\"required\"`\n" +
	"	Schema      *Schema `json:\"schema,omitempty\"`\n" +
	"}\n" +
	"\n" +
	"type Response struct {\n" +
	"	Description string `json:\"description\"`\n" +
	"	Schema      struct {\n" +
	"		Type  string `json:\"type\"`\n" +
	"		Items struct {\n" +
	"			Ref string `json:\"$ref\"`\n" +
	"		} `json:\"items\"`\n" +
	"		Ref string `json:\"$ref\"`\n" +
	"	} `json:\"schema\"`\n" +
	"}\n" +
	"\n" +
	"type Property struct {\n" +
	"	Type        string      `json:\"type,omitempty\"`\n" +
	"	Format      string      `json:\"format,omitempty\"`\n" +
	"	Description string      `json:\"description,omitempty\"`\n" +
	"	Ref         string      `json:\"$ref,omitempty\"`\n" +
	"	Items       *Definition `json:\"items,omitempty\"`\n" +
	"	Properties  *Definition `json:\"properties,omitempty\"`\n" +
	"\n" +
	"	AdditionalProperties *AdditionalProperties `json:\"additionalProperties,omitempty\"`\n" +
	"}\n" +
	"\n" +
	"type AdditionalProperties struct {\n" +
	"	Type string `json:\"type\"`\n" +
	"	Ref  string `json:\"$ref\"`\n" +
	"}\n" +
	"\n" +
	"type NestedProperty struct {\n" +
	"	Id   string `json:\"id,omitempty\"`\n" +
	"	Name string `json:\"name,omitempty\"`\n" +
	"}\n" +
	"\n" +
	"type Definition struct {\n" +
	"	Type                 string                 `json:\"type,omitempty\"`\n" +
	"	Format               string                 `json:\"format,omitempty\"`\n" +
	"	Items                *Definition            `json:\"items,omitempty\"`\n" +
	"	Properties           map[string]*Definition `json:\"properties,omitempty\"`\n" +
	"	AdditionalProperties *Definition            `json:\"additionalProperties,omitempty\"`\n" +
	"	Ref                  string                 `json:\"$ref,omitempty\"`\n" +
	"}\n" +
	"\n" +
	"type Swagger struct {\n" +
	"	Swagger     string                 `json:\"swagger\"`\n" +
	"	Info        *Info                  `json:\"info\"`\n" +
	"	Host        string                 `json:\"host\"`\n" +
	"	BasePath    string                 `json:\"basePath\"`\n" +
	"	Schemes     []string               `json:\"schemes\"`\n" +
	"	Paths       map[string]Method      `json:\"paths\"`\n" +
	"	Definitions map[string]*Definition `json:\"definitions\"`\n" +
	"}\n" +
	"\n" +
	"type Method map[string]Router\n" +
	"\n" +
	"func MapType(typ string) string {\n" +
	"	switch typ {\n" +
	"	case \"int\", \"int8\", \"int16\", \"int32\", \"int64\", \"uint\", \"uint8\", \"uint16\", \"uint32\", \"uint64\":\n" +
	"		return \"integer\"\n" +
	"	case \"string\", \"str\", \"s\":\n" +
	"		return \"string\"\n" +
	"	case \"bool\", \"boolean\", \"b\":\n" +
	"		return \"boolean\"\n" +
	"	case \"object\", \"obj\", \"o\":\n" +
	"		return \"object\"\n" +
	"	case \"float32\", \"float64\":\n" +
	"		return \"number\"\n" +
	"	case \"array\", \"slice\":\n" +
	"		return \"array\"\n" +
	"	case \"map\":\n" +
	"		return \"map\"\n" +
	"	}\n" +
	"	return \"{}\" // any\n" +
	"}"

var templateParser = "// +build sample_swagger\n" +
	"\n" +
	"package sample_swagger\n" +
	"\n" +
	"import (\n" +
	"	\"encoding/json\"\n" +
	"	\"fmt\"\n" +
	"	\"reflect\"\n" +
	"	\"strings\"\n" +
	")\n" +
	"\n" +
	"const (\n" +
	"	typeString = \"string\"\n" +
	"	typeInt    = \"integer\"\n" +
	"	typeBool   = \"boolean\"\n" +
	"	typeNumber = \"number\"\n" +
	"	typeObject = \"object\"\n" +
	"	typeArray  = \"array\"\n" +
	"	typeMap    = \"map\"\n" +
	")\n" +
	"\n" +
	"var buildInTypes = map[string]string{\n" +
	"	\"time.Time\":  \"string\",\n" +
	"	\"*time.Time\": \"string\",\n" +
	"}\n" +
	"\n" +
	"var definitions = make(map[string]model)\n" +
	"\n" +
	"func parse() string {\n" +
	"	var swagger Swagger\n" +
	"	err := json.Unmarshal([]byte(generatorJson), &swagger)\n" +
	"	if err != nil {\n" +
	"		fmt.Printf(\"unmarshal error: %s\", err.Error())\n" +
	"		return \"\"\n" +
	"	}\n" +
	"\n" +
	"	for _, v := range generatorModels {\n" +
	"		rt := reflect.TypeOf(v)\n" +
	"		rv := reflect.ValueOf(v)\n" +
	"		parseDefines(&rv, rt)\n" +
	"	}\n" +
	"\n" +
	"	for _, v := range definitions {\n" +
	"		if swagger.Definitions == nil {\n" +
	"			swagger.Definitions = make(map[string]*Definition)\n" +
	"		}\n" +
	"		name, definition := v.toDefinition(false)\n" +
	"		if definition != nil && name != \"\" {\n" +
	"			swagger.Definitions[name] = definition\n" +
	"		}\n" +
	"	}\n" +
	"	if swagger.Swagger == \"\" {\n" +
	"		swagger.Swagger = \"2.0\"\n" +
	"	}\n" +
	"	bs, err := json.Marshal(swagger)\n" +
	"	if err != nil {\n" +
	"		fmt.Printf(\"marshal error: %s\", err.Error())\n" +
	"		return \"\"\n" +
	"	}\n" +
	"	return string(bs)\n" +
	"}\n" +
	"\n" +
	"type model struct {\n" +
	"	Name   string\n" +
	"	Type   string\n" +
	"	Object *model   `json:\"object\"`\n" +
	"	Fields []*model `json:\",omitempty\"` // properties\n" +
	"\n" +
	"	Anonymous bool\n" +
	"}\n" +
	"\n" +
	"func (m *model) expandFields() []*model {\n" +
	"	var fds []*model\n" +
	"	for _, f := range m.Fields {\n" +
	"		if f.Anonymous && f.Object != nil {\n" +
	"			pm, ok := definitions[f.Object.Name]\n" +
	"			if ok {\n" +
	"				fds = append(fds, pm.expandFields()...)\n" +
	"			}\n" +
	"		} else {\n" +
	"			fds = append(fds, f)\n" +
	"		}\n" +
	"	}\n" +
	"	return fds\n" +
	"}\n" +
	"\n" +
	"func (m *model) toDefinition(ref bool) (name string, definition *Definition) {\n" +
	"\n" +
	"	if m == nil {\n" +
	"		return\n" +
	"	}\n" +
	"	if !ref && isBaseDefinitions(m.Type) {\n" +
	"		return m.Type, nil\n" +
	"	}\n" +
	"\n" +
	"	var d Definition\n" +
	"	name = m.Name\n" +
	"	d.Type = m.Type\n" +
	"\n" +
	"	switch m.Type {\n" +
	"	case typeObject:\n" +
	"		if ref {\n" +
	"			if m.Object != nil {\n" +
	"				if isBaseDefinitions(m.Object.Type) {\n" +
	"					return m.Name, &Definition{Type: m.Object.Type}\n" +
	"				}\n" +
	"				if !isNestedObject(m.Object.Name) {\n" +
	"					return m.Name, &Definition{Type: typeObject, Ref: \"#/definitions/\" + m.Object.Name}\n" +
	"				}\n" +
	"			} else {\n" +
	"				if isBaseDefinitions(m.Type) {\n" +
	"					return m.Name, &Definition{Type: m.Type}\n" +
	"				}\n" +
	"				if !isNestedObject(m.Name) {\n" +
	"					return m.Name, &Definition{Type: typeObject, Ref: \"#/definitions/\" + m.Name}\n" +
	"				}\n" +
	"			}\n" +
	"		}\n" +
	"		if !ref && isNestedObject(name) {\n" +
	"			return m.Name, nil\n" +
	"		}\n" +
	"\n" +
	"		if m.Object != nil {\n" +
	"			_, d := m.Object.toDefinition(true)\n" +
	"			return m.Name, d\n" +
	"		} else {\n" +
	"			if len(m.Fields) > 0 {\n" +
	"				ps := make(map[string]*Definition, len(m.Fields))\n" +
	"				fds := m.expandFields()\n" +
	"				for _, v := range fds {\n" +
	"					_, ps[v.Name] = v.toDefinition(true)\n" +
	"				}\n" +
	"				d.Properties = ps\n" +
	"			}\n" +
	"		}\n" +
	"	case typeArray:\n" +
	"		if m.Object != nil {\n" +
	"			_, d := m.Object.toDefinition(true)\n" +
	"			return m.Name, &Definition{Type: typeArray, Items: d}\n" +
	"		}\n" +
	"	case typeMap:\n" +
	"		if m.Object != nil {\n" +
	"			_, d := m.Object.toDefinition(true)\n" +
	"			return m.Name, &Definition{Type: typeObject, AdditionalProperties: d}\n" +
	"		}\n" +
	"	}\n" +
	"	definition = &d\n" +
	"	return\n" +
	"}\n" +
	"\n" +
	"func isBaseDefinitions(typ string) bool {\n" +
	"	switch typ {\n" +
	"	case typeObject, typeArray, typeMap:\n" +
	"		return false\n" +
	"	}\n" +
	"	return true\n" +
	"}\n" +
	"\n" +
	"func isNestedObject(name string) bool {\n" +
	"	return strings.Contains(name, \"struct {\")\n" +
	"}\n" +
	"\n" +
	"func parseField(value reflect.Value, typ reflect.Type, fd reflect.StructField) *model {\n" +
	"	// unexport field\n" +
	"	if fd.Name[0] > 'Z' || fd.Name[0] < 'A' {\n" +
	"		return nil\n" +
	"	}\n" +
	"	var f model\n" +
	"	f.Name = strings.Split(fd.Tag.Get(\"json\"), \",\")[0]\n" +
	"	// ignore field\n" +
	"	if f.Name == \"-\" {\n" +
	"		return nil\n" +
	"	}\n" +
	"	f.Anonymous = fd.Anonymous\n" +
	"	if f.Name == \"\" {\n" +
	"		f.Name = fd.Name\n" +
	"	}\n" +
	"	switch typ.Kind() {\n" +
	"	case reflect.String:\n" +
	"		f.Type = typeString\n" +
	"	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,\n" +
	"		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:\n" +
	"		f.Type = typeInt\n" +
	"	case reflect.Float32, reflect.Float64:\n" +
	"		f.Type = typeNumber\n" +
	"	case reflect.Bool:\n" +
	"		f.Type = typeBool\n" +
	"	case reflect.Struct:\n" +
	"		if buildIn, ok := buildInTypes[typ.String()]; ok {\n" +
	"			f.Type = buildIn\n" +
	"			break\n" +
	"		}\n" +
	"		f.Type = typeObject\n" +
	"		m := parseDefines(&value, typ)\n" +
	"		f.Object = &m\n" +
	"	case reflect.Ptr:\n" +
	"		v, t := indirectType(fd.Type)\n" +
	"		return parseField(v, t, fd)\n" +
	"	case reflect.Array, reflect.Slice:\n" +
	"		f.Type = typeArray\n" +
	"		v, t := indirectType(fd.Type.Elem())\n" +
	"		m := parseDefines(&v, t)\n" +
	"		f.Object = &m\n" +
	"	case reflect.Map:\n" +
	"		f.Type = typeMap\n" +
	"		v, t := indirectType(fd.Type.Elem())\n" +
	"		vm := parseDefines(&v, t)\n" +
	"		f.Object = &vm\n" +
	"	}\n" +
	"	return &f\n" +
	"}\n" +
	"\n" +
	"func nameOfType(t reflect.Type) string {\n" +
	"	return t.String()\n" +
	"}\n" +
	"\n" +
	"func indirectType(t reflect.Type) (reflect.Value, reflect.Type) {\n" +
	"	switch t.Kind() {\n" +
	"	case reflect.Ptr:\n" +
	"		return reflect.Indirect(reflect.New(t.Elem())), t.Elem()\n" +
	"	}\n" +
	"	return reflect.Indirect(reflect.New(t)), t\n" +
	"}\n" +
	"\n" +
	"func parseDefines(v *reflect.Value, t reflect.Type) model {\n" +
	"	if v == nil {\n" +
	"		return model{}\n" +
	"	}\n" +
	"\n" +
	"	switch t.Kind() {\n" +
	"	case reflect.Ptr:\n" +
	"		if v.IsNil() {\n" +
	"			v, t := indirectType(t)\n" +
	"			return parseDefines(&v, t)\n" +
	"		}\n" +
	"	}\n" +
	"\n" +
	"	if t.Kind() == reflect.Ptr {\n" +
	"		obj := reflect.Indirect(*v).Interface()\n" +
	"		v := reflect.ValueOf(obj)\n" +
	"		t := reflect.TypeOf(obj)\n" +
	"		return parseDefines(&v, t)\n" +
	"	}\n" +
	"\n" +
	"	key := nameOfType(t)\n" +
	"	if v, ok := definitions[key]; ok {\n" +
	"		if v.Type == typeObject && !strings.Contains(v.Name, \"struct { \") {\n" +
	"			return model{Name: v.Name, Type: v.Type}\n" +
	"		}\n" +
	"		return v\n" +
	"	}\n" +
	"	// block dead loop\n" +
	"	definitions[key] = sampleModel(key, t)\n" +
	"\n" +
	"	var m model\n" +
	"	switch t.Kind() {\n" +
	"	case reflect.String:\n" +
	"		m.Name = typeString\n" +
	"		m.Type = typeString\n" +
	"		return m\n" +
	"	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,\n" +
	"		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:\n" +
	"		m.Name = typeInt\n" +
	"		m.Type = typeInt\n" +
	"		return m\n" +
	"	case reflect.Float32, reflect.Float64:\n" +
	"		m.Name = typeNumber\n" +
	"		m.Type = typeNumber\n" +
	"	case reflect.Bool:\n" +
	"		m.Name = typeBool\n" +
	"		m.Type = typeBool\n" +
	"	case reflect.Struct:\n" +
	"		m.Name = key\n" +
	"		m.Type = typeObject\n" +
	"		var fields []*model\n" +
	"		for i := 0; i < v.NumField(); i++ {\n" +
	"			v := v.Field(i)\n" +
	"			f := parseField(v, t.Field(i).Type, t.Field(i))\n" +
	"			if f == nil {\n" +
	"				continue\n" +
	"			}\n" +
	"			fields = append(fields, f)\n" +
	"		}\n" +
	"		m.Fields = fields\n" +
	"	case reflect.Array, reflect.Slice:\n" +
	"		m.Type = typeArray\n" +
	"		fmt.Println(\"name->\", t.Elem().Name())\n" +
	"		v, t := indirectType(t.Elem())\n" +
	"		mm := parseDefines(&v, t)\n" +
	"		m.Object = &mm\n" +
	"	case reflect.Map:\n" +
	"		m.Type = typeMap\n" +
	"		v, t := indirectType(t.Elem())\n" +
	"		fmt.Println(\"typ:\", t)\n" +
	"		mm := parseDefines(&v, t)\n" +
	"		m.Object = &mm\n" +
	"	}\n" +
	"	definitions[key] = m\n" +
	"	if m.Type == typeObject && !isNestedObject(m.Name) {\n" +
	"		return model{Name: m.Name, Type: m.Type}\n" +
	"	}\n" +
	"	return m\n" +
	"}\n" +
	"\n" +
	"func sampleModel(key string, t reflect.Type) model {\n" +
	"	switch t.Kind() {\n" +
	"	case reflect.String:\n" +
	"		return model{Name: key, Type: typeString}\n" +
	"	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,\n" +
	"		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:\n" +
	"		return model{Name: key, Type: typeInt}\n" +
	"	case reflect.Float32, reflect.Float64:\n" +
	"		return model{Name: key, Type: typeNumber}\n" +
	"	case reflect.Bool:\n" +
	"		return model{Name: key, Type: typeBool}\n" +
	"	default:\n" +
	"		return model{Name: key, Type: typeObject}\n" +
	"	}\n" +
	"	return model{}\n" +
	"}\n" +
	""

var templateServer = "// +build sample_swagger\n" +
	"\n" +
	"package sample_swagger\n" +
	"\n" +
	"import (\n" +
	"	\"html/template\"\n" +
	"	\"net/http\"\n" +
	")\n" +
	"\n" +
	"var serverJson string\n" +
	"\n" +
	"var htmlTemp = `\n" +
	"<!-- HTML for static distribution bundle build -->\n" +
	"<!DOCTYPE html>\n" +
	"<html lang=\"en\">\n" +
	"<head>\n" +
	"    <meta charset=\"UTF-8\">\n" +
	"    <title>Swagger UI</title>\n" +
	"    <link rel=\"stylesheet\" type=\"text/css\"\n" +
	"          href=\"https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui.css\">\n" +
	"    <link rel=\"icon\" type=\"image/png\" href=\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAEPElEQVR4Ab1XA8xlWQx+a9u2bcRex17btj1G8Nu2bdu2bfOi02/yzjN/NWnuYdtbH42joKrqqbKy+aSsSD/yN0KS16t4PAKU5I0a/kZhD2dwVrNToKjyhUz4e0WR+ohUcgRwlgX5ke9ebI++nT+W3mYiU4Iwr9HwdAOVtHpSYuXPFF70IZDHP1Fxqwfv1RPOCMBdpvEuaDnJXLmY/yBOEFvfXKa8puP0b8Td9Jmnxib+E3EX5TYe4ztLOqFZiATQdJT51XyhQVyu6gqhX4KvAXGn8Jegq6myMwg0hFmamPa1dv9cMGfHopD8t0FsWxic/yaxwxoIoV5sy+ZxgrlHxisgsCPonvaioRAJFn0CDgd1AUPy38LFHcWgvDdI0Idjmqr+QvbYCWKo6go2uvi1zzlU0RFAq+tzNDRVa5fR4FQNrW7Mw/70te+5RnsVnQG66DAKUcS58HZTh4su/YIA43NtlN1wyK4A2fUHaWy2hQCxZV8b7f0cdJUuOpC0dLZnu/RgMY/Dx5QgogCwP/ohh9X9X9R9BKjtiTDby2k4ovUFue+kLyB1ipBDDJteaBlMIcDvoTc5LAC0CGgdTDPb+zv8Dl1oMu+nNVAFJshiOGAuQBoBfgu53mEBfgy8HEyodSjd0j58SUTEz9BABCYlrR4WD/eNl0FS+srnbIcF+NL7DIQdDUxWWdwvanUTGojSoKphgnxuenBf9APICdQ7Xoo5fe51mpEgYAQUc+zhDMZdowVgYNF3Eip+IADzrtOgnGKComJ4KLNuH3F40vLaNB2KfRxr1D6cSSvrs/RryHVgTJPzXTS10MPjM9nu1/LZGeoYycFZOhj7KC2tToIGZdUfMKIdVvi+MMGYNQH40n5cpqW1KSb2mFaAbOQD+AMEAHOaXuyFAPQrCwDhOkfycJYOxDxCi6sToIHQtC4Aq9iqCfbHPKw1QQnm9AWr92sjE5wJNDLBFzoT5MMEEN6uCSIwKbbmhBPCCc9y2gn7JystO2GLcEIpZnfCMMBmGCJdCwEQhpJBIrpzDxLR7QaJSHpam4rlPizkNh41u1DdFUqAfc6k4sh7baTiw8IBB8DbpBgtoZMxuhBT+iW2UGAQGXaZI3xHZ5sJEFf2jUkxupLWNhdFEvrZqPNVVGUKG5WdgWblGKV1bWOBhqbr7AmANMtnF6E5+sb3PKO98g5/AoCXWcesGDQkaKN2uiEJzHvNtCGx2JIlaOOT3NNf2jHmrmnPg6ZwvCTwstqUonEUQkAT2//z1w36QbnFbnuO1lkIAXXBJ34WjukEovtBK8c0DJir1zr8MNF2rySiI7fxCP1tIU9YaDgQanxn0eBhspkEmlt5mr1r8jRDJKCeI5+jqAAxxhoiwORppsww8/e39VhFuOChycIMkIOAJIM439bj1IJGTkcPhzaKvzGoZCinQIxRWJDbGZ/FWUfpngCleTNdmkrhIgAAAABJRU5ErkJggg==\n" +
	"\" sizes=\"32x32\"/>\n" +
	"    <link rel=\"icon\" type=\"image/png\" href=\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABhElEQVR4AZVTA0x1cRz9jPHD7DCnKc1htms2spv5smvINS/bjXF4xpRtzHvvf0//82zc7fLg/vjJ/1AUES8UW6NQxKG8P/F0PjcScxMDhco3SVIJYbVqj6cxvl2I9rlM8ByTz9rjKRCTZo3kBhPP3jyb0TSVipKBT0HPRondPJsgTeZ9TOQHFcX14/9JDHvWjf9zmTS6c2Zorj+vaxpwcncQIDy528eGptEZSQqosdeEeTNnFzFUJLVjf3H7YnG/a44nHVGwwiySC3h8O0bl8O8Ag/KhH3h8P3G9y8IWgFpG8MRK8+PkbjHMF2tgOn3LeeiYy0LHfDZ6l3LRNJ0G0/mK5JSQi7bZDFDrY0DQfL4qyTTIp5gnzWgA49kypnZLfQxCpPArRArHgSmwEJrjqdiLeOQooruNbA0Btoot8zc4vt3HprbJ2cZk2Fxt9AySCXVRDtL1s9Hxd99RFrM0YSShxMQoVrxG2d+kkemwJiwSK82TzxwcYtzKAHF068wtFIn+/A9tMmqI7MxzGAAAAABJRU5ErkJggg==\n" +
	"\" sizes=\"16x16\"/>\n" +
	"    <style>\n" +
	"        html {\n" +
	"            box-sizing: border-box;\n" +
	"            overflow: -moz-scrollbars-vertical;\n" +
	"            overflow-y: scroll;\n" +
	"        }\n" +
	"\n" +
	"        *,\n" +
	"        *:before,\n" +
	"        *:after {\n" +
	"            box-sizing: inherit;\n" +
	"        }\n" +
	"\n" +
	"        body {\n" +
	"            margin: 0;\n" +
	"            background: #fafafa;\n" +
	"        }\n" +
	"    </style>\n" +
	"</head>\n" +
	"\n" +
	"<body>\n" +
	"<div id=\"swagger-ui\"></div>\n" +
	"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui-bundle.js\"></script>\n" +
	"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.18.2/swagger-ui-standalone-preset.js\"></script>\n" +
	"<script>\n" +
	"\n" +
	"    var spec = {{.Spec}};\n" +
	"\n" +
	"	spec = JSON.parse(spec);\n" +
	"\n" +
	"    window.onload = function () {\n" +
	"        const ui = SwaggerUIBundle({\n" +
	"            spec: spec,\n" +
	"            dom_id: '#swagger-ui',\n" +
	"            deepLinking: true,\n" +
	"            presets: [\n" +
	"                SwaggerUIBundle.presets.apis,\n" +
	"                SwaggerUIStandalonePreset.slice(1) // here\n" +
	"            ],\n" +
	"            layout: \"StandaloneLayout\"\n" +
	"        });\n" +
	"        window.ui = ui;\n" +
	"    }\n" +
	"</script>\n" +
	"</body>\n" +
	"</html>\n" +
	"`\n" +
	"\n" +
	"func ServerHTTP(w http.ResponseWriter, r *http.Request) {\n" +
	"	if serverJson == \"\" {\n" +
	"		serverJson = parse()\n" +
	"		if serverJson == \"\" {\n" +
	"			serverJson = \"{}\"\n" +
	"		}\n" +
	"	}\n" +
	"	t, err := template.New(\"swagger\").Parse(htmlTemp)\n" +
	"	if err != nil {\n" +
	"		w.Write([]byte(err.Error()))\n" +
	"		return\n" +
	"	}\n" +
	"	err = t.Execute(w, map[string]string{\"Spec\": serverJson})\n" +
	"	if err != nil {\n" +
	"		w.Write([]byte(err.Error()))\n" +
	"	}\n" +
	"}\n" +
	""

var templateServer2 = "// +build !sample_swagger\n" +
	"\n" +
	"package sample_swagger\n" +
	"\n" +
	"import \"net/http\"\n" +
	"\n" +
	"func ServerHTTP(w http.ResponseWriter, r *http.Request) {\n" +
	"   w.Write([]byte(`Please use build tag \"sample_swagger\" to open swagger!`))\n" +
	"}\n" +
	""

const templateVars = "// +build sample_swagger\n" +
	"\n" +
	"package sample_swagger\n" +
	"\n" +
	"import (\n" +
	"       {{Imports}}\n" +
	")\n" +
	"\n" +
	"var generatorJson = {{GeneratorJson}}\n" +
	"\n" +
	"var generatorModels = []interface{}{\n" +
	"{{GeneratorModels}}\n" +
	"}\n" +
	""
