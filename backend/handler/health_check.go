package handler

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(check *gin.Engine) {
	check.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
