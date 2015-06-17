package imap

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"
)

func handleCapability(c net.Conn, seq, msg string) {
	send(c, "* CAPABILITY IMAP4rev1 STARTTLS")
	seqSend(c, seq, "OK CAPABILITY COMPLETED")
}

func handleListCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle list")
	send(c, fmt.Sprintf(`* LIST (\Noselect) "/" ""`))
	seqSend(c, seq, "OK LIST Completed")
}

func handleCreateCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle create")
}

func handleSelectCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle select")
	send(c, fmt.Sprintf(`* 100 EXISTS`))
	send(c, fmt.Sprintf(`* 1 RECENT`))
	send(c, fmt.Sprintf(`* OK [UNSEEN 12]`))
	seqSend(c, seq, "OK [READ-WRITE] SELECT completed")
}

func handleUIDCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle UID")
}

func handleLoginCommand(c net.Conn, seq, msg string) error {
	args := strings.Split(msg, " ")
	if len(args) < 3 {
		return errors.New("not enough arguments")
	}
	username := args[1]
	password := args[2]
	for _, s := range []*string{&username, &password} {
		*s = helpers.UnquoteIfNeeded(*s, '"')
	}

	log.Println(username, password)
	if users.CheckPassword(helpers.MailAddress(username), []byte(password)) {
		seqSend(c, seq, "OK LOGIN")
		return nil
	} else {
		seqSendError(c, seq, "NO invalid login")
		return errors.New("invalid login")
	}

}
