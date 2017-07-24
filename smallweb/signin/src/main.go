package main

import (
	"engine/datastore"
	"engine/logz"
	"errors"
	"net/http"
	"os"
	_ "routes"
)

var port string

func init() {
	//serverData, err := datastore.ParseDataFromFile("../etc/server.cfg")
	serverData, err := datastore.ParseDataFromFile(os.Getenv("ETCDIR") + "/server.cfg")
	if err != nil {
		panic(err)
	}
	var ok bool
	port, ok = serverData["signin_port"]
	if !ok {
		panic(errors.New("no signin_port in server.cfg"))
	}
	logz.NewLogDebug(true).Info("listen at ", port)
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(port, nil)
}
