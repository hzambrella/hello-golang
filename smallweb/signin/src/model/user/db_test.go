package user

import (
	"fmt"
	"testing"
)

/*单元测试，用来检查错误。在写包的时候，最好做这个测试，用处是1.检查，2.用例子告诉别人怎么用你写的包  3.排查bug
测试指令
	go test
	go test -run=TestGetUserByName
开发时间紧时，可以不写
*/
var db UserDB

func init() {

	var err error
	db, err = NewUserDB()
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

func TestAddUser(t *testing.T) {
	fmt.Println("TestAddUser")
	uid, err := db.AddUser("nihao", "caonima")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ok   ", uid)
}

func TestUpdateUserStatus(t *testing.T) {
	fmt.Println("TestUpdateUserStatus")
	_, err := db.UpdateUserStatus("123", "2")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ok   ")
}
