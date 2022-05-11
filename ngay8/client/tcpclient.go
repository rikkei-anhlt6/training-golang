package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

// func sendData(msg string) {

// }
func main() {
	connection, err := net.Dial("tcp", "localhost:3002")
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string)
	go readdata(ch)
	for msg := range ch {
		msg = fmt.Sprintf("%s\n", msg[:len(msg)+0])
		connection.Write(([]byte(msg)))
	}
	connection.Close()
}
func readdata(data chan string) {
	fptr := flag.String("fpath", "C:/Code/Golang/training-golang/ngay8/readdata/data.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		data <- s.Text()
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
