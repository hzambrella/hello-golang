package main

import (
	"engine/datastore"
	"fmt"
	"log"
	"net/http"
)

var data datastore.Data

func main() {
	r := []byte(`{"test1":"d"}`)
	dat, err := datastore.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	datastore.SaveDataToFile("test", dat)

	d, err := datastore.ParseDataFromFile("test")
	if err != nil {
		log.Fatal(err)
		d = make(datastore.Data)
	}

	data = d
	fmt.Println("data ok:", data)
	http.HandleFunc("/test/view", View)
	http.HandleFunc("/test/add", Add)
	http.HandleFunc("/test/modify", Modify)
	http.HandleFunc("/test/delete", Delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func View(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var meiyou bool = true
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()
	key := r.FormValue("key")
	if len(key) == 0 {
		result = data.ToJson()
		return
	}
	for i := range data {
		if i == key {
			meiyou = false
		}
		if meiyou {
			result = []byte("查询不到相关KEY")
			return
		}
	}
	val := data.Get(key)
	result = []byte(val)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()
	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key不能为空 ")
		return
	}
	for i := range data {
		if i == key {
			result = []byte("已存在，不能添加")
			return
		}
	}
	val := r.FormValue("data")
	data.Set(key, val)
	result = []byte("add")
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
}
func Modify(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var meiyou bool = true
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
			return
		}
	}()
	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key can't be nil")
		return
	}
	for i := range data {
		if i == key {
			meiyou = false
		}
	}
	if meiyou {
		result = []byte("key isn't exist")
		return
	}
	val := r.FormValue("data")
	data.Set(key, val)
	result = []byte("modify")
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var meiyou bool = true
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()
	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key can't be nil")
		return
	}
	for i := range data {
		if i == key {
			meiyou = false
		}
	}
	if meiyou {
		result = []byte("key isn't exist")
		return
	}
	delete(data, key)
	result = []byte("delete")
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
}
