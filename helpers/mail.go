package helpers

import (
	"fmt"
	"net/mail"
	"strings"
)

type MailAddress string

func (addr MailAddress) Valid() bool {
	if _, err := mail.ParseAddress(string(addr)); err != nil {
		return false
	}
	return true
}

func (addr MailAddress) DomainPart() (string, error) {
	l := strings.Split(string(addr), "@")
	if len(l) != 2 {
		return "", fmt.Errorf("Mail address %s is not valid.", addr)
	}
	return l[1], nil
}
