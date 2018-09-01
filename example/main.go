package main

import (
	"fmt"
	"github.com/heramerom/sample-swagger/example/handler"
	"github.com/heramerom/sample-swagger/example/sample-swagger"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由
	http.HandleFunc("/v1/class/detail", handler.SayHello)
	http.HandleFunc("/swagger.html", sample_swagger.ServerHTTP)

	log.Fatal(http.ListenAndServe(":8089", nil))
}
