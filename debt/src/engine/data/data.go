package data
import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
)

type Relation struct{
	OwnerId int64 `json:"ownerid"`//zhai quan  ren
	DebtorId int64	`json:"debtorid"`// zhai wu ren
	Amount float64 `json:"amount"`// amount
}

type Data []*Relation

var (
	ErrDataNotFound=errors.New("data isn't exist")
	ErrDataExist=errors.New("data is already exist ,can't add")
	fileName string="data"
)

func New()(Data){
	data,err:=parseDataFromFile(fileName)
	if err!=nil{
		panic(err)
	}
	return data
}

//add
func (d Data)Add(relation *Relation)error{
	d=append(d,relation)
	return saveDataToFile(fileName,d)
}

//get
func (d Data)Get(key int)(*Relation,error){
	result:=d[key]
	if result==nil{
		return nil,ErrDataNotFound
	}else{
		return result,nil
	}
}

//set
func (d Data)Set(key int,relation *Relation)(error){
	d[key]=relation
	return saveDataToFile(fileName,d)
}

//delete
func (d Data)Delete(key int)error{
	d=append(d[:key],d[key+1:]...)
	return saveDataToFile(fileName,d)
}

func parseDataFromFile(fileName string)(Data,error){
	dByte,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		log.Println(err)
		return nil,err
	}

	d:=make(Data,0)

	if len(dByte)==0{
		return d,nil
	}

	if err:=json.Unmarshal(dByte,&d);err!=nil{
		log.Println(err)
		return nil,err
	}
	return d,nil

}

func saveDataToFile(fileName string,d Data)error{
	dbyte,err:=json.MarshalIndent(d,"","")
	if err!=nil{
		return err
	}

	return ioutil.WriteFile(fileName,dbyte,0666)
}
