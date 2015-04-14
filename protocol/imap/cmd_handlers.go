package imap

import (
	"fmt"
	"net"
)

func handleListCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle list")
	send(c, fmt.Sprintf(`* LIST (\Noselect) "/" ""`))
	send(c, fmt.Sprintf("%s OK LIST Completed", seq))
}

func handleCreateCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle create")
}

func handleSelectCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle select")
	send(c, fmt.Sprintf(`* 100 EXISTS`))
	send(c, fmt.Sprintf(`* 1 RECENT`))
	send(c, fmt.Sprintf(`* OK [UNSEEN 12]`))
	send(c, fmt.Sprintf("%s OK [READ-WRITE] SELECT completed", seq))
}

func handleUIDCommand(c net.Conn, seq, msg string) {
	fmt.Println("TODO: implement handle UID")
}
