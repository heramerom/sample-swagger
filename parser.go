package main

import (
	"errors"
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
	router.Tags = strings.Split(r.tag, ",")
	for _, v := range r.params {
		router.Parameters = append(router.Parameters, v.toParameters())
	}
	router.Responses = responses(r.response).toResponses()
	for _, v := range r.methods {
		method[v] = router
	}
	return models.Method(method)
}

func (r *router) String() string {
	return fmt.Sprintf("[%s] %s %s", strings.Join(r.methods, ","), r.path, r.desc)
}

type param struct {
	name    string
	typ     string
	object  string
	in      string
	require string
	desc    string
}

func (p *param) toParameters() models.Parameter {
	require, _ := strconv.ParseBool(p.require)
	parameter := models.Parameter{
		Name:        p.name,
		In:          p.in,
		Type:        models.MapType(p.typ),
		Description: p.desc,
		Required:    require,
	}
	if p.object != "" {
		parameter.Schema = &models.Schema{
			Ref: "#/definitions/" + p.object,
		}
	}
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
	resp.Schema.Type = models.MapType(r.typ)
	switch resp.Schema.Type {
	case "object":
		if r.model != "" {
			resp.Schema.Type = "object"
			resp.Schema.Ref = "#/definitions/" + r.model
		}
	case "array":
		if r.model != "" {
			resp.Schema.Type = "array"
			resp.Schema.Items.Ref = "#/definitions/" + r.model
		}
	case "map":
		// do nothing
	}
	return resp
}

// @sw:r get,post, path, tag, desc
func parseRouter(s *Scanner) router {
	var r router
	var next string
Loop:
	for {
		next = s.nextString(',')
		switch next {
		case "get", "post", "option", "put", "delete", "patch":
			r.methods = append(r.methods, next)
		default:
			break Loop
		}
	}

	r.path = next
	r.tag = s.nextString(',')
	r.desc = s.nextString()
	return r
}

func parseParam(s *Scanner) param {
	p := param{}
	p.name = s.nextString(',')
	p.in = s.nextString(',')
	p.typ = s.nextString(',')
	if p.typ == "obj" || p.typ == "object" {
		p.typ = "object"
		p.object = s.nextString(',')
	}
	p.require = s.nextString(',')
	p.desc = s.nextString()
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
func parseResponse(s *Scanner) response {
	var resp response
	next := s.nextString(',')
	resp.code = next
	if next == "" {
		resp.code = "default"
	}
	next = s.nextString(',')
	resp.typ = strings.ToLower(next)
	if resp.typ == "" {
		resp.typ = "string"
	}
	switch resp.typ {
	case "string", "s":
		resp.typ = "string"
		resp.desc = s.nextString()
	case "object", "obj", "o":
		resp.typ = "object"
		next = s.nextString(',')
		resp.model = next
		resp.desc = s.nextString()
	case "array", "a":
		resp.typ = "array"
		next = s.nextString(',')
		resp.model = next
		resp.desc = s.nextString()
	}
	return resp
}

type definition struct {
	path  string
	model string
	tag   string
	desc  string
}

// @sw:m import_path,package,desc
func parseModel(s *Scanner, pkgPath string, pkg string, scanner *FileScanner) (def definition, err error) {
	def.path = s.nextString(',')
	def.model = s.nextString(',')
	def.desc = s.nextString()
	if def.path == "" {
		def.path = pkgPath
	}
	if def.model == "" {
		def.model, err = getNextModel(scanner)
		def.model = pkg + "." + def.model
	}
	return
}

func getNextModel(scanner *FileScanner) (model string, err error) {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "type") {
			model = strings.Split(strings.TrimSpace(strings.Replace(line, "type", "", 1)), " ")[0]
			return
		}
		if line != "" {
			err = errors.New("syntax error: model define")
			return
		}
	}
	return
}
