package imap

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
)

var listening = true

func listen(listenHost string, listenPort int) {
	ln, err := net.Listen("tcp", fmt.Sprintf("%v:%v", listenHost, listenPort))
	if err != nil {
		panic(err)
	}

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
	log.Printf("New connection with %v\n", c.RemoteAddr())
	var (
		err  error
		seq  string // sequence ID
		imsg string // incoming message
	)
	rdr := textproto.NewReader(bufio.NewReader(c))
	defer c.Close()

	send(c, "* OK IMAP4rev1 Service Ready")

	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, fmt.Sprintf("%s", err))
		return
	}

	if strings.ToUpper(imsg) == "CAPABILITY" {
		send(c, "* CAPABILITY IMAP4rev1 STARTTLS")
		send(c, fmt.Sprintf("%s OK CAPABILITY COMPLETED", seq))
	} else {
		sendError(c, "wrong command order")
	}

	// Now we *need* STARTTLS
	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, fmt.Sprintf("%s", err))
		return
	}
	if imsg != "STARTTLS" {
		sendError(c, "Only connections with STARTTLS are supported; for your own safety")
		return
	}

	send(c, fmt.Sprintf("%s OK Starting TLS", seq))

	c = tls.Server(c, tlsConf)
	rdr = textproto.NewReader(bufio.NewReader(c))
	log.Println("info: STARTTLS successful")

	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, "capability problem")
		return
	}
	if strings.ToUpper(imsg) == "CAPABILITY" {
		send(c, "* CAPABILITY IMAP4rev1 AUTH=PLAIN")
		send(c, fmt.Sprintf("%s OK CAPABILITY completed", seq))
	} else {
		sendError(c, "at this moment I want to be asked about my CAPABILITY")
		return
	}

	// LOGIN
	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, "Login problem")
		return
	}
	magicWord := strings.ToUpper(strings.Split(imsg, " ")[0])
	if magicWord == "AUTHENTICATE" {
		send(c, fmt.Sprintf("%s NO please LOGIN", seq))
		seq, imsg, err = receiveInSequence(rdr)
		if err != nil {
			sendError(c, "i am confused")
		}
		magicWord = strings.ToUpper(strings.Split(imsg, " ")[0])
	}
	if magicWord == "LOGIN" {
		send(c, fmt.Sprintf("%s OK", seq))
	} else {
		sendError(c, fmt.Sprintf("%s BAD no idea what you intended to do", seq))
		return
	}

}

func send(c net.Conn, data string) {
	log.Println("sent: ", data)
	c.Write([]byte(data + "\n"))
}

func sendError(c net.Conn, err string) {
	log.Println("sent errormsg: ", err)
	c.Write([]byte(err))
	c.Close()
}

func receive(r *textproto.Reader) (string, error) {
	s, err := r.ReadLine()
	if err != nil {
		log.Println("read: [ERROR]")
		return "", err
	}
	log.Printf("read: %v", s)
	return s, nil
}

func receiveInSequence(r *textproto.Reader) (string, string, error) {
	imsg, err := receive(r)
	if err != nil {
		return "", "", err
	}

	sl := strings.SplitN(imsg, " ", 2)
	if len(sl) != 2 {
		return "", "", fmt.Errorf("Command not in sequence: %s", imsg)
	}
	return sl[0], sl[1], nil
}
