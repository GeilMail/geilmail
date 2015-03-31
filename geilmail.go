package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/GeilMail/geilmail/configuration"
	"github.com/GeilMail/geilmail/protocol/imap"
	"github.com/GeilMail/geilmail/protocol/smtp"
	"github.com/GeilMail/geilmail/storage"

	_ "github.com/facebookgo/ensure" // we need this, so go get will pull it
)

func main() {
	log.Println("Starting GeilMail")

	//TODO: read config from file
	conf := &configuration.Config{
		SQLite: &configuration.SQLiteConfig{DBPath: "geilmail.db"},
	}

	imap.Boot(conf)
	smtp.Boot(conf)
	storage.Boot(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c // blocks until interrupt
}
