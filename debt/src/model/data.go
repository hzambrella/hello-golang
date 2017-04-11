package model
import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

var (
	ErrDataNotFound=errors.New("data isn't exist")
	ErrDataExist=errors.New("data is already exist ,can't add")
	fileName string="data"
)


type Data map[string]string

type DataOp interface{
	Get(key string)(string,error)
	Add(key ,val string)error
	Set(key ,val string)error
	Delete(key string)error
}

func New()(Data){
	data,err:=parseData(fileName)
	if err!=nil{
		panic(err)
	}
	return data
}

//get
func (d Data)Get(key string)(string,error){
	result,ok:=d[key]
	if !ok{
		return "",ErrDataNotFound
	}else{
		return result,nil
	}
}

//add
func (d Data)Add(key,val string)error{
	_,ok:=d[key]
	if ok{
		return ErrDataExist
	}
	d[key]=val
	return saveDataToFile(fileName,d)
}

//set
func (d Data)Set(key,val string)error{
	_,ok:=d[key]
	if !ok{
		return ErrDataNotFound
	}

	d[key]=val
	return saveDataToFile(fileName,d)
}

//delete
func (d Data)Delete(key string)error{
	delete(d,key)
	return saveDataToFile(fileName,d)
}

func parseData(fileName string)(Data,error){
	dByte,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		log.Println(err)
		return nil,err
	}

	d:=make(Data)
	if err:=json.Unmarshal(dByte,&d);err!=nil{
		log.Println(err)
		return nil,err
	}
	fmt.Println("data:",d)
	return d,nil

}

func saveDataToFile(fileName string,d Data)error{
	dbyte,err:=json.Marshal(d)
	if err!=nil{
		return err
	}

	return ioutil.WriteFile(fileName,dbyte,0666)
}
