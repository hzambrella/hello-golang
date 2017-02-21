package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/", index)
	r.GET("/login", login)
	r.POST("/login", pass)
	r.Run(":8080")
}

func index(c *gin.Context) {
	c.String(200, "hello")
}

func login(c *gin.Context) {
	c.HTML(200, "login", gin.H{})
}

func pass(c *gin.Context) {
	name:= c.PostForm("username")
	password:= c.PostForm("password")
	c.String(200, fmt.Sprintln(name, password))
}
