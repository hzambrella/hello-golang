package slice 
import(
	"fmt"
)

type slice struct{
	a []int
}

var aa []int=[]int{1,2,3,4,5,6}

func sliceStudy(){
	s:=new(slice)
	s.a=aa
	s.remove(4)
	fmt.Println(s.a)
}

func (a *slice)remove(key int){
	fmt.Println(a.a[:key-1])
	fmt.Println(a.a[key-1:])
	a.a=append(a.a[:key-1],a.a[key:]...)
	fmt.Println(a.a)
}
