package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":3002")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		for {
			reader := bufio.NewReader(conn)
			msg, err1 := reader.ReadString('\n')
			fmt.Print(msg)
			if err1 != nil {
				break
			}
		}
	}
}
