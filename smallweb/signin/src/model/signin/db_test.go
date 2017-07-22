package signin

import (
	"engine/datastore"
	"fmt"
)

var data datastore.Data = make(datastore.Data, 0)
var db SigninDB
var dsnCfg string

func init() {
	data, err := datastore.ParseDataFromFile("../../../etc/db.cfg")
	if err != nil {
		panic(err)
	}

	dsnCfg = data["dbname"] + data["dsn"]
	fmt.Println(dsnCfg)
	db, err = NewSigninDB(dsnCfg)
	if err != nil {
		panic(err)
	}

}
