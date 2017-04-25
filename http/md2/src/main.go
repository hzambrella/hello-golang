package main

//practice http middleware
//only  serve apointed host
import(
	"fmt"
	"net/http"
)

var allowedHost string="localhost:8080"
func SingleHost(h http.Handler,allowedHost string)http.Handler{
	outfunc:=func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.Host)
		if allowedHost==r.Host{
			w.Write([]byte("we use middleware2\n"))
			h.ServeHTTP(w,r)
			w.Write([]byte("we use middleware2\n"))
		}else{
			w.Write([]byte("forbid"))
	//	w.WriteHeader(403)
		}
	}
//	return outfunc
	return http.HandlerFunc(outfunc)
}

func myHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello world\n"))
}

func main(){
	s:=SingleHost(http.HandlerFunc(myHandler),allowedHost)
	fmt.Println("listen and serve at 8080")
	err:=http.ListenAndServe(":8080",s)
	if err!=nil{
		fmt.Println(err.Error())
	}
}
