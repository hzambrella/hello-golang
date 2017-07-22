package routes

import "net/http"

func init() {
	http.Handle("/test", ReqURLPrt(http.HandlerFunc(Test)))
}

func Test(w http.ResponseWriter, r *http.Request) {
	String(w, 500, "hello")
}
