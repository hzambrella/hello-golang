package datastore

import (
	"errors"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"sync"
)

var lock sync.Mutex

type Data map[string]string

//
func Parse(src []byte) (Data, error) {
	lock.Lock()
	defer lock.Unlock()
	d := make(Data)
	fmt.Println("data.go:14",src)

	if err := json.Unmarshal(src, &d); err != nil {
		return nil, err
	}
	return d, nil
}

//
func (d Data) ToJson() []byte {
	lock.Lock()
	defer lock.Unlock()
	result, err := json.MarshalIndent(d, "","	")
	if err != nil{
		panic(err)
	}
	return result
}

func (d Data) Get(key string) string {
	lock.Lock()
	defer lock.Unlock()
	result, _ := d[key]
	return result
}

func (d Data) Set(key, val string) {
	lock.Lock()
	defer lock.Unlock()
	d[key] = val
}

// read from file
func ParseDataFromFile(fileName string)(Data, error){
	_,err:=os.Stat("test")
	if err!=nil{
		if os.IsNotExist(err){
			panic(err)
		}else{
			panic(errors.New("file not exist"))
		}
	}

	fmt.Println("data.go 55:file exist")
	data, err := ioutil.ReadFile(fileName)
	fmt.Println("data.go 45",data)
	if err != nil{
		return nil, err
	}
	return Parse(data)
}

// save to file
func SaveDataToFile(fileName string, d Data) error{
	return ioutil.WriteFile(fileName, d.ToJson(), 0666)
}
