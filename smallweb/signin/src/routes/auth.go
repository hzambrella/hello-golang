package routes

import (
	"encoding/json"
	"engine/datastore"
	"errors"
	"net/http"
	"time"
	//uuid是唯一识别码
	"github.com/satori/go.uuid"
)

//用户鉴权,模仿柳丁的auth.go
/*
步骤：
1.读取cookie：  若为空，进行step2:用户登陆。若不为空，取cookie,cookie.value是sessionid,用sessionid从缓存中换取用户信息。
2.用户登陆：读取原始请求链接，缓存，重定向到登陆页面对应的链接。用户在登陆页面进行登陆。
3.完成登陆：若用户登陆成功，取出缓存的原始请求链接，缓存用户信息于session,session存于cookie。
4.注册同理

这个流程的效果是，用户请求页面时，若没登陆过，就是登陆页面。登陆过的话，浏览器的页面上会知道用户是谁，不需要重复登陆。
*/

/* 鉴权的目的是因http的无状态性。服务器要想办法记住客户端的信息。
鉴权拥有两种机制，cookie和session。
cookie直接将用户信息编码后存入cookie。一旦让人摸清服务器的编码加密机制，就会暴露个人信息。
session将个人信息缓存至服务端，将sessionid存到cookie中，虽然还是能被firefox等伪造，但相对安全。
这里采用的是session。uuid is session id
*/

//TODO:缓存的时间,缓存要有时间限制
//成熟的缓存器：redis ,memcached, 基于消息队列的beantalkd

//cookie存活时间
const COOKIE_MAX_MAX_AGE = time.Hour * 1 / time.Second // 单位：秒。
var (
	//cookies name
	cookie string = "user_login"
	maxAge        = int(COOKIE_MAX_MAX_AGE)
)

//原始请求链接的简单缓存。没有缓存时间功能。
var reqUrlStore datastore.Data = make(datastore.Data, 0)

//用户信息
type userInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"user_name"`
}

var userNotFound = errors.New("user not found in session")

type session map[string][]byte

func NewSession() session {
	sess := make(session, 0)
	return sess
}

// 用户信息的简单缓存。没有缓存时间功能。
var userSession session = NewSession()

func (s session) GetUser(key string) (*userInfo, error) {
	ub, ok := s[key]
	if !ok {
		return nil, userNotFound
	}
	user := &userInfo{}
	err := json.Unmarshal(ub, &user)
	if err != nil {
		return nil, err
	}

	s.Delete(key)
	return user, nil

}

func (s session) SetUser(key string, user *userInfo) error {
	ub, err := json.Marshal(user)
	if err != nil {
		return err
	}
	s[key] = ub
	return nil

}

func (s session) Delete(key string) {
	delete(s, key)
}

func (u *userInfo) setCookie(w http.ResponseWriter) {
	uuid := uuid.NewV4().String()
	userSession.SetUser(uuid, u)

	uid_cookie := &http.Cookie{
		Name:     cookie,
		Value:    uuid,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   maxAge,
	}

	http.SetCookie(w, uid_cookie)
}

func auth(w http.ResponseWriter, r *http.Request) (*userInfo, bool) {
	ck, err := r.Cookie(cookie)
	if err != nil {
		logl.Println(err)
		login(w, r)
		return nil, false
	}

	u, err := userSession.GetUser(ck.Value)
	if err != nil {
		if err == userNotFound {
			login(w, r)
			return nil, false
		} else {
			logl.Error(err)
			return nil, false
		}
	}

	u.setCookie(w)

	return u, true
}
