package main

import(
	"net/http"
	"fmt"
	"io"
)

func main(){
	fmt.Println("listen and serve:8080")
	http.HandleFunc("/hello",Hello)
	http.ListenAndServe(":8080",nil)
}


func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello")
	_,err:=io.WriteString(w,"hello v1")
	if err!=nil{
		fmt.Println(err.Error())
	}
}
