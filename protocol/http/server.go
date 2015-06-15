package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listen(hostport string) {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("", rootURL)
		api.POST("/account/create", createAccount)
	}

	router.Run(hostport)
}

func rootURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "up and running"})
}
