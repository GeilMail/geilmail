package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listen(hostport string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/", rootURL)
		api.POST("/account/create", createAccount)
	}

	router.Run(hostport)
}

func rootURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "up and running"})
}

type RequestCreateAccount struct {
	MailAddress string `form:"mailaddress" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

func createAccount(c *gin.Context) {
	//TODO: add create account impl
	var req RequestCreateAccount
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": req.MailAddress})
}
