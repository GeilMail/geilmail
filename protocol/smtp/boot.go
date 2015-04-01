package smtp

import (
	"crypto/tls"
	"log"
	"os"
	"path"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage/mail"
)

var (
	tlsConf     *tls.Config
	mailStorage mail.Storage
)

func Boot(c *cfg.Config) {
	gopath := os.Getenv("GOPATH")
	if string(gopath[len(gopath)-1]) == ":" {
		gopath = gopath[:len(gopath)-1]
	}
	log.Println(gopath)
	cert, err := tls.LoadX509KeyPair(path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.crt"), path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.key")) //TODO: make configurable
	if err != nil {
		panic(err)
	}

	log.Println("Booting SMTP server")
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	go listen(c.SMTP.ListenIP, c.SMTP.Port)
}

func ShutDown() {
	listening = false
}
