package datastore
import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var (
	driver string="mysql"
	dsn string=" ----copy  ws/life/etc.cfg------ "
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

	db,err:=sql.Open(driver,dsn)
	if err!=nil{
		panic(err)
	}


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

	defer db.Close()
}
