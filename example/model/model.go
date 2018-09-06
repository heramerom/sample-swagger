package model

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
}

// @sw:m github.com/heramerom/sample-swagger/example/model, model.Self,
type Self struct {
	Value string
	Left  *Self
	Right *Self
}
