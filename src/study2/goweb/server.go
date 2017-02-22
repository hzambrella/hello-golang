package main
import(
	"fmt"
	"time"
	"strings"
	"net/http"
	"log"
	"html/template"
	"session"
)

type User struct{
	username string
	password string
}

var globalSession *session.Manager

func  main(){
	http.HandleFunc("/login",login)
	http.HandleFunc("/login2",loginBySession)
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/cookie",getCookie)
	fmt.Println("listen and serve at 8080:")
	http.ListenAndServe(":8080",nil)
}

func login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println("method:",r.Method)

	if r.Method=="GET"{
		fmt.Fprintln(w,"hello")
		t,err:=template.ParseFiles("template/login.tmpl")

		if err!=nil{
			log.Fatal(err)
			return
		}

		t.Execute(w,nil)
	}else{
		User1:=&User{
			username:r.Form["username"][0],
			password:r.Form["password"][0],
		}
		fmt.Fprintln(w,User1.username,User1.password)
		expirement:=time.Now().AddDate(1,0,0)
		cookie:=&http.Cookie{
			Name:"testCookie",
			Value:User1.username,
			Expires:expirement,
		}

		http.SetCookie(w,cookie)
		fmt.Println(cookie)
	}
	fmt.Println("200,login")
}

func loginBySession(w http.ResponseWriter,r *http.Request){
	sess:=globalSessions.SessionStart(w,r)
}
func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Method,":",r.URL.Host)
	fmt.Println(r.URL.Path,r.URL.Scheme,r.URL.Scheme,r.URL.Query)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Println(http.StatusOK,"index")
}

func getCookie(w http.ResponseWriter, r *http.Request ){
	r.ParseForm()
	cookie,err:=http.Cookie("testCookie")
	if err!=nil{
		log.Fatal(err)
		fmt.Println(http.StatusNotFound,"cookie wrong")
		return
	}
	fmt.Println(cookie)
	fmt.Fprintln(w,cookie)
	fmt.Println("200,getCookie")
}
