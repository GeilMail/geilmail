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
	conf := cfg.ReadConfig("config.yaml")

	go imap.Boot(conf)
	go smtp.Boot(conf)
	go http.Boot(conf)
	go storage.Boot(conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c // blocks until interrupt
}
