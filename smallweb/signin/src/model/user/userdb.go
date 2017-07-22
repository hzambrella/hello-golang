package user

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type UserDB interface {
	Ping() error
	Close()
	GetUserByName(name string) (*User, error)
}

var (
	driver string = "mysql"
	// 封装错误，方便route层针对不同的错误进行不同的处理。如仅仅是查找结果为空和数据库系统错误的处理方式就不一样
	UserDataNotFound = errors.New("user data not found")
)

//userDB要实现UserDB的接口
type userDB struct {
	*sql.DB
}

func NewUserDB(dsn string) (UserDB, error) {
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
	userdb := &userDB{db}
	return userdb, nil
}

func (db *userDB) Ping() error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (db *userDB) Close() {
	db.Close()
}

const (
	getUserByName = `
SELECT
	*
FROM
	user_info
WHERE
	user_name=?
`
)

func (db *userDB) GetUserByName(name string) (*User, error) {
	user := User{}
	err := db.QueryRow(getUserByName, name).Scan(
		&user.UserId,
		&user.Password,
		&user.UserName,
		&user.Status,
		&user.CreateTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserDataNotFound
		} else {
			return nil, err
		}
	}
	return &user, nil
}
