package smtp

import (
	"crypto/tls"
	"log"
)

var (
	tlsConf tls.Config
)

func Boot() {
	log.Println("Booting SMTP server")
	go listen()
}

func ShutDown() {
	listening = false
}
