package sample_swagger

import (
	sw_model "github.com/heramerom/sample-swagger/example/model"
)

var generatorJson = `{"swagger":"","info":{"description":"hello world","version":"1.0.0","title":"swagger example","termsOfService":"","contact":{"email":"heramerom@gmail.com"},"license":{"name":"","url":""}},"host":"","basePath":"/v1","schemes":null,"paths":{"/class/detail":{"get":{"tags":["class"],"summary":" class-detail","description":"","operationId":"","consumes":null,"produces":null,"parameters":[{"in":"query","name":"id","type":"string","description":" class id","required":false}],"responses":{"200":{"description":"","schema":{"type":"object","items":{"$ref":""},"$ref":"#/definitions/model.Class"}}}},"post":{"tags":["class"],"summary":" class-detail","description":"","operationId":"","consumes":null,"produces":null,"parameters":[{"in":"query","name":"id","type":"string","description":" class id","required":false}],"responses":{"200":{"description":"","schema":{"type":"object","items":{"$ref":""},"$ref":"#/definitions/model.Class"}}}}}},"definitions":null}`

var generatorModels = []interface{}{
	new(sw_model.Class),
}
