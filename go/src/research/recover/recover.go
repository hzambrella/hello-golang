package recover
import(
	"fmt"
	"errors"
)

func recoverA(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	f()
}

func f(){
	errA:=errors.New("chucuola!")
	panic(errA)
}
