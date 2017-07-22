package routes

import (
	"engine/datastore"
	"fmt"
)

var datadb datastore.Data = make(datastore.Data, 0)
var dsnCfg string

func init() {
	datadb, err := datastore.ParseDataFromFile("../etc/db.cfg")
	if err != nil {
		panic(err)
	}

	dsnCfg = datadb["dbname"] + datadb["dsn"]
	fmt.Println(dsnCfg)
}
