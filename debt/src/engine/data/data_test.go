package data
import(
	"testing"
	"fmt"
)
/*
func TestGet(t *testing.T){
	d:=New()
	data,err:=d.Get("1")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(data)
}
*/
func TestAdd(t *testing.T){
	d:=New()
	fmt.Println("before add")
	for _,v:=range d{
		fmt.Println(v)
	}

	err:=d.Add(&Relation{
		OwnerId:1,
		DebtorId:2,
		Amount:50,
	})
	if err!=nil{
		t.Fatal(err)
	}

	fmt.Println("add ok")
	for _,v:=range d{
		fmt.Println(v)
	}
}

/*
func TestDelete(t *testing.T){
	d:=New()
	fmt.Println("before delete",d)
	d,err=d.Delete(len(d)-1)
	fmt.Println("delete ok",d)
	d:=New()
	err:=d.Delete("3")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("delete ok")
}

/*
func TestSet(t *testing.T){
	d:=New()
	err:=d.Set("3","4")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("set ok")
}

*/
