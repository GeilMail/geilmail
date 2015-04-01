package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/protocol/imap"
	"github.com/GeilMail/geilmail/protocol/smtp"
	"github.com/GeilMail/geilmail/storage"

	_ "github.com/facebookgo/ensure" // we need this, so go get will pull it
)

func main() {
	log.Println("Starting GeilMail")

	//TODO: read config from file
	conf := &cfg.Config{
		SQLite: &cfg.SQLiteConfig{DBPath: "geilmail.db"},
		IMAP: &cfg.IMAPConfig{
			ListenIP: "0.0.0.0",
			Port:     1143, //TODO: set to 143
		},
		SMTP: &cfg.SMTPConfig{
			ListenIP: "0.0.0.0",
			Port:     1587, //TODO: set to 587
		},
	}

	imap.Boot(conf)
	smtp.Boot(conf)
	storage.Boot(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c // blocks until interrupt
}
