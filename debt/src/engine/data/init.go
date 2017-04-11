package data
import(
	"fmt"
	"os"
	myos"lib/os"
)

var data=[]*Relation{
	&Relation{OwnerId:1,DebtorId:2,Amount:10},
	&Relation{OwnerId:2,DebtorId:3,Amount:20},
	&Relation{OwnerId:3,DebtorId:1,Amount:30},
	&Relation{OwnerId:5,DebtorId:6,Amount:40},
	&Relation{OwnerId:7,DebtorId:8,Amount:10},
	&Relation{OwnerId:8,DebtorId:9,Amount:5},
}

func init(){
	exist,err:=myos.CheckFilesExist(fileName)
	if err!=nil{
		panic(err)
	}
	if exist{
		if err:=os.Remove(fileName);err!=nil{
			panic(err)
		}
	}

	_,err=os.Create(fileName)
	if err!=nil{
		panic(err)
	}

	for _,v:=range data{
		_,err:=AddAll(v)
		if err!=nil{
			panic(err)
		}
	}
	fmt.Println("add all ok")
}
