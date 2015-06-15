package http

import (
	"net/http"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"
	"github.com/gin-gonic/gin"
)

type errorMsg map[string]string

type RequestCreateAccount struct {
	MailAddress string `form:"mailaddress" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

func badReq(c *gin.Context, msg string) {
	if msg == "" {
		msg = "bad request data"
	}
	c.JSON(http.StatusBadRequest, errorMsg{"error": msg})
}

func createAccount(c *gin.Context) {
	var req RequestCreateAccount
	err := c.Bind(&req)
	if err != nil {
		badReq(c, err.Error())
		return
	}
	maddr := helpers.MailAddress(req.MailAddress)
	if !maddr.Valid() {
		badReq(c, "invalid mail address")
		return
	}
	dp, err := maddr.DomainPart()
	if err != nil {
		badReq(c, "invalid mail address")
		return
	}
	pwHash, err := users.HashPassword([]byte(req.Password))
	if err != nil {
		badReq(c, "couldn't hash password")
	}

	if err = users.New(users.Domain{DomainName: dp}, req.MailAddress, pwHash); err != nil {
		badReq(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": req.MailAddress})
}
