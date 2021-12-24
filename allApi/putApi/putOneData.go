package main

import (
	dbConnect "demo/templates/database"
	User "demo/templates/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.PUT("/updateData", putData)
	router.Run(":8080")
}

func putData(c *gin.Context) {
	var input User.User
	var user User.User
	c.ShouldBindJSON(&input)

	db := dbConnect.ConnectDB()
	db.Table("User").Where("name = ? OR id = ?", input.Name, input.ID).Find(&user)
	if user.Name != "" && user.Message != "" && user.ID != 0 {
		if err := db.Table("User").Select("ID", "Name", "Message").Updates(&User.User{ID: input.ID, Name: input.Name, Message: input.Message}).Error; err != nil {
			c.String(http.StatusBadRequest, "%v \n", err)
		} else {
			c.JSON(http.StatusOK, input)
		}
	} else {
		c.String(http.StatusBadRequest, "Data not found.\n")
	}

}
