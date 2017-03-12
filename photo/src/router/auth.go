package router

import(
	"net/http"
	"log"
	"fmt"
	"time"
	"strconv"

	"lib/code"
)

type UserInfo struct{
	Name string
}

type UserSession struct{
	Token string
	Name string
}

const CookieExpire=60*1
func makeUserKey(w http.ResponseWriter, r *http.Request,name string)error{
	userKey:=strconv.FormatInt(time.Now().Unix(),10)+name
	redisClient.HMSet("userKey",map[string]string{userKey:userKey})

	userSessionNew:=UserSession{
		Token:userKey,
		Name:name,
	}

	if err:=userSessionNew.saveToCookie(w,r);err!=nil{
		log.Println("makeUserKey:saveToCookie",err.Error)
		return err
	}
	return nil
}

func (uss *UserSession)saveToCookie(w http.ResponseWriter,r *http.Request)error{
	NewSessionStr,err:=code.Encode(uss)
	if err!=nil{
		return err
	}
	http.SetCookie(w,&http.Cookie{
		Name:"userSess",
		Value:NewSessionStr,
		Path:"/",
		MaxAge:CookieExpire,
		HttpOnly:false,
	})
	return nil
}

func doAuth(w http.ResponseWriter,r *http.Request)(*UserInfo,bool){
	userSession:=&UserSession{}
	cookie,err:=r.Cookie("userSess")

	if err!=nil{
		//http.Error(w,"poor connection",404)
		log.Println("doAuth.go,cookie:",err.Error())
		return nil,false
	}

	userSess:=cookie.Value

	if code.Decode(userSess,userSession);err!=nil{
		log.Fatal("doAuth.go:get cookie fail",err)
		return nil,false
	}

	userKey:=userSession.Token
	userKeyRedis,err:=redisClient.HGet("userKey",userKey).Result()
	if err!=nil{
		log.Fatal("doAuth.go:get session in redis fail",err)
		return nil,false
	}

	if userKey!=userKeyRedis {
		log.Println("doauth.go:permission:diff userkey",err)
		return nil,false
	}

	redisClient.Del("auth",userKey)

	name:=userSession.Name
	if err=makeUserKey(w,r,name);err!=nil{
		log.Fatal("doAuth.go:session:Encode fail",err)
		return nil,false
	}
	log.Println("doAuth is success")

	return &UserInfo{
		Name:name,
	},true
}

func Auth(w http.ResponseWriter,r *http.Request)(*UserInfo,bool){
	log.Println("begin Auth")
	u,ok:=doAuth(w,r)
	log.Println("Auth.go :ready to redirect")

	if !ok{
		log.Println("Auth.go:user need auth!")
		http.Redirect(w,r,r.URL.Host+"/login/view",302)
	}else{
		log.Println(fmt.Sprintf("Auth.go:user:%s pass auth!",u.Name))
		http.Redirect(w,r,r.URL.Host+fmt.Sprintf("/hello?name=%s",u.Name),302)
	}
	return u,ok
}
