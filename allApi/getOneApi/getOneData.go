package main

import (
	dbConnect "demo/templates/database"
	User "demo/templates/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getData", getOne)
	router.Run(":8080")
}

func getOne(c *gin.Context) {
	var input User.User
	c.ShouldBindJSON(&input)
	db := dbConnect.ConnectDB()
	db.Table("User").Where("name = ? OR id = ?", input.Name, input.ID).Find(&input)
	if input.Name != "" || input.ID != 0 || input.Message != "" {
		c.String(http.StatusOK, "Your target data : %v \n", input)
	} else {
		c.String(http.StatusBadRequest, "Data not found.\n")
	}

}
