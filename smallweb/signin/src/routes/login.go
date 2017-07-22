package routes

import (
	"errors"
	"fmt"
	"model/user"
	"net/http"

	"github.com/satori/go.uuid"
)

const (
	LoginIndexPath    = "/login/index"
	DoLoginPath       = "/login/api"
	RegisterIndexPath = "/register/index"
	DoRegisterPath    = "/register/api"
	JumpStr           = "?state=%s"
)

func init() {
	http.Handle(LoginIndexPath, ReqURLPrt(http.HandlerFunc(loginIndex)))
	http.Handle(DoLoginPath, ReqURLPrt(http.HandlerFunc(doLogin)))
	http.Handle(RegisterIndexPath, ReqURLPrt(http.HandlerFunc(registerIndex)))
	http.Handle(DoRegisterPath, ReqURLPrt(http.HandlerFunc(doRegister)))
}

func login(w http.ResponseWriter, r *http.Request) {
	state := uuid.NewV4().String()
	reqUrlStore.Set(state, r.Host+r.RequestURI)
	reLink := r.Host + LoginIndexPath + fmt.Sprintf(JumpStr, state)
	logl.Println("redirect to:", reLink)
	http.Redirect(w, r, reLink, 302)
}

func loginIndex(w http.ResponseWriter, r *http.Request) {
	state := FormValue("state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil")
		return
	}

	Render(w, 200, "public/user/login.html",
		H{
			"state": state,
		})
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	userName := FormValue("user_name")
	if len(userName) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error("user name is nil")
		String(w, 400, "请输入用户名")
		return
	}

	password := FormValue("password")
	if len(password) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error("password is nil")
		String(w, 400, "请输入密码 ")
		return
	}

	state := FormValue("state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil")
		return
	}

	//TODO:密码的加密解密。如base64,XXtea
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

	if user.Password != password {
		logl.Error(errors.New("密码错误"), username, password, user.Password)
		String(w, 400, "密码错误")
		return
	}

	if user.State != 1 {
		String(w, 400, "你被封号了，请注意你的行为举止")
	}

	//TODO:用户登录信息database记录
	logl.Info("LOGIN", "欢迎"+username+"登录系统")

	reLink := reqUrlStore.Get(state)
	if len(reLink) == 0 {
		logl.Error(errors.New("relink is nil"))
		String(w, 500, "出错了！,请重新打开网页 relink is nil")
		return
	}

	http.Redirect(w, r, reLink, 302)
}

func registerIndex(w http.ResponseWriter, r *http.Request) {
	state := FormValue("state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil")
		return
	}

	Render(w, 200, "public/user/register.html",
		H{
			"state": state,
		})
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	userName := FormValue("user_name")
	if len(userName) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error("user name is nil")
		String(w, 400, "请输入用户名")
		return
	}

	password := FormValue("password")
	if len(password) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error("password is nil")
		String(w, 400, "请输入密码 ")
		return
	}

	state := FormValue("state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil")
		return
	}

	//TODO:密码的加密解密。如base64,XXtea
	userDB, err := user.NewUserDB(dsnCfg)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	err := userDB.AddUser(name, password)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	if user.Password != password {
		logl.Error(errors.New("密码错误"), username, password, user.Password)
		String(w, 400, "密码错误")
		return
	}

	//TODO:用户登录信息database记录
	logl.Info("LOGIN", "欢迎"+username+"注册系统")

	reLink := reqUrlStore.Get(state)
	if len(reLink) == 0 {
		logl.Error(errors.New("relink is nil"))
		String(w, 500, "出错了！,请重新打开网页 relink is nil")
		return
	}

	http.Redirect(w, r, reLink, 302)
}
