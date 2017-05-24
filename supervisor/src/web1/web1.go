package web1
import(
	"loghz"
	"net/http"
)

func web1(){
	loghz.Println(":8181")
	http.HandleFunc("/test1",Test1)
	http.ListenAndServe(":8181",nil)
}

func Test1(resp http.ResponseWriter,req *http.Request){
	loghz.Println(req)

}
