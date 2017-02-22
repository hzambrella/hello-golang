package main

import "github.com/gin-gonic/gin"
// import "github.com/gin-gonic/gin/binding"
import "net/http"

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"nil":     "haha",
			"nill":    nil,
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
        //默认的值
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.DefaultPostForm("message", "adasds")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// ce shi
	// curl -v --form user=user --form password=password http://localhost:8080/login

	router.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.BindWith(&form, binding.Form)
		// or you can simply use autobinding with Bind method:
		var form LoginForm
		// in this case proper binding will be automatically selected
		if c.Bind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")

}
