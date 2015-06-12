package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/protocol/http"
	"github.com/GeilMail/geilmail/protocol/imap"
	"github.com/GeilMail/geilmail/protocol/smtp"
	"github.com/GeilMail/geilmail/storage"

	_ "github.com/facebookgo/ensure" // we need this, so go get will pull it
)

func main() {
	log.Println("Starting GeilMail")

	gopath := os.Getenv("GOPATH")
	if string(gopath[len(gopath)-1]) == ":" {
		gopath = gopath[:len(gopath)-1]
	}

	conf := cfg.ReadConfig("config.yaml")
	//TODO: read config from file
	// conf := &cfg.Config{
	// 	SQLite: &cfg.SQLiteConfig{DBPath: "geilmail.db"},
	// 	IMAP: &cfg.IMAPConfig{
	// 		ListenIP: "0.0.0.0",
	// 		Port:     1143, //TODO: set to 143
	// 	},
	// 	SMTP: &cfg.SMTPConfig{
	// 		ListenIP: "0.0.0.0",
	// 		Port:     1587, //TODO: set to 587
	// 	},
	// 	TLS: &cfg.TLSConfig{
	// 		CertPath: path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.crt"),
	// 		KeyPath:  path.Join(gopath, "src/github.com/GeilMail/geilmail/certs/server.key"),
	// 	},
	// 	HTTP: &cfg.HTTPConfig{
	// 		Listen: "0.0.0.0:51488",
	// 	},
	// }

	go imap.Boot(conf)
	go smtp.Boot(conf)
	go http.Boot(conf)
	go storage.Boot(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c // blocks until interrupt
}
