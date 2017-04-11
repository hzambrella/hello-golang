package model
import(
	"testing"
	"fmt"
)

func TestGet(t *testing.T){
	d:=New()
	data,err:=d.Get("1")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(data)
}

func TestAdd(t *testing.T){
	d:=New()
	err:=d.Add("3","4")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("add ok")
}

func TestSet(t *testing.T){
	d:=New()
	err:=d.Set("3","4")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("set ok")
}

func TestDelete(t *testing.T){
	d:=New()
	err:=d.Delete("3")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("delete ok")
}
