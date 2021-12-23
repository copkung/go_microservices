package main

import (
	dbConnect "demo/templates/database"
	User "demo/templates/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.DELETE("/delData", delData)
	router.Run(":8080")
}

func delData(c *gin.Context) {
	var input User.User
	var user User.User
	c.ShouldBindJSON(&input)
	// getOneApi.GetOne(c, input)
	// replace with available func. later

	db := dbConnect.ConnectDB()
	db.Table("User").Where("name = ? OR id = ?", input.Name, input.ID).Find(&user)
	if user.Name != "" && user.ID == 0 {
		if err := db.Table("User").Select("ID", "Name", "Message").Delete(user).Error; err != nil {
			c.String(http.StatusBadRequest, "%v \n", err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.String(http.StatusBadRequest, "Data not found.\n")
	}

}
