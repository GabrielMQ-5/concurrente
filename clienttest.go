package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Msg: ")
		msg, _ := r.ReadString('\n')
		fmt.Fprintf(conn, msg+"\n")
		if strings.Contains(msg, "quit") {
			break
		}
		server_msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server: " + server_msg)
	}
}
