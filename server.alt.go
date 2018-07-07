package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		if strings.Contains(msg, "exit") {
			conn.Write([]byte("Bye"))
			break
		}
		fmt.Print("Recieved: ", string(msg))
		conn.Write([]byte(msg + "\n"))
	}
}
