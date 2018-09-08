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
