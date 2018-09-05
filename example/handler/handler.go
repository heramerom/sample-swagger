package handler

import "net/http"

// @sw:r  get,post, /class/detail, class, class-detail
// @sw:p  id, query, , , class id
// @sw:resp  200, object, model.Class,
func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// @sw:r  post, /class/body, class, class-title, router desc
// @sw:p  , body, object, model.Class, true , class id
// @sw:resp  200, object, model.Class,
func TestBody(w http.ResponseWriter, r *http.Request) {

}
