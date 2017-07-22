package routes

import (
	"io"
	"net/http"
)

func init() {
	http.Handle("/test", ReqURLPrt(http.HandlerFunc(Test)))
}

func Test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}
