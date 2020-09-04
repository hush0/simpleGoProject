package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

//User Object
type User struct {
	name    string `form:"name" json:"name" binding:"required"`
	gendary string `form:"gendary" json:"gendary"`
	age     int    `form:"age" json:"age"`
}

//DBconnect
type ConnectionInfo struct {
	MyUser   string
	Password string
	Host     string
	Port     int
	Db       string
}

func main() {

	router := gin.Default()
	//get reauest example
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstName", "Guest")
		lastname := c.Query("lastName")
		c.String(http.StatusOK, "Hello,  %s %s", firstname, lastname)
	})

	//post request example
	router.POST("/user", func(c *gin.Context) {
		name := c.DefaultPostForm("name", "Guest")
		age := c.PostForm("age")
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"name": name,
			"age":  age,
		})
	})

	//form submit example
	router.POST("/addUser", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err = c.BindWith(&user, binding.Form)
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"name":    user.name,
			"gendary": user.gendary,
			"age":     user.age,
		})
	})

	//any router
	router.Any("any", func(c *gin.Context) {
		c.String(200, "any")
	})

	//redirect example
	router.GET("/redict/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//static router
	router.Static("/resources", "./resources")
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("index.html", "./index.html")

	//泛绑定
	router.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello guys")
	})

	/*********************** get params *************************/
	router.GET("/test", func(c *gin.Context) {
		name := c.Query("name")
		message := c.DefaultQuery("message", "hello!")
		c.String(http.StatusOK, "%s,%s", name, message)
	})

	router.Run(":9090")
}
