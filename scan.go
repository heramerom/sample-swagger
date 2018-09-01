package main

import (
	"bytes"
	"fmt"
	models "github.com/heramerom/sample-swagger/template"
	"strconv"
	"strings"
)

type router struct {
	path     string
	methods  []string
	tag      string
	desc     string
	params   []param
	response []response
}

var emptyRouter = router{}

func (r *router) toMethod() models.Method {
	method := make(map[string]models.Router, 1)
	var router models.Router
	router.Summary = r.desc
	router.Tags = strings.Split(r.tag, "/")
	for _, v := range r.params {
		fmt.Println(v)
		router.Parameters = append(router.Parameters, v.toParameters())
	}
	router.Responses = responses(r.response).toResponses()
	for _, v := range r.methods {
		method[v] = router
	}
	return models.Method(method)
}

type param struct {
	name    string
	typ     string
	in      string
	require string
	desc    string
}

func (p *param) toParameters() models.Parameter {
	require, _ := strconv.ParseBool(p.require)
	parameter := models.Parameter{
		Name:        p.name,
		In:          p.in,
		Type:        p.typ,
		Description: p.desc,
		Required:    require,
	}
	fmt.Println(parameter)
	return parameter
}

type responses []response

func (rs responses) toResponses() map[string]models.Response {
	var r = make(map[string]models.Response)
	for _, v := range rs {
		r[v.code] = v.toResponse()
	}
	return r
}

type response struct {
	code  string
	typ   string
	model string
	desc  string
}

func (r response) toResponse() models.Response {
	var resp models.Response
	resp.Description = r.desc
	if r.typ == "object" {
		if strings.HasPrefix(r.model, "#") {
			resp.Schema.Type = "object"
			resp.Schema.Ref = r.model
		}
	} else if r.typ == "array" {
		if strings.HasPrefix(r.model, "#") {
			resp.Schema.Type = "array"
			resp.Schema.Items.Ref = r.model
		}
	}
	return resp
}

// @sw:r get,post, path, tag, desc
func parseRouter(line string) router {
	var r router
	var next string
Loop:
	for {
		next, line = scanNext(line, ',')
		switch next {
		case "get", "post", "option", "put", "delete", "patch":
			r.methods = append(r.methods, next)
		default:
			break Loop
		}
	}

	r.path = next
	next, line = scanNext(line, ',')
	r.tag = next
	r.desc = line
	return r
}

func parseParam(line string) param {
	p := param{}
	p.name, line = scanNext(line, ',')
	p.in, line = scanNext(line, ',')
	p.typ, line = scanNext(line, ',')
	p.require, line = scanNext(line, ',')
	p.desc = line
	if p.typ == "" {
		p.typ = "string"
	}
	if p.in == "" {
		p.in = "query"
	}
	if p.require == "" {
		p.require = "false"
	}
	return p
}

// 100,obj,model,description
func parseResponse(line string) response {
	var resp response
	next, line := scanNext(line, ',')
	resp.code = next
	if next == "" {
		resp.code = "default"
	}
	next, line = scanNext(line, ',')
	resp.typ = strings.ToLower(next)
	if resp.typ == "" {
		resp.typ = "string"
	}
	switch resp.typ {
	case "string", "s":
		resp.typ = "string"
		resp.desc = line
	case "object", "obj", "o":
		resp.typ = "object"
		next, line = scanNext(line, ',')
		resp.model = "#/definitions/" + next
		resp.desc = line
	case "array", "a":
		resp.typ = "array"
		next, line = scanNext(line, ',')
		resp.model = "#/definitions/" + next
		resp.desc = line
	}
	return resp
}

type definition struct {
	path  string
	model string
	tag   string
	desc  string
}

// @sw:m import_path,package,tag,desc
func parseModel(line string) definition {
	var def definition
	def.path, line = scanNext(line, ',')
	def.model, line = scanNext(line, ',')
	def.tag, line = scanNext(line, ',')
	def.desc = line
	return def
}

func scanNext(line string, sepSet ...byte) (s1, s2 string) {
	if len(sepSet) == 0 {
		s1 = line
		return
	}
	line = strings.TrimLeft(line, string(sepSet))
	buf := bytes.NewBufferString("")
	for idx, char := range line {
		for _, sep := range sepSet {
			if char == int32(sep) {
				s1 = strings.TrimSpace(buf.String())
				s2 = strings.TrimSpace(strings.TrimLeft(line[idx:], string(sepSet)))
				return
			}
		}
		buf.WriteByte(byte(char))
	}
	return
}
