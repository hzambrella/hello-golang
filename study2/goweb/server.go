package main
import(
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func  main(){
	http.HandleFunc("/login",login)
	http.HandleFunc("/",index)
	fmt.Println("listen and serve at 8080:")
	http.ListenAndServe(":8080",nil)
}

func login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if r.Method=="GET"{
		fmt.Fprintln(w,"hello")
		t,err:=template.ParseFiles("template/login.tmpl")
		if err!=nil{
			log.Fatal(err)
			return
		}
		t.Execute(w,nil)
	}else{
		fmt.Fprintln(w,r.Form["username"],r.Form["password"])
	}
}

func index(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintln(w,r.Method,r.Host,r.URL.Scheme,r.URL.Query)
}
