package routes

import "net/http"

//中间件，可以不要
//打印请求的链接
func ReqURLPrt(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logl.Info("request from", r.Host+r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
