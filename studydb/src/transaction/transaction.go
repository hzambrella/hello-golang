package transaction
import (
	"io/ioutil"
	"fmt"
	"log"
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
	updateOrderTableSql=`
		UPDATE
			order_table
		SET
			table_name=?
		WHERE
			id=?
	`
)

func datastore(){
	fmt.Println("start study db!")

	defer func(){
		if err:=recover();err!=nil{
			log.Fatal(err)
		}
	}()

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

	//begin transaction
	//get conn
	tx,err:=db.Begin()
	if err!=nil{
		panic(err)
	}

	stmt,err:=tx.Prepare(updateOrderTableSql)
	if err!=nil{
		panic(err)
	}

	//tx.Exec(18)// don't use it wrong
	_,err=stmt.Exec("ha1",26)
	if err!=nil{
		panic(err)
	}
//sql.Tx一旦释放，连接就回到连接池中，这里stmt在关闭时就无法找到连接。所以必须在Tx commit或rollback之前关闭statement。
	stmt.Close()

	stmt,err=tx.Prepare(updateOrderTableSql)
	if err!=nil{
		panic(err)
	}
//故意制造一个panic,你会发现，前面的id=26的不会被修改成功
//	panic("do wrong deliberately")
	_,err=stmt.Exec("ha2",18)
	if err!=nil{
		panic(err)
	}

	stmt.Close()
	// if exception happend ,rollback
	defer tx.Rollback()
	tx.Commit()

}
