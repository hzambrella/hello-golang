package xml
import(
	"net/http"
	"log"
	"fmt"
	"bytes"
)

//design middleware
func checkXML(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//check requestbody
		if r.ContentLength==0{
			w.WriteHeader(400)
			w.Write([]byte("content is nil\n"))
			return
		}

		//check file type
		buf:=new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		b:=buf.Bytes()
		fmt.Println(string(b))
		fileType:=http.DetectContentType(b)
		fmt.Println("filetype:",fileType)
		if fileType!="text/xml;charset=utf-8"{
			w.WriteHeader(400)
			w.Write([]byte("type wrong\n"))
		}
		next.ServeHTTP(w,r)
	})
}

func final (w http.ResponseWriter, r *http.Request){
		log.Println("execute final")
		w.Write([]byte("final md\n"))
}

func md(){
	finalHandle:=http.HandlerFunc(final)
	http.Handle("/",checkXML(finalHandle))
	log.Println("listen at :8080")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err.Error())
	}
}
