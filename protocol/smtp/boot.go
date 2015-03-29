package smtp

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/storage/mail"
)

var (
	tlsConf     *tls.Config
	mailStorage mail.Storage
)

func Boot() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

	log.Println("Booting SMTP server")
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	go listen()
}

func ShutDown() {
	listening = false
}
