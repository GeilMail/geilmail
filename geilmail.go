package main

import (
	"log"

	"github.com/GeilMail/geilmail/protocol/imap"
	"github.com/GeilMail/geilmail/protocol/smtp"

	_ "github.com/facebookgo/ensure" // we need this, so go get will pull it
)

func main() {
	log.Println("Starting GeilMail")

	imap.Boot()
	smtp.Boot()
}
