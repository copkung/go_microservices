package templates

type User struct {
	ID      int    `"json:id"`
	Name    string `"json:name"`
	Message string `"json:message"`
}

// var (
// 	router  = gin.Default()
// 	dsn     = "root:mysql@tcp(127.0.0.1:3306)/usersql"
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// )
