package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func GetAllUser(c *gin.Context) {

	var testUser User

	testUser.Username = "codingno"
	testUser.Firstname = "codingno"
	testUser.Lastname = "code"

	c.JSON(http.StatusOK, testUser)
}

type Authorize struct {
	Authorization string `header:"Authorization"`
}

func GetAuthorize(c *gin.Context) {
	authorize := Authorize{}

	if err := c.ShouldBindHeader(&authorize); err != nil {
		c.JSON(200, err)
	}

	fmt.Println(authorize)
	c.JSON(200, gin.H{"Authorization": authorize.Authorization})
}

func ApiRouter(r *gin.RouterGroup) {
	rUser := r.Group("/user")
	{
		rUser.GET("", GetAllUser)
		rUser.GET("/", GetAllUser)
	}

	rAuth := r.Group("/auth")
	{
		rAuth.GET("", GetAuthorize)
	}
}
