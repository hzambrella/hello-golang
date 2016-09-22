package test

import(
	"fmt"
"testing"
)
func TestKrand(t *testing.T){
	result:=Krand(16,3)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(result)
}
