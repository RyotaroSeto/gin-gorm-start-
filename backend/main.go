package main

import (
	"github.com/gin-gonic/gin"
	"golang-mysql/handler"
)

func main() {
	r := gin.Default()
	handler.HealthCheck(r)
	r.Run()
}
