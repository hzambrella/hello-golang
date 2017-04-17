package data
import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
)

type DataS struct{
	 DStore Data
}

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

func New()*DataS{
	data,err:=getAll()
	if err!=nil{
		panic(err)
	}
	return &DataS{DStore:data}
}

func getAll()(Data,error){
	data,err:=parseDataFromFile(fileName)
	if err!=nil{
		return nil,err
	}
	return data,nil
}

func writeDataToFile(d Data)error{
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


//add
func (d *DataS)Add(relation *Relation)error{
	d.DStore=append(d.DStore,relation)
	return saveDataToFile(fileName,d.DStore)
}

//get
func (d *DataS)Get(key int)(*Relation,error){
	result:=d.DStore[key]
	if result==nil{
		return nil,ErrDataNotFound
	}else{
		return result,nil
	}
}

//set
func (d *DataS)Set(key int,relation *Relation)(error){
	d.DStore[key]=relation
	return saveDataToFile(fileName,d.DStore)
}

//delete
func (d *DataS)Delete(key int)error{
	d.DStore=append(d.DStore[:key],d.DStore[key+1:]...)
	return saveDataToFile(fileName,d.DStore)
}

