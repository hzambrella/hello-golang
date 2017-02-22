package main

import (
	"fmt"
	"log"
	"net/http"

	"engine/datastore"

	"github.com/gin-gonic/gin"
)

var data datastore.Data

func main() {
	d, err := datastore.ParseDataFromFile("test")
	if err != nil {
		log.Println(err)
	}
	data = d
	fmt.Println("----------")
	fmt.Println(data)
	router := gin.Default()
	router.GET("/test/view", view)
	router.GET("/test/add", add)
	router.GET("/test/modify", modify)
	router.GET("/test/delete", move)
	router.Run(":8080")
}

func view(c *gin.Context) {
	var result []byte
	// key := c.Param("key")

	key := c.Query("key")
	if len(key) == 0 {
		result = data.ToJson()
		fmt.Println(data)
		fmt.Println("-----1-----")
		c.String(http.StatusOK, string(result))
		return
	}

	data := data.Get(key)
	if len(data) == 0 {
		c.String(404, "key not found")
		return
	}
	fmt.Println(data)
	fmt.Println("-----2-----")
	c.String(http.StatusOK, key+":"+data)
}

func add(c *gin.Context) {
	var result []byte
	// key := c.Param("key")

	key := c.Query("key")
	if len(key) == 0 {
		c.String(400, "need key")
		return
	}

	for i := range data {
		if key == i {
			c.String(400, "key is exist")
			return
		}
	}
	value := c.Query("data")
	data.Set(key, value)

	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		c.String(400, string(result))
		return
	}

	c.String(200, "add is ok")
}

func modify(c *gin.Context) {
	var result []byte
	var flag bool = false
	// key := c.Param("key")

	key := c.Query("key")
	if len(key) == 0 {
		c.String(400, "need key")
		return
	}

	for i := range data {
		if key == i {
			flag = true
			break
		}
	}
	if !flag {
		c.String(404, "key not found")
		return
	}

	value := c.Query("data")
	data.Set(key, value)

	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		c.String(400, string(result))
		return
	}

	c.String(200, "modify is ok")
}

func move(c *gin.Context) {
	var flag bool = false
	key := c.Query("key")
	for i := range data {
		if key == i {
			flag = true
		}
		if !flag {
			c.String(404, "key not found")
			return
		}
		delete(data, key)
		if err := datastore.SaveDataToFile("test", data); err != nil {
			result := []byte(err.Error())
			c.String(400, string(result))
			return
		}
		c.String(200, "delete is ok")
	}
}
