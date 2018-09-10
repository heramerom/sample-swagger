package model

import "time"

type Student struct {
	Name string
	Age  string
}

// @sw:m	github.com/heramerom/sample-swagger/example/model, model.Class, json, "hello"
type Class struct {
	Name     string
	Students []Student
	Map      map[string]int
	Map2     map[string]Student
	Map3     map[string]struct {
		Int int `json:"int"`
	}
}

type Base struct {
	Name string
}

// @sw:m github.com/heramerom/sample-swagger/example/model, model.Sub,
type Sub struct {
	Base
	Age int

	unExportField string

	BirthDay time.Time `json:"birth_day"`

	Map  map[string]int
	Map2 map[string]struct {
		Name string `json:"name"`
	} `json:"map_2"`
	Map3 map[string]*Class
}

// @sw:m github.com/heramerom/sample-swagger/example/model, model.Self,
type Self struct {
	Value string
	Left  *Self
	Right *Self
}

// @sw:m github.com/heramerom/sample-swagger/example/model, model.ArrayObject,
type ArrayObject struct {
	Names []string
	Subs  []*Sub
}

// @sw:m github.com/heramerom/sample-swagger/example/model, model.NestObject,
type NestObject struct {
	Name string `json:"name"`
	Data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"data"`
}
