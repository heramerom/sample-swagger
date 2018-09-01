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
