package main
import (
	"net/http"
	"router"
	"log"
	_"lib"
)
const(
	helloPath="/hello"
	uploadPath="/upload"
	viewPath="/view"
	listViewPath="/listview"
	TEMPLATE_DIR="./public"
)



//TemplateSlice:=make(map[string]*template.Template)

func main(){
	mux:=http.NewServeMux()
	lenOfRouter:=20
	routerSlice:=[]string{helloPath,uploadPath,viewPath,listViewPath}
	mux.HandleFunc(helloPath,safeHandleFunc(router.SayHello))
	mux.HandleFunc(uploadPath,safeHandleFunc(router.Upload))
	mux.HandleFunc(viewPath,safeHandleFunc(router.View))
	mux.HandleFunc(listViewPath,safeHandleFunc(router.ListView))
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
