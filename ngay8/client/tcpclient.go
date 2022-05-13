package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:3003")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {
		msg := s.Text() + "\n"
		fmt.Print(msg)
		connection.Write([]byte(msg))
		time.Sleep(1 * time.Second)
	}
	connection.Close()
}
