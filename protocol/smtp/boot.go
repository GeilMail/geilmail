package smtp

import (
	"crypto/tls"
	"log"
	"os"
	"path"

	"github.com/GeilMail/geilmail/storage/mail"
)

var (
	tlsConf     *tls.Config
	mailStorage mail.Storage
)

func Boot() {
	gopath := os.Getenv("GOPATH")
	cert, err := tls.LoadX509KeyPair(path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.crt"), path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.key")) //TODO: make configurable
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
