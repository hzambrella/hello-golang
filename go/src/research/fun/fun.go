package fun
import(
	"fmt"
)


func fun(){
	//anonymous function 
	func (i int){
		fmt.Println("i=",i)
	}(2)//this is input

	//return by func
	out:=funType(1)
	fmt.Println("out1 c+a+b=",out(1,2))
	fmt.Println("out2 c+a+b=",out(1,2))

	//function  type
	funcType2(out).Serve(1,2)
}

//function as return
func funType(c int)func(a int,b int)(int){
	out:=func(a int, b int)int{
		c=c+1//close package
		fmt.Println("c=",c)
		return c+a+b
	}
	return out
}

//func type
type funcType2 func(a int,b int)(int)

func (f funcType2)Serve(a,b int){
	fmt.Println("f(a,b)=",f(a,b))
}
