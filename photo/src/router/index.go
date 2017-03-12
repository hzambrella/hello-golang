package router
import(
	"net/http"
	"log"
	"io"
	"lib/code"
//	"fmt"

	tmpl"lib/template"
)


func LoginView(w http.ResponseWriter,r *http.Request){
		log.Println("router/Login.go:is called")
		if err:=tmpl.RenderHTML(w,"photo/login",nil);err!=nil{
			log.Println("router/Login.go:GET TMPl fali: "+err.Error())
		}
		log.Println(http.StatusOK,"loginView")
		return
}

func Login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	userName:=r.FormValue("username")
	userpasswd:=r.FormValue("pasword")
	log.Println(userName+"in"+userpasswd)

	// TODO:count people num that online
	if err:=makeUserKey(w,r,userName);err!=nil{
		log.Println("/router/Index.go:login,fail to makeUserKey",err.Error())
	}

	log.Println("success to  makeUserKey")
	log.Println(http.StatusOK,"login")
//	http.Redirect(w,r,r.URL.Host+fmt.Sprintf("/hello?name=%s",userName),302)
	http.Redirect(w,r,"/hello",302)
	//http.Redirect(w,r,r.URL.Host+"/listview",302)
	return
}

func SayHello(w http.ResponseWriter,r *http.Request){
//	userName:=r.FormValue("name")
	userSession:= &UserSession{}
	cookie,err:=r.Cookie("userSess")

	if err!=nil{
		//http.Error(w,"poor connection",404)
		log.Println("doAuth.go,cookie:",err.Error())
		return
	}

	userSess:=cookie.Value
	log.Println("SayHello:userSess:",userSess)

	if code.Decode(userSess,userSession);err!=nil{
		log.Fatal("doAuth.go:get cookie fail",err)
		return
	}
	userName:=userSession.Name
	log.Println("sayHello:userName:",userName)

	helloMes:="hello "+userName+"!"
	_,err=io.WriteString(w,helloMes)

	if err!=nil{
		log.Fatal("SayHello:",err)
		http.Error(w,err.Error(),500)
		return
	}

	log.Println(http.StatusOK,r.Method,"SayHello")
	return
}
