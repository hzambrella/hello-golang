package main

//practice http middleware
//only  serve apointed host
import(
	"fmt"
	"net/http"
)


type SingleHost struct{
	handler http.Handler
	AllowedHost string
}

func (s *SingleHost)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Host)
	if s.AllowedHost==r.Host{
		s.handler.ServeHTTP(w,r)
	}else{
		w.Write([]byte("forbid"))
	//	w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello world"))
}

func main(){
	s:=&SingleHost{
		handler:http.HandlerFunc(myHandler),//notice it!!!
		//http.HandlerFunc acheived Handle interface
		// use func(f HandlerFunc)ServeHttp( w.r..){f(w,r)}
		// so s.handler.ServeHTTP(w,r)execute function myHandler(w,r) 
		AllowedHost:"localhost:8080",
	}
	fmt.Println("listen and serve at 8080")
	err:=http.ListenAndServe(":8080",s)
	if err!=nil{
		fmt.Println(err.Error())
	}
}
