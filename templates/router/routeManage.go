package templates

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func RunRoute(c *gin.Engine) {
	router = c
	router.Run(":3000")
}
