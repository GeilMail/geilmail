package smtp

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/storage/mail"
)

var (
	tlsConf     tls.Config
	mailStorage mail.Storage
)

func Boot() {
	log.Println("Booting SMTP server")
	go listen()
}

func ShutDown() {
	listening = false
}
