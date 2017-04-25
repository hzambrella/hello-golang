package main
import(
	"io"
	"net/http"
	"log"
	"fmt"
	"time"
)

var mux map[string]func(w http.ResponseWriter,r *http.Request)
func main(){
	myHandler:=&myHandler{}
	server:=&http.Server{
		Addr:":8080",
		Handler:myHandler,
		WriteTimeout:5*time.Second,
	}

	mux=make(map[string]func(w http.ResponseWriter,r *http.Request),0)
	mux["/hello"]=Hello

	log.Println("listen and serve :8080")
	server.ListenAndServe()
}

type myHandler struct{
}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	if v,ok:= mux[r.URL.String()];ok {
		v(w,r)
	}
}

func Hello(w http.ResponseWriter,r *http.Request){
	_,err:=io.WriteString(w,"hello v3")
	if err!=nil{
		fmt.Println(err.Error())
	}
}
