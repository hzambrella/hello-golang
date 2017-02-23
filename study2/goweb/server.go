package main
import(
	"fmt"
	"time"
	"strings"
	"net/http"
	"log"
	"html/template"
	"session"
	_"memory"
)

var globalSession *session.Manager

func init(){
	globalSession,_=session.NewSessionManager("memory","goSessionid",60)
	fmt.Println(globalSession)
	go globalSession.GC()
	fmt.Println("main:init complete")
}

func  main(){
	http.HandleFunc("/login",login)
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/count",count)
//	http.HandleFunc("/cookie",getCookie)
	fmt.Println("listen and serve at 8080:")
	http.ListenAndServe(":8080",nil)
}


func login(w http.ResponseWriter,r *http.Request){
	sess:=globalSession.SessionStart(w,r)
	fmt.Println("sess login:start:",sess)
	r.ParseForm()
	if r.Method=="GET"{
		t,err:=template.ParseFiles("template/login.tmpl")
		if err!=nil{
			log.Fatal(err)
			return
		}
		//w.Header().Set("Content-Type","text/html")
		t.Execute(w,sess.Get("username"))
	}else{
		sess.Set("username",r.Form["username"])
		fmt.Println("sess login post:",sess)
		http.Redirect(w,r,"/count",302)
	}
	fmt.Println(http.StatusOK,"login2")
}

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("method:",r.Method,"host:",r.URL.Host)
	fmt.Println("path:",r.URL.Path,"scheme",r.URL.Scheme,"Query:",r.URL.Query)
	fmt.Println("url_long:",r.Form["url_long"])
	fmt.Println("......................")
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintln(w,"study cookie and session")
	fmt.Println(http.StatusOK,"sayhello")
}



func count(w http.ResponseWriter,r *http.Request){
	sess:=globalSession.SessionStart(w,r)
	r.ParseForm()
	createtime:=sess.Get("createtime")

	if createtime==nil{
		sess.Set("createtime",time.Now().Unix())
	}else if (createtime.(int64)+360)<(time.Now().Unix()){
		globalSession.SessionDestory(w,r)
		sess=globalSession.SessionStart(w,r)
	}
	ct:=sess.Get("countnum")

	if ct==nil{
		sess.Set("countnum",1)
	}else{
		sess.Set("countnum",(ct.(int)+1))
	}

	t,err:=template.ParseFiles("template/count.tmpl")

	if err!=nil{
		log.Fatal(err)
		return
	}

//	w.Header().Set("content-type","text/html")
	t.Execute(w,sess.Get("countnum"))
	fmt.Println("sess: count",sess)
	fmt.Println(http.StatusOK,"count")
}

func observe(w http.ResponseWriter,r *http.Request){

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
