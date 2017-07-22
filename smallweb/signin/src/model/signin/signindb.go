package signin

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type SigninDB interface {
	Ping() error
	Close()
}

var (
	driver string = "mysql"
	// 封装错误，方便route层针对不同的错误进行不同的处理。如仅仅是查找结果为空和数据库系统错误的处理方式就不一样
	DataNotFound = errors.New("data not found")
)

//signinDB要实现SigninDB的接口
type signinDB struct {
	*sql.DB
}

func NewSigninDB(dsn string) (SigninDB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	/*
		err = db.Ping()
		if err != nil {
			return nil, err
		}
	*/
	signindb := &signinDB{db}
	return signindb, nil
}

func (db *signinDB) Ping() error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (db *signinDB) Close() {
	db.Close()
}
