package datastore
import (
	"io/ioutil"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"

	"bs"
)
var (
	driver string="mysql"
//	dsn string=" ----copy  ws/life/etc.cfg------ "
//	dsn2 string="n2d_admin:[32n2d_admin15]@tcp(test.orientaltele.com:3306)/n2d_center?timeout=30s&strict=true&loc=Local&parseTime=true&allowOldPasswords=1"
)

const(
	getCitysByFirstLetterSql=`
	SELECT
		code,name,first_letter
	FROM
		life_city
	WHERE
		first_letter=?
	`
)

type City struct{
	Code string
	Name string
	FirstLetter string
}

func datastore(){
	fmt.Println("start study db!")

	/*
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
*/


	b,err:=ioutil.ReadFile("../db.cfg")
	if err!=nil{
		panic(err)
	}


	//b last byte is 10
	dsn:=bs.B2S(b[:len(b)-1])
	fmt.Println("46:",dsn)

	db,err:=sql.Open(driver,dsn)
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	err=db.Ping()
	if err!=nil{
		panic(err)
	}

	rows,err:=db.Query(getCitysByFirstLetterSql,"a")
	if err!=nil{
		panic(err)
	}
	fmt.Println("rows",rows)

	defer rows.Close()


	cityList:=make([]*City,0)
	for rows.Next(){
		city:=&City{}
		err:=rows.Scan(
			&city.Code,
			&city.Name,
			&city.FirstLetter,
		)

		if err!=nil{
			panic(err)
		}
		fmt.Println("city",city)
		cityList=append(cityList,city)
	}

	if err:=rows.Err();err!=nil{
		panic(err)
	}

	fmt.Println("cityList",cityList)
	for _,v:=range cityList{
		fmt.Println(v)
	}

}
