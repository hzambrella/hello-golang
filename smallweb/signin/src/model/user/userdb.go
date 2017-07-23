package user

// model层尽量少的引用三方包，减小耦合。
// 在 vim 输入  "%s/要被修改的/修改/g"，即可完成批量修改。如"%s/u s e r/s i g n i n/g"
import (
	"database/sql"
	"engine/datastore"
	"errors"
	"fmt"
	"model"
)

type UserDB interface {
	Ping() error
	Close()
	GetUserByName(name string) (*User, error)
	AddUser(name, password string) (int, error)
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

// 一开始的sql.Open()方法写在这里，即没有缓存，这样做的坏处是，每次调用这个函数就连接一下数据库。连接数据库是很大的一笔系统开销。当用户多的时候，系统就会满载。
// 现在将其移动到model/db.go/GetDB方法
func NewUserDB() (UserDB, error) {
	data, err := datastore.ParseDataFromFile("../../../etc/db.cfg")
	if err != nil {
		return nil, err
	}

	dbname, ok := data["dbname"]
	fmt.Println("dbname:", dbname)
	if !ok || len(dbname) == 0 {
		return nil, errors.New("配置文件错误")
	}
	db, err := model.LinkStore.GetDB(dbname)
	if err != nil {
		return nil, err
	}

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
	getUserByNameSql = `
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
	err := db.QueryRow(getUserByNameSql, name).Scan(
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

const (
	addUserSql = `
INSERT INTO
	user_info
	(user_name,password,status)
VALUE
	(?,?,1)
	`
)

func (db *userDB) AddUser(name, password string) (int, error) {
	result, err := db.Exec(addUserSql, name, password)
	if err != nil {
		return 0, err
	}
	uid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(uid), err
}
