package http

import (
	"log"

	"github.com/GeilMail/geilmail/cfg"
)

func Boot(c *cfg.Config) {
	log.Println("Booting HTTP server")
	go listen(c.HTTP.Listen)
}

func ShutDown() {
}
