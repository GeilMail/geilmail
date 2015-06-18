package smtp

import (
	"fmt"
	"log"
	"net"
	"net/textproto"
)

func writeError(c net.Conn, msg string) error {
	return write(c, fmt.Sprintf("500 %s\n", msg))
}

func okMsg(c net.Conn) {
	write(c, "250 Ok\n")
}

func read(rdr *textproto.Reader) (string, error) {
	msg, err := rdr.ReadLine()
	if err == nil {
		log.Printf("SMTP <<< %s", msg)
	}
	return msg, err
}

func write(c net.Conn, msg string) error {
	_, err := c.Write([]byte(msg))
	if err == nil {
		log.Printf("SMTP >>> %s", msg)
	}
	return err
}
