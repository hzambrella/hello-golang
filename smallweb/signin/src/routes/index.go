package routes

import (
	"fmt"
	"model/user"
	"net/http"
)

const (
	signinHtmlPath = "/signin/index"
)

func init() {
	http.Handle(signinHtmlPath, ReqURLPrt(http.HandlerFunc(signinHTML)))
}

func signinHTML(w http.ResponseWriter, r *http.Request) {
	u, ok := auth(w, r)
	if !ok {
		logl.Println("0o")
		return
	}

	name := u.UserName

	userDB, err := user.NewUserDB()
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
