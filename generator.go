package main

import (
	"encoding/json"
	models "github.com/heramerom/sample-swagger/template"
	"strings"
)

type Api struct {
	swagger     models.Swagger
	definitions []definition
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) AddRouters(routers ...router) {
	if a.swagger.Paths == nil {
		a.swagger.Paths = make(map[string]models.Method)
	}
	for _, r := range routers {
		a.swagger.Paths[r.path] = r.toMethod()
	}
}

func (a *Api) AddDefinitions(definitions ...definition) {
	a.definitions = append(a.definitions, definitions...)
}

func (a *Api) Json() string {
	bs, _ := json.MarshalIndent(a.swagger, "", "  ")
	return string(bs)
}

func (a *Api) DefinitionImports() string {
	ism := make(map[string]struct{})
	for _, d := range a.definitions {
		ism[d.path] = struct{}{}
	}
	var is string
	for path := range ism {
		paths := strings.Split(path, "/")
		is += "sw_" + paths[len(paths)-1] + " \"" + path + "\"" + "\n"
	}
	return is
}

func (a *Api) DefinitionObjects() string {
	var ds string
	for _, d := range a.definitions {
		ds += "new(sw_" + d.model + "),\n"
	}
	return ds
}
