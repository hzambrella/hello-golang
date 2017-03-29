package hzerror
import(
	"fmt"
	"time"
//	"testing"
)

type HzError struct{
	When time.Time
	What string
}

// interface
//type error interface{
// Error() string
//}

func (e HzError)Error() string{
	return fmt.Sprintf("%v:%v",e.When,e.What)
}

func oops()error{
	return HzError{
		time.Date(1989,3,15,22,30,0,0,time.UTC),
		"hz error study",
	}
}

func Example(){
	err:=oops()
	if err=oops();err!=nil{
		fmt.Println(err)
	}
}
