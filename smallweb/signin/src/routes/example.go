package routes

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"model/user"
	"net/http"
	"strconv"
)

func init() {
	http.Handle("/test/string", ReqURLPrt(http.HandlerFunc(TestString)))
	http.Handle("/test/json", ReqURLPrt(http.HandlerFunc(TestJSON)))
	http.Handle("/test/html", ReqURLPrt(http.HandlerFunc(TestHTML)))
	http.Handle("/test/html2", ReqURLPrt(http.HandlerFunc(test)))
}

func TestString(w http.ResponseWriter, r *http.Request) {
	String(w, 200, "hello")
	return
}

func TestJSON(w http.ResponseWriter, r *http.Request) {
	name := FormValue(r, "name")
	if len(name) == 0 {
		// 这里是日志打印,目的见log.go和相关资料
		logl.Error(errors.New("name is nil"))
		// 返回错误。建议中文. newding app 喜欢内嵌网页，多数错误都是url 参数不对，缺少参数，然后甩锅给我们，如果返回中文错误，就能甩锅
		String(w, 500, "缺少参数：name")
	}

	userDB, err := user.NewUserDB(dsnCfg)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	user, err := userDB.GetUserByName(name)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	JSON(w, 200, H{"user": user})
	return

}

func TestHTML(w http.ResponseWriter, r *http.Request) {
	name := FormValue(r, "name")
	if len(name) == 0 {
		// 这里是日志打印,目的见log.go和相关资料
		logl.Error(errors.New("name is nil"))
		// 返回错误。建议中文. newding app 喜欢内嵌网页，多数错误都是url 参数不对，缺少参数，然后甩锅给我们，如果返回中文错误，就能甩锅
		String(w, 500, "缺少参数：name")
	}

	userDB, err := user.NewUserDB(dsnCfg)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	user, err := userDB.GetUserByName(name)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}
	defer func() {
		if err := recover(); err != nil {
			logl.Error(err.(error))
		}
	}()
	fmt.Println(user)

	Render(w, 200, "public/user/info.html",
		H{
			"user": user,
		})
}

func test(w http.ResponseWriter, r *http.Request) {
	code := 200
	w.WriteHeader(code)
	w.Header().Set("codelog", strconv.Itoa(code))
	tempName := "public/user/info.html"
	t, err := template.ParseFiles(tempName)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.Execute(w, nil)
}
