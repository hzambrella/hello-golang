package main
import(
	"fmt"
	//"time"
	"strings"
	"net/http"
	"log"
	"html/template"
	"session"
	_"memory"
)

var globalSession *session.Manager

func init(){
	globalSession,_=session.NewSessionManager("memory","goSessionId",60)
	go globalSession.GC()
	fmt.Println("main:init complete")
}

func  main(){
	http.HandleFunc("/login",login)
	http.HandleFunc("/login2",loginBySession)
	http.HandleFunc("/",sayHello)
//	http.HandleFunc("/cookie",getCookie)
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
			username:=r.Form["username"]
			password:=r.Form["password"]
		fmt.Println(w,username," come in")
		fmt.Fprintln(w,username,password)
	//	expirement:=time.Now().AddDate(1,0,0)
	//	cookie:=&http.Cookie{
	//		Name:"testCookie",
	//		Value:User1.username,
	//		Expires:expirement,
	//	}

	//	http.SetCookie(w,cookie)
	//	fmt.Println(cookie)
	}
	fmt.Println("200,login")
}

func loginBySession(w http.ResponseWriter,r *http.Request){
	sess:=globalSession.SessionStart(w,r)
	r.ParseForm()
	if r.Method=="GET"{
		t,_:=template.ParseFiles("login")
		t.Execute(w,sess.Get("username"))
	}else{
		sess.Set("username",r.Form["username"])
		http.Redirect(w,r,"/",302)
	}
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
	fmt.Fprintln(w,"study cookie and session")
	fmt.Println(http.StatusOK,"sayhello")
}

/*
func getCookie(w http.ResponseWriter, r *http.Request ){
	r.ParseForm()
	cookie,err:=http.Cookie()
	if err!=nil{
		log.Fatal(err)
		fmt.Println(http.StatusNotFound,"cookie wrong")
		return
	}
	fmt.Println(cookie)
	fmt.Fprintln(w,cookie)
	fmt.Println("200,getCookie")
}
*/
