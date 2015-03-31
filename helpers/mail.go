package helpers

import (
	"fmt"
	"strings"
)

type MailAddress string

//TODO: implement
func ValidMailAddress(addr MailAddress) bool {
	panic("not implemented yet")
}

func MailDomainPart(addr MailAddress) (string, error) {
	l := strings.Split(string(addr), "@")
	if len(l) != 2 {
		return "", fmt.Errorf("Mail address %s is not valid.", addr)
	}
	return l[1], nil
}
