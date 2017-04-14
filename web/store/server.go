package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"
	"strconv"
)

var testmap map[string]string
var num int=1

func  main(){
	testmap=make(map[string]string)
	http.HandleFunc("/test",test)
	http.HandleFunc("/ping",ping)
	http.HandleFunc("/pong",pong)
	fmt.Println("listen and serve at 8080:")
	http.ListenAndServe(":8080",nil)
}


func test(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
		t,err:=template.ParseFiles("index.html")
		if err!=nil{
			log.Fatal(err)
			return
		}
		t.Execute(w,nil)
}

func ping(w http.ResponseWriter,r *http.Request){
	fmt.Println(testmap)
	fmt.Println(num)
	log.Println(200,"testmap")

}

func pong(w http.ResponseWriter,r *http.Request){
	numstr:=strconv.Itoa(num)
	testmap[numstr]=numstr
	num=num+1
	fmt.Println(testmap)
	fmt.Println(num)
}
