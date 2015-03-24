package smtp

import (
	"crypto/tls"
	"log"
)

var (
	tlsConf tls.Config
)

const (
	smtpPort        = 1587               //TODO: set to 587 later
	hostName        = "mail.example.com" //TODO: make configurable
	errMsgBadSyntax = "message not understood"
	maxReceivers    = 10
)

func Boot() {
	log.Println("Booting SMTP server")
	listen()
	// tlsConf := tls.Config{}
}
