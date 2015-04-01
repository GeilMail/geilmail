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
		imsg string // incoming message
		cmd  []string
	)
	bufread := bufio.NewReader(c)
	rdr := textproto.NewReader(bufread)
	defer c.Close()

	c.Write([]byte("* OK IMAP4rev1 Service Ready\n"))

	imsg, err = rdr.ReadLine()
	if err != nil {
		fmt.Println("I didn't like that")
	}
	cmd = strings.Split(imsg, " ")
	if len(cmd) < 2 {
		if strings.ToUpper(cmd[1]) == "CAPABILITY" {
			c.Write([]byte("* CAPABILITY IMAP4rev1 STARTTLS\n"))
		} else {
			panic("hä?")
		}
	}

	fmt.Println(imsg)
}
