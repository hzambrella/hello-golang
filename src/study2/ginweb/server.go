package main

import (
	"fmt"
	"time"
	"net/http"
	"engine/github.com/gin-gonic/gin"
)
type user struct{
	Username string
	Password string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/", index)
	r.GET("/login", login)
	r.POST("/login", pass)
	r.GET("/getCookie", getCookie)
	r.Run(":8080")
}

func index(c *gin.Context) {
	c.String(200, "hello")
}

func login(c *gin.Context) {
	c.HTML(200, "login", gin.H{})
}

func pass(c *gin.Context) {
	user.Username:= c.PostForm("username")
	user.Password:= c.PostForm("password")
	c.String(200, fmt.Sprintln(name, password))
}

func getCookie(c *gin.Context){
	req:=c.Request
	req.ParseForm()
	w:=c.Writer
	req.Form["username"]
	expiration:=time.Now()
	expiration=expiration.AddDate(1,0,0)
	cookie:=http.Cookie{
		cookie:user.Username,
		Expire:expiration ,
	}

	for _,cookie:=range req.Cookies(){
		fmt.Println("reading cookie")
		fmt.Fprintln(w,cookie.Name)
	}
}
