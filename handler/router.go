package handler

import (
	"github.com/gin-gonic/gin"
)

// InitializeAPI start liveliness endpoint
func InitializeAPI() *gin.Engine {
	webEngine := gin.Default()
	webEngine.GET("/liveliness", LivelinessHandler())
	_ = webEngine.Run("localhost:12345")
	return webEngine
}
