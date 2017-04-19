package fun
import(
	"fmt"
)

func fun(){
	//anonymous function 
	func (i int){
		fmt.Println(i)
	}(2)//this is input
}
