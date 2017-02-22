package main

import (
	"log"
	"net/http"
	//"net/http/httputil"

	"engine/datastore"
)

var data datastore.Data

func main() {

	d, err := datastore.ParseDataFromFile("test")
	if err != nil {
		log.Println(err)
	}
	data = d
	http.HandleFunc("/test/view", View)

	http.HandleFunc("/test/add", Add)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
func View(w http.ResponseWriter, r *http.Request) {
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
			return
		}
	}()
	key := r.FormValue("key")
	if len(key) == 0 {
		result = data.ToJson()
		return
	}
	value := data.Get(key)
	if len(value) == 0 {
		result = []byte("no")
		return
	}
	result = []byte(value)
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
		result = []byte("a")
		return
	}
	val := r.FormValue("data")
	if len(data) == 0 {
		result = []byte("l")
		return
	}
	data.Set(key, val)
	if err := datastore.SaveDataToFile("test", data); err != nil {
		log.Fatal(err)
		return
	}
}
