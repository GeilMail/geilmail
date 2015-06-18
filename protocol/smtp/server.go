package smtp

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/mail"
)

const (
	errMsgBadSyntax = "message not understood"
	maxReceivers    = 10
)

var (
	listening    = true
	capabilities = []string{"AUTH", "STARTTLS", "LOGIN"}
	hostName     string
)

// listen for plain SMTP connections
func listen(listenHost string, listenPort int, hostName string, rdy chan bool) {
	hostName = hostName
	ln, err := net.Listen("tcp", fmt.Sprintf("%v:%v", listenHost, listenPort))
	if err != nil {
		panic(err)
	}
	// accepting connections now
	rdy <- true
	for {
		if !listening {
			ln.Close()
			break
		}

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
	//TODO: pipeline?
	rdr := textproto.NewReader(bufio.NewReader(c))
	defer c.Close()

	// welcome the client
	advMsg := fmt.Sprintf("220 %s SMTP service ready\n", hostName)
	write(c, advMsg)

	// read HELO/EHLO
	imsg, err = read(rdr)
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
	welcomeMsg := fmt.Sprintf("250-Hello %s, we are ready over here\n250 %s\n", "dummy", strings.Join(capabilities, " ")) //TODO: refactor this and signalize further capabilities, if needed
	write(c, welcomeMsg)

	imsg, err = read(rdr)
	if err != nil {
		writeError(c, errMsgBadSyntax)
		return
	}

	// check if STARTTLS has been annouced
	if imsg == "STARTTLS" {
		write(c, "220 i love encryption\n")
		// overwriting connection and reader for transparent encryption handling
		c = tls.Server(c, tlsConf)
		rdr = textproto.NewReader(bufio.NewReader(c))
		// usually the client will now send EHLO
		imsg, err = read(rdr)
		if len(imsg) >= 4 {
			if imsg[:4] == "EHLO" {
				advMsg := fmt.Sprintf("250 ready\n")
				write(c, advMsg)
			} else {
				writeError(c, "only accepting EHLO at that place")
			}
		}
		// allow continuing with MAIL FROM:
		imsg, err = read(rdr)
	}

	// read MAIL FROM:
	if len(imsg) < 10 || imsg[:10] != "MAIL FROM:" {
		writeError(c, "invalid MAIL FROM message")
		return
	}
	fromAddr := imsg[10:] //TODO: forgive whitespaces
	if string(fromAddr[0]) != "<" || string(fromAddr[len(fromAddr)-1]) != ">" {
		writeError(c, "invalid MAIL FROM address")
		return
	}
	fromAddr = fromAddr[1 : len(fromAddr)-1]
	okMsg(c)

	// read receivers (RCPT TO) or wait DATA to start
	receivers := []helpers.MailAddress{}
	for {
		if len(receivers) > maxReceivers {
			writeError(c, "too many receivers")
			return
		}
		imsg, err = read(rdr)
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
		receivers = append(receivers, helpers.MailAddress(rcptAddr)) //TODO: validate mail address
		okMsg(c)
	}

	write(c, "354 End data with <CR><LF>.<CR><LF>\n")

	mailData, err := rdr.ReadDotBytes()
	if err != nil {
		writeError(c, errMsgBadSyntax)
		return
	}

	write(c, "250 Ok: queued as 1337\n") //TODO queue id
	log.Println("Received message")
	mail.MailDrop(receivers, mailData)

	imsg, err = read(rdr)
	if err != nil {
		writeError(c, errMsgBadSyntax)
		return
	}
	if imsg == "QUIT" {
		write(c, "221 Bye\n")
	}
}
