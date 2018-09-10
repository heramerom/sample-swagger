// +build !sample_swagger

package sample_swagger

import "net/http"

func ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Please use build tag "sample_swagger" to open swagger!`))
}
