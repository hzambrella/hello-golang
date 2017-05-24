package web2
import(
	"loghz"
	"net/http"
)

func web2(){
	loghz.Println(":8182")
	http.HandleFunc("/test1",Test1)
	http.ListenAndServe(":8182",nil)
}

func Test1(resp http.ResponseWriter,req *http.Request){
	loghz.Println(req)

}
