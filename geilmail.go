package main

import (
	"log"

	"github.com/GeilMail/geilmail/protocol/imap"
	"github.com/GeilMail/geilmail/protocol/smtp"
)

func main() {
	log.Println("Starting GeilMail")

	imap.Boot()
	smtp.Boot()
}
