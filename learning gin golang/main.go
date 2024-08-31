package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//What is context
// This context contains all the information about the request that the handler might need to process it
// its easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request
// :9090/api/path?name=go guru&email=somevalue
// like header cookies etc
// json {} {}
// XML <tag></tag>

//usage of context
// err := c.Bind(&obj)  json and xml
// err := c.BindQuery(&obj)
// err := c.BindXML(&obj)
// err := c.BindYAML(&obj)
// err := c.BindHeader(&obj)
// err := c.BindJSON(&obj)
// c.Header("user-id", "186492719480")
// c.setCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
// val, err := c.Cookie("name")
// err := c.saveUploadFile(file, dest)
// form := c.MultipartForm()
// key := c.PostForm("key", "default value")
// id := c.Query("id")
// id := c.Param("id")
// name := c.DefaultQuery("name", "jack")
// c.GetFloat64("key")
// c.Set("key", "value")
// c.Get("key")
// c.MustGet("key")
// c.GetString("key")
// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
// c.isAborted()
// c.Abort()
// c.AbortWithStatus(http.StatusBadRequest)

// c.XTML(
//   http.StatusOk, "index.html", gin.H{
//     "message": "HTML message",
//     "data": "whateevr can be object etc"
//   })
// )

// c.XML(
//   http.StatusOk, gin.H{
//     "message": "lol"
//   })
// )

// c.YAML(
//   http.StatusOk, gin.H{
//     "message": "lmao"
//   })
// )

// type api struct {
// 	Name  string `json:"name" form:"name" header:"name" binding:"required"`
// 	Email string `json:"email"`
// }

type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var data api

// using thrid party apis in gin
var url = "http://date.jsontest.com/"

func main() {
	router := gin.Default()
	router.GET("/inky", inky)
	router.GET("/get", getValues)
	router.POST("/post", postValues)
	router.PUT("/put", putValues)
	router.DELETE("/delete", deleteValues)
	router.Run(":9090")

	//how to use 3rd party urls

}

func inky(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "inky",
	})
}

func getValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func postValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
