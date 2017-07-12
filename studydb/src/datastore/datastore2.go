package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driver string = "mysql"
	dsn    string = "haozhao:haozhoa@tcp(192.168.0.114:3306)/hz_db"
)

const (
	getUserLoginSql = `
	SELECT
		id
	FROM
		user_login
	WHERE
		id=?
	`
)

func main() {
	fmt.Println("start study db!")

	/*
		defer func(){
			if err:=recover();err!=nil{
				fmt.Println(err)
			}
		}()
	*/

	/*
		b, err := ioutil.ReadFile("../db.cfg")
		if err != nil {
			panic(err)
		}

		//b last byte is 10
			dsn := B2S(b[:len(b)-1])
			fmt.Println("46:", dsn)
	*/

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var id int64
	err = db.QueryRow(getUserLoginSql, "1").Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

}

/*
func B2S(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}
*/
