package imap

import (
	"log"

	"github.com/GeilMail/geilmail/cfg"
)

func Boot(c *cfg.Config) {
	log.Println("Booting IMAP server")
	go listen(c.IMAP.ListenIP, c.IMAP.Port)
}
