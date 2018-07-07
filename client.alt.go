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
		fmt.Print("Input: ")
		input, _ := r.ReadString('\n')
		fmt.Fprintf(conn, input+"\n")
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Output: " + msg)
		if strings.Contains(msg, "Bye") {
			break
		}
	}
}
