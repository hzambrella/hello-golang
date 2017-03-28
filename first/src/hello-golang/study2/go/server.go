package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"
)


func  main(){
	http.HandleFunc("/test",test)
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
