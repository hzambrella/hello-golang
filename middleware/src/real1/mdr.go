package real1
import(
	"net/http"
	"log"
)

//design middleware
func mdone(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("execute 1")
		w.Write([]byte("first md\n"))
		next.ServeHTTP(w,r)
		log.Println("execute 1 again")
		w.Write([]byte("first md again\n"))
	})
}

func mdtwo(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("execute 2")
		w.Write([]byte("second md\n"))
		next.ServeHTTP(w,r)
		log.Println("execute 2 again")
		w.Write([]byte("second md again\n"))
	})
}
func final (w http.ResponseWriter, r *http.Request){
		log.Println("execute final")
		w.Write([]byte("final md\n"))
}

func md(){
	finalHandle:=http.HandlerFunc(final)
	http.Handle("/",mdone(mdtwo(finalHandle)))
	log.Println("listen at :8080")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err.Error())
	}
}
