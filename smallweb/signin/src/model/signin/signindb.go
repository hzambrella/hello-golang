package signin

// model层尽量少的引用三方包，减小耦合。
// 在 vim 输入  "%s/要被修改的/修改/g"，即可完成批量修改。如"%s/u s e r/s i g n i n/g",增加开发效率
import (
	"database/sql"
	"engine/datastore"
	"errors"
	"model"
	"os"
)

type SigninDB interface {
	Ping() error
	Close()
}

var (
	driver string = "mysql"
	// 封装错误，方便route层针对不同的错误进行不同的处理。如仅仅是查找结果为空和数据库系统错误的处理方式就不一样
	SigninDataNotFound = errors.New("signin data not found")
)

//signinDB要实现SigninDB的接口
type signinDB struct {
	*sql.DB
}

// 一开始的sql.Open()方法写在这里，即没有缓存，这样做的坏处是，每次调用这个函数就连接一下数据库。连接数据库是很大的一笔系统开销。当用户多的时候，系统就会满载。
// 现在将其移动到model/db.go/GetDB方法
func NewSigninDB() (SigninDB, error) {
	//	data, err := datastore.ParseDataFromFile("../../../etc/db.cfg")
	data, err := datastore.ParseDataFromFile(os.Getenv("ETCDIR") + "/db.cfg")
	if err != nil {
		return nil, err
	}

	dbname, ok := data["dbname"]
	if !ok || len(dbname) == 0 {
		return nil, errors.New("配置文件错误")
	}
	db, err := model.LinkStore.GetDB(dbname)
	if err != nil {
		return nil, err
	}

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
