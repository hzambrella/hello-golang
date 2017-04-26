package main

import(
	"net/http"
	"fmt"
	"html/template"
	"io"
)

func main(){
	fmt.Println("listen and serve:8080")
	http.Handle("/hello",printUrl(http.HandlerFunc(Hello)))
	h:=http.StripPrefix("/public/",http.FileServer(http.Dir("public")))
	http.Handle("/public/",printUrl(h))
	http.ListenAndServe(":8080",nil)
}

func printUrl(next http.Handler)http.Handler{
	out:=func(w http.ResponseWriter,r *http.Request){
		fmt.Println("urlpath:",r.URL.Path)
		fmt.Println("host",r.Host)
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(out)
}

func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello")
	_,err:=io.WriteString(w,"hello v1")
	if err!=nil{
		fmt.Println(err.Error())
	}
}

