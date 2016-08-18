package datastore

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Data map[string]string

var l sync.Mutex

func Parse(src []byte) (Data, error) {

	d := make(Data)

	if err := json.Unmarshal(src, &d); err != nil {
		return nil, err
	}
	return d, nil
}

func (d Data) ToJson() []byte {
	l.Lock()
	defer l.Unlock()
	result, err := json.MarshalIndent(d, "", "	")
	if err != nil {
		panic(err)
	}
	return result
}

func (d Data) Get(key string) string {

	l.Lock()
	defer l.Unlock()
	result, _ := d[key]
	return result
}

func (d Data) Set(key, val string) {
	l.Lock()
	defer l.Unlock()
	d[key] = val
}

// read from file
func ParseDataFromFile(fileName string) (Data, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// save to file
func SaveDataToFile(fileName string, d Data) error {
	return ioutil.WriteFile(fileName, d.ToJson(), 0666)
}
