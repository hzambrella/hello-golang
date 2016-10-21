package main

import (
	"engine/datastore"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

var data datastore.Data

func main() {
	d, err := datastore.ParseDataFromFile("test")
	if err != nil {
		log.Println(err)
		d = make(datastore.Data)
	}
	data = d
	fmt.Println(data)
	http.HandleFunc("/test/view", View)
	http.HandleFunc("/test/add", Add)
	http.HandleFunc("/test/modify", Modify)
	http.HandleFunc("/test/delete", Delete)

	addr := ":8080"
	log.Println("Start at " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func View(w http.ResponseWriter, r *http.Request) {
	var result []byte
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
	val := data.Get(key)
	if len(val) == 0 {
		result = []byte("key not found")
		return
	}
	result = []byte(val)

}
func Add(w http.ResponseWriter, r *http.Request) {
	debug, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(debug))
	}
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key not found")
		return
	}

	for i := range data {
		if key == i {
			result = []byte("已经存在")
			return
		}
	}

	value := r.FormValue("data")
	if len(value) == 0 {
		result = []byte("need data value")
		return
	}
	data.Set(key, value)
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
	result = []byte("增加 ")

}
func Modify(w http.ResponseWriter, r *http.Request) {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()
	var k bool = true
	debug, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(debug))
	}
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key not found")
		return
	}
	for i := range data {
		if key == i {
			k = false
			break
		}
	}
	if k {
		result = []byte("不能改正key 不存在的 ")
		return
	}

	value := r.FormValue("data")
	if len(value) == 0 {
		result = []byte("need data value")
		return
	}
	data.Set(key, value)
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
	result = []byte("已修改 ")

}
func Delete(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var k bool = true

	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	key := r.FormValue("key")
	if len(key) == 0 {
		result = []byte("key not found")
		return
	}
	for i := range data {
		if key == i {
			k = false
			break
		}
	}
	if k {
		result = []byte("不能删除key 不存在的 ")
		return
	}
	delete(data, key)
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
	result = []byte("删除了")

}
