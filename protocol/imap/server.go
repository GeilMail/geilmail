package imap

import (
	"bufio"
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
		cmd  []string
	)
	bufread := bufio.NewReader(c)
	rdr := textproto.NewReader(bufread)
	defer c.Close()

	send(c, "* OK IMAP4rev1 Service Ready\n")

	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, err)
		return
	}

	if strings.ToUpper(imsg) == "CAPABILITY" {
		send(c, "* CAPABILITY IMAP4rev1 STARTTLS\n")
		send(c, fmt.Sprintf("%s OK CAPABILITY COMPLETED\n", seq))
	} else {
		sendError(c, "wrong command order")
	}

	// Now we *need* STARTTLS
	seq, imsg, err = receiveInSequence(rdr)
	if err != nil {
		sendError(c, err)
		return
	}
	if imsg != "STARTTLS" {
		sendError(c, fmt.Errorf("Only connections with STARTTLS are supported; for your own safety"))
		return
	}

	c = tls.Server(c, tlsConf)

}

func send(c net.Conn, data string) {
	log.Println("sent: ", data)
	c.Write([]byte(data))
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
