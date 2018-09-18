### sample-swagger

A sample swagger tool for golang web app.


#### Install
``` sh
go install github.com/heramerom/sample-swagger
```

#### Usage
1. Add tags in source golang code.
```go
// tag ruls

// @sw:r  [http methods], path, url tags, description
// @sw:p  name, position, type, require, description
// @sw:res  response code, type, 
func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

}


// @sw:m  import package path, reference
type Response struct {

}

```    

2. Generator sample-swagger package
```sh
sample-swagger .
```

3. Add router handler
```go 
http.HandleFunc("/swagger.html", sample_swagger.ServerHTTP)
```

4. run web app
```sh
go run -tags sample_swagger main.go
```

5. access the swagger *http://location/swagger.html*
