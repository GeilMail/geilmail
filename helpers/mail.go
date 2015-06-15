package helpers

import (
	"fmt"
	"log"
	"strings"
)

type MailAddress string

//TODO: implement
func (addr MailAddress) Valid() bool {
	log.Println("implement mail address validation!")
	return true
}

func (addr MailAddress) DomainPart() (string, error) {
	l := strings.Split(string(addr), "@")
	if len(l) != 2 {
		return "", fmt.Errorf("Mail address %s is not valid.", addr)
	}
	return l[1], nil
}
