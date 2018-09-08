package sample_swagger

import (
	sw_model "github.com/heramerom/sample-swagger/example/model"
)

var generatorJson = `{"swagger":"","info":{"description":"hello world","version":"1.0.0","title":"swagger example","termsOfService":"","contact":{"email":"heramerom@gmail.com"},"license":{"name":"","url":""}},"host":"","basePath":"/v1","schemes":null,"paths":{"/class/body":{"post":{"tags":["class"],"summary":"class-title, router desc","description":"","operationId":"","consumes":null,"produces":null,"parameters":[{"in":"body","name":"","type":"object","description":"class id","required":true,"schema":{"$ref":"#/definitions/model.Class"}}],"responses":{"200":{"description":"","schema":{"type":"object","items":{"$ref":""},"$ref":"#/definitions/model.Class"}}}}},"/class/detail":{"get":{"tags":["class"],"summary":"class-detail","description":"","operationId":"","consumes":null,"produces":null,"parameters":[{"in":"query","name":"id","type":"string","description":"class id","required":false}],"responses":{"200":{"description":"","schema":{"type":"object","items":{"$ref":""},"$ref":"#/definitions/model.Class"}}}},"post":{"tags":["class"],"summary":"class-detail","description":"","operationId":"","consumes":null,"produces":null,"parameters":[{"in":"query","name":"id","type":"string","description":"class id","required":false}],"responses":{"200":{"description":"","schema":{"type":"object","items":{"$ref":""},"$ref":"#/definitions/model.Class"}}}}}},"definitions":null}`

var generatorModels = []interface{}{
	new(sw_model.Class),
	new(sw_model.Sub),
	new(sw_model.Self),
	new(sw_model.NestObject),
	new(sw_model.ArrayObject),
}
