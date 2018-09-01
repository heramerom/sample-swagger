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
	Tags        []string            `json:"tags"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	OperationID string              `json:"operationId"`
	Consumes    []string            `json:"consumes"`
	Produces    []string            `json:"produces"`
	Parameters  []Parameter         `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
}

type Parameter struct {
	In          string `json:"in"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Schema      *struct {
		Ref string `json:"$ref"`
	} `json:"schema,omitempty"`
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
	Type        string          `json:"type,omitempty"`
	Format      string          `json:"format,omitempty"`
	Description string          `json:"description,omitempty"`
	Ref         string          `json:"$ref,omitempty"`
	Items       *Property       `json:"items,omitempty"`
	Properties  *NestedProperty `json:"properties,omitempty"`
}

type NestedProperty struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Definition struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
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