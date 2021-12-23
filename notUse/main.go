package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	router  = gin.Default()
	dsn     = "root:mysql@tcp(127.0.0.1:3306)/usersql"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)

func main() {
	// routeMapping()
	// routeMap.routeMapping()
	// fmt.Println("Hello!!!")
	router.Run(":3000")
}
