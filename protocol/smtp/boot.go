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
	smtpPort        = 1587               //TODO: set to 587 later
	hostName        = "mail.example.com" //TODO: make configurable
	errMsgBadSyntax = "message not understood"
	maxReceivers    = 10
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
		imsg string // incoming message
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
		writeError(c, errMsgBadSyntax)
		return
	}

	// verify HELO/EHLO
	if len(imsg) < 4 || (imsg[:4] != "EHLO" && imsg[:4] != "HELO") {
		writeError(c, "invalid EHLO")
		return
	}
	//TODO: verify host name

	// signalize waiting for content
	welcomeMsg := fmt.Sprintf("250 Hello %s, we are ready over here\n", "dummy")
	c.Write([]byte(welcomeMsg))

	// now we expect a MAIL FROM:
	imsg, err = rdr.ReadLine()
	if err != nil {
		writeError(c, errMsgBadSyntax)
		return
	}
	if len(imsg) < 10 || imsg[:10] != "MAIL FROM:" {
		writeError(c, "invalid MAIL FROM message")
		return
	}
	fromAddr := imsg[10:] //TODO: forgive whitespaces
	if string(fromAddr[0]) != "<" || string(fromAddr[len(fromAddr)-1]) != ">" {
		writeError(c, "invalid MAIL FROM address")
		return
	}
	okMsg(c)

	// read receivers (RCPT TO) or wait DATA to start
	receivers := []string{}
	for {
		if len(receivers) > maxReceivers {
			writeError(c, "too many receivers")
			return
		}
		imsg, err = rdr.ReadLine()
		if err != nil {
			writeError(c, errMsgBadSyntax)
			return
		}
		// right now, we can receive another RCPT or an DATA statement
		if imsg[:4] == "DATA" {
			break
		}
		if len(imsg) < 8 || imsg[:8] != "RCPT TO:" {
			writeError(c, "invalid RCPT TO message")
			return
		}
		rcptAddr := imsg[8:] //TODO: forgive whitespaces
		if string(rcptAddr[0]) != "<" || string(rcptAddr[len(rcptAddr)-1]) != ">" {
			writeError(c, "invalid RCPT TO address")
			return
		}
		//TODO: mail address validation
		rcptAddr = rcptAddr[1 : len(rcptAddr)-1]
		receivers = append(receivers, rcptAddr)
		okMsg(c)
	}

	c.Write([]byte("354 End data with <CR><LF>.<CR><LF>\n"))

	for {
		//TODO: implement data limit
		imsg, err = rdr.ReadLine()
		if err != nil {
			writeError(c, errMsgBadSyntax)
		}
		//TODO:<CR><LF>.<CR><LF>
		if imsg == "." {
			// schedule mail for delivery
			c.Write([]byte("250 Ok: queued as 1337\n")) //TODO queue id
			log.Println("Received message")
			return
		}
	}
}

func writeError(c net.Conn, msg string) {
	c.Write([]byte(fmt.Sprintf("500 %s\n", msg)))
}

func okMsg(c net.Conn) {
	c.Write([]byte("250 Ok\n"))
}
