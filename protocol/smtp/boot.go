package smtp

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/textproto"
)

var (
	tlsConf tls.Config
)

const (
	smtpPort = 1587               //TODO: set to 587 later
	hostName = "mail.example.com" //TODO: make configurable
)

func Boot() {
	log.Println("Booting SMTP server")
	listen()
	// tlsConf := tls.Config{}
}

// listen for plain SMTP connections
func listen() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", smtpPort))
	if err != nil {
		// panic("Could not listen on port 587 for SMTP")
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Could not accept new connection via SMTP")
		}

		go handleIncomingConnection(conn)
	}
}

func handleIncomingConnection(c net.Conn) {
	var (
		err  error
		imsg string
	)
	//TODO: pipeline
	rdr := textproto.NewReader(bufio.NewReader(c))
	defer c.Close()

	// welcome the client
	advMsg := fmt.Sprintf("220 %s SMTP service ready\n", hostName)
	c.Write([]byte(advMsg))

	// read HELO/EHLO
	imsg, err = rdr.ReadLine()
	if err != nil {
		writeError(c)
		return
	}

	// verify HELO/EHLO
	if imsg[:4] != "EHLO" && imsg[:4] != "HELO" {
		writeError(c)
		return
	}
	//TODO: verify host name

	// signalize waiting for content
	welcomeMsg := fmt.Sprintf("250 Hello %s, we are ready over here", "dummy")
	c.Write([]byte(welcomeMsg))

}

func writeError(c net.Conn) {
	c.Write([]byte("500 Error"))
}
