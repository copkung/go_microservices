package main

import (
	dbConnect "demo/templates/database"
	User "demo/templates/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getAllData", GetAll)
	router.Run(":8080")
}

func GetAll(c *gin.Context) {
	db := dbConnect.ConnectDB()
	var users []User.User
	db.Table("User").Find(&users)
	c.JSON(200, users)
}
