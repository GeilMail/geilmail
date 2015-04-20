package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listen(hostport string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api")
	api.GET("/", rootURL)
	api.GET("/mail/add", addMailAddr)

	router.Run(hostport)
}

func rootURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "up and running"})
}

func addMailAddr(c *gin.Context) {
	//TODO: add create account impl
	c.JSON(http.StatusOK, gin.H{"message": "TODO"})
}
