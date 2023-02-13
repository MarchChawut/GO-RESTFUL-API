package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

// /////////////////////////////////
func GetHelloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

///////////////////////////////////

func PostData(c *gin.Context) {
	var u User
	c.BindJSON(&u)//context+parser+address
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func GetUsername(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

// update ทุกfields
func PutUpdate(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

func DeleteUser(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func main() {
	r := gin.Default()

	////////////////////////////
	r.GET("/", GetHelloworld)
	////////////////////////////
	r.POST("/register", PostData)
	r.GET("/userinfo", GetUsername)
	r.PUT("/updates/:username/:password", PutUpdate)
	r.DELETE("/deluser", DeleteUser)

	//run server
	r.Run()
}
