package recover
import(
	"fmt"
	"loghz"
//	"errors"
)

func recoverA(){
	defer func(){
		if err:=recover();err!=nil{
			loghz.Error(err.(error))
		}
	}()
	f()
}

func f(){
	aslice:=[]int{1,2,3}
	 fmt.Println(aslice[3])
}
