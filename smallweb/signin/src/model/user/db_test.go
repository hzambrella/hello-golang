package user

import (
	"engine/datastore"
	"fmt"
	"testing"
)

var data datastore.Data = make(datastore.Data, 0)
var db UserDB
var dsnCfg string

func init() {
	data, err := datastore.ParseDataFromFile("../../../etc/db.cfg")
	if err != nil {
		panic(err)
	}

	dsnCfg = data["dbname"] + data["dsn"]
	fmt.Println(dsnCfg)
	db, err = NewUserDB(dsnCfg)
	if err != nil {
		panic(err)
	}

}

func TestGetUserByName(t *testing.T) {
	user, err := db.GetUserByName("11")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}
