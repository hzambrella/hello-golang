package main
import(
	"fmt"
)
type Human struct{
	name string
}

func(h Human)String()string{
return h.name+"hello"
}


func main(){
	bob:= Human{"bob"}
	fmt.Println(bob)
fmt.Println(bob.String())
}
