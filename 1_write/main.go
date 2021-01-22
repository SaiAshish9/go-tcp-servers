package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// IETF

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	// telnet localhost 8081

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		conn.Close()
	}
}
