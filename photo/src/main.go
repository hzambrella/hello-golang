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
	mux.HandleFunc(helloPath,router.SayHello)
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

func safeHandleFunc()http.HandleFunc{
}
