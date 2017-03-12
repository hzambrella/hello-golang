// important thing I should say three times!
// don't use http.Error()


package main
import (
	"net/http"
	"router"
	"log"
	_"lib/template"
)
const(
	helloPath="/hello"
	loginViewPath="/login/view"
	loginPath="/login"


	uploadPath="/upload"
	viewPath="/view"
	listViewPath="/listview"

	TEMPLATE_DIR="./public"
)

//TemplateSlice:=make(map[string]*template.Template)

func main(){

	mux:=http.NewServeMux()
	lenOfRouter:=20

	routerSlice:=[]string{helloPath,loginViewPath,loginPath,uploadPath,viewPath,listViewPath}

	mux.HandleFunc(helloPath,router.SayHello)
	mux.HandleFunc(loginViewPath,router.LoginView)
	mux.HandleFunc(loginPath,router.Login)

	mux.HandleFunc(uploadPath,router.Upload)
	mux.HandleFunc(viewPath,router.View)
	mux.HandleFunc(listViewPath,router.ListView)

	addr:=":8080"
	for _,v:=range routerSlice {
		if len(v)<lenOfRouter{
			vend:=len(v)
			for ;vend<lenOfRouter;vend++{
				v=v+" "
			}
		log.Println("Router "+v+ " is register!")
		}
	}

	log.Println("runing at",addr)

	if	err:=http.ListenAndServe(addr,mux);err!=nil{
		log.Fatal("ListenAndServe",err)
	}
}

func safeHandleFunc(fn http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if e,ok:=recover().(error);ok{
				http.Error(w,e.Error(),http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Warn panic in %v - %v",fn,e)
			}
		}()
		fn(w,r)
	}
}

