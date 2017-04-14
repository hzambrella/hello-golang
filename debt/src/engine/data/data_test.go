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
	ds:=New()
	fmt.Println("before add")
	fmt.Println(ds.DStore)
	for _,v:=range ds.DStore{
		fmt.Println(v)
	}

	err:=ds.Add(&Relation{
		OwnerId:1,
		DebtorId:2,
		Amount:50,
	})
	if err!=nil{
		t.Fatal(err)
	}

	fmt.Println("add ok")
	for _,v:=range ds.DStore{
		fmt.Println(v)
	}
}

func TestDelete(t *testing.T){
	ds:=New()
	fmt.Println("before delete")
	for _,v:=range ds.DStore{
		fmt.Println(v)
	}
	if err:=ds.Delete(len(ds.DStore)-1);err!=nil{
		t.Fatal(err)
	}
	fmt.Println("delete ok")
	for _,v:=range ds.DStore{
		fmt.Println(v)
	}
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
