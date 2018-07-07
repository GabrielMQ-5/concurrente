package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ls, _ := net.Listen("tcp", "localhost:8000")
	defer ls.Close()
	conn, _ := ls.Accept()
	defer conn.Close()
	for {
		r := bufio.NewReader(conn)
		msg, _ := r.ReadString('\n')
		msg = strings.Trim(msg, " ")
		if strings.Contains(msg, "quit") {
			break
		}
		fmt.Printf("Client: %s", msg)
		fmt.Fprintf(conn, "%s", msg)
	}
}
