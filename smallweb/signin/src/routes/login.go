package routes

import (
	"errors"
	"fmt"
	"model/user"
	"net/http"

	"github.com/satori/go.uuid"
)

//用户登录，模仿柳丁的login.go
//TODO:密码的加密解密。
//TODO:用户登录信息database记录
//TODO:修改个人信息，忘记密码，短信验证码（没发短信的功能接口）
//TODO:微信阿里的网页鉴权

const (
	//登录主页
	LoginIndexPath = "/login/index"
	//账号密码验证及跳转原始链接
	DoLoginPath = "/login/api"
	// 注册主页
	RegisterIndexPath = "/register/index"
	// 注册提交
	DoRegisterPath = "/register/api"
	JumpStr        = "?state=%s"
)

func init() {
	http.Handle(LoginIndexPath, ReqURLPrt(http.HandlerFunc(loginIndex)))
	http.Handle(DoLoginPath, ReqURLPrt(http.HandlerFunc(doLogin)))
	http.Handle(RegisterIndexPath, ReqURLPrt(http.HandlerFunc(registerIndex)))
	http.Handle(DoRegisterPath, ReqURLPrt(http.HandlerFunc(doRegister)))
}

func login(w http.ResponseWriter, r *http.Request) {
	//state
	state := uuid.NewV4().String()
	//缓存原始请求链接，key:state  value:链接
	reqUrlStore.Set(state, r.Host+r.RequestURI)
	//一定注意，加上http://，否则不会重定向
	reLink := "http://" + r.Host + LoginIndexPath + fmt.Sprintf(JumpStr, state)
	logl.Println("redirect to:", reLink)
	//TODO weixin and ali
	http.Redirect(w, r, reLink, 302)
	return
}

//登录主页
func loginIndex(w http.ResponseWriter, r *http.Request) {
	state := FormValue(r, "state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil,您是否已经登陆过？")
		return
	}

	logl.Println("login index")

	Render(w, 200, "public/user/login.html",
		H{
			"state": state,
		})
}

//账号密码验证及跳转原始链接
func doLogin(w http.ResponseWriter, r *http.Request) {
	userName := FormValue(r, "username")
	if len(userName) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error(errors.New("user name is nil"))
		String(w, 400, "请输入用户名")
		return
	}

	password := FormValue(r, "password")
	if len(password) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error(errors.New("password is nil"))
		String(w, 400, "请输入密码 ")
		return
	}

	state := FormValue(r, "state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil,您是否已经登陆过？")
		return
	}

	//TODO:密码的加密解密。如base64,XXtea
	userDB, err := user.NewUserDB()
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	userinfo, err := userDB.GetUserByName(userName)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	if userinfo.Password != password {
		logl.Error(errors.New("密码错误"), userName, password, userinfo.Password)
		String(w, 400, "密码错误")
		return
	}

	if userinfo.Status != 1 {
		String(w, 400, "你被封号了，请注意你的行为举止")
	}

	//TODO:用户登录信息database记录
	logl.Info("LOGIN", "欢迎"+userName+"登录系统")

	reLink := reqUrlStore.Get(state)
	delete(reqUrlStore, state)
	if len(reLink) == 0 {
		logl.Error(errors.New("relink is nil"))
		String(w, 500, "出错了！,请重新打开网页 relink is nil")
		return
	}

	u := &userInfo{
		UserName: userinfo.UserName,
		Uid:      userinfo.UserId,
	}
	u.setCookie(w)
	logl.Println(reLink)
	//	http.Redirect(w, r, "http://"+reLink, 302)
	JSON(w, 200,
		H{
			"relink": "http://" + reLink,
		})
	return
}

func registerIndex(w http.ResponseWriter, r *http.Request) {
	state := FormValue(r, "state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil,您是否已经登陆过？")
		return
	}

	Render(w, 200, "public/user/register.html",
		H{
			"state": state,
		})
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	userName := FormValue(r, "username")
	if len(userName) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error(errors.New("user name is nil"))
		String(w, 400, "请输入用户名")
		return
	}

	password := FormValue(r, "password")
	if len(password) == 0 {
		// 验证输入的工作尽量在前端完成
		logl.Error(errors.New("password is nil"))
		String(w, 400, "请输入密码 ")
		return
	}

	state := FormValue(r, "state")
	if len(state) == 0 {
		logl.Error(errors.New("state is nil"))
		String(w, 500, "出错了！,请重新打开网页 state is nil,您是否已经登陆过？")
		return
	}

	//TODO:密码的加密解密。如base64,XXtea
	userDB, err := user.NewUserDB()
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	_, err = userDB.GetUserByName(userName)
	if err != nil {
		if err == user.UserDataNotFound {
			//do nothing
		} else {
			logl.Error(err)
			String(w, 500, err.Error())
			return
		}
	} else {
		logl.Println(fmt.Sprintf("用户名 %s已注册. ", userName))
		String(w, 400, fmt.Sprintf("用户名 %s已注册. ", userName))
		return

	}

	uid, err := userDB.AddUser(userName, password)
	if err != nil {
		logl.Error(err)
		String(w, 500, err.Error())
		return
	}

	//TODO:用户登录信息database记录
	logl.Info("LOGIN", "欢迎"+userName+"注册系统")

	reLink := reqUrlStore.Get(state)
	delete(reqUrlStore, state)
	if len(reLink) == 0 {
		logl.Error(errors.New("relink is nil"))
		String(w, 500, "出错了！,请重新打开网页 relink is nil")
		return
	}

	u := &userInfo{
		UserName: userName,
		Uid:      uid,
	}

	u.setCookie(w)

	logl.Println(reLink)
	//一定注意，加上http://，否则不会重定向
	//http.Redirect(w, r, "http://"+reLink, 302)
	JSON(w, 200,
		H{
			"relink": "http://" + reLink,
		})
	return
}
