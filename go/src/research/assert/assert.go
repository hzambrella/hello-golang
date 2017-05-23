package assert
import(
	"fmt"
)

//var e map[string]interface{}=map[string]interface{"haha":map[string]string{"he":"he","ha":"ha"},"hiahia":map[string]string{"1":"1","2","2"},"papa":map[string]string{"pa","pa"}}

func assert(){
//	fmt.Println(e)
	v,ok:=interface{}(uint32(12)).(int)
	//false
	fmt.Println(v,ok)
}
