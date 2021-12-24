package main

import (
	dbConnect "demo/templates/database"

	User "demo/templates/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/postData", postData)
	router.Run(":8080")
}

func postData(c *gin.Context) {
	var input User.User
	c.ShouldBindJSON(&input)
	if input.ID == 0 || input.Name == "" || input.Message == "" {
		c.String(http.StatusBadRequest, "Input incomplete! Please check before posting again.")
	}
	db := dbConnect.ConnectDB()
	if err := db.Table("User").Select("ID", "Name", "Message").Create(&input).Error; err != nil {
		c.String(http.StatusBadRequest, "%v \n", err)
	} else {
		c.JSON(http.StatusCreated, input)
	}
}
