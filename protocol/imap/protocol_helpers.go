package imap

import (
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
)

func send(c net.Conn, data string) {
	log.Println("sent:", data)
	c.Write([]byte(data + "\r\n"))
}

func seqSend(c net.Conn, seq, data string) {
	send(c, fmt.Sprintf("%s %s", seq, data))
}

func sendError(c net.Conn, err string) {
	log.Println("sent errormsg:", err)
	c.Write([]byte(err + "\r\n"))
	c.Close()
}

func seqSendError(c net.Conn, seq, err string) {
	sendError(c, fmt.Sprintf("%s %s", seq, err))
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
