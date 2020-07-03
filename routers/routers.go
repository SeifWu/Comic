package routers

import (
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
