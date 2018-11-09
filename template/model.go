// +build sample_swagger

package template

type Info struct {
	Description    string `json:"description"`
	Version        string `json:"version"`
	Title          string `json:"title"`
	TermsOfService string `json:"termsOfService"`
	Contact        struct {
		Email string `json:"email"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

type Router struct {
	Tags        []string            `json:"tags,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationID string              `json:"operationId,omitempty"`
	Consumes    []string            `json:"consumes,omitempty"`
	Produces    []string            `json:"produces,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
}

type Schema struct {
	Ref string `json:"$ref,omitempty"`
}

type Parameter struct {
	In          string  `json:"in,omitempty"`
	Name        string  `json:"name,omitempty"`
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

type Response struct {
	Description string `json:"description"`
	Schema      struct {
		Type  string `json:"type"`
		Items struct {
			Ref string `json:"$ref"`
		} `json:"items"`
		Ref string `json:"$ref"`
	} `json:"schema"`
}

type Property struct {
	Type        string      `json:"type,omitempty"`
	Format      string      `json:"format,omitempty"`
	Description string      `json:"description,omitempty"`
	Ref         string      `json:"$ref,omitempty"`
	Items       *Definition `json:"items,omitempty"`
	Properties  *Definition `json:"properties,omitempty"`

	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
}

type AdditionalProperties struct {
	Type string `json:"type"`
	Ref  string `json:"$ref"`
}

type NestedProperty struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Definition struct {
	Type                 string                 `json:"type,omitempty"`
	Format               string                 `json:"format,omitempty"`
	Items                *Definition            `json:"items,omitempty"`
	Properties           map[string]*Definition `json:"properties,omitempty"`
	AdditionalProperties *Definition            `json:"additionalProperties,omitempty"`
	Ref                  string                 `json:"$ref,omitempty"`
}

type Swagger struct {
	Swagger     string                 `json:"swagger"`
	Info        *Info                  `json:"info"`
	Host        string                 `json:"host"`
	BasePath    string                 `json:"basePath"`
	Schemes     []string               `json:"schemes"`
	Paths       map[string]Method      `json:"paths"`
	Definitions map[string]*Definition `json:"definitions"`
}

type Method map[string]Router

func MapType(typ string) string {
	switch typ {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "string", "str", "s":
		return "string"
	case "bool", "boolean", "b":
		return "boolean"
	case "object", "obj", "o":
		return "object"
	case "float32", "float64":
		return "number"
	case "array", "slice":
		return "array"
	case "map":
		return "map"
	}
	return "{}" // any
}
