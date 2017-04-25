package main
import(
	"io"
	"net/http"
	"log"
	"fmt"
)

func main(){
	mux:=http.NewServeMux()
	// Handle(patter,handler)
	// handler is interface,need ServeHTTP
//	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/hello",Hello)
	log.Println("listen and serve:8080")
	// wrong use :nil
//	http.ListenAndServe(":8080",nil)
	http.ListenAndServe(":8080",mux)
}

type myHandler struct{
}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println("/")
	io.WriteString(w,r.URL.String())
}

func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello")
	_,err:=io.WriteString(w,"hello v2")
	if err!=nil{
		fmt.Println(err.Error())
	}
}
