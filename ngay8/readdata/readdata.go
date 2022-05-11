// package main

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fptr := flag.String("fpath", "C:/Code/Golang/training-golang/ngay8/readdata/data.txt", "file path to read from")
// 	flag.Parse()

// 	f, err := os.Open(*fptr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer func() {
// 		if err = f.Close(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		a := strings.Split(s.Text(), " ")
// 		fmt.Println(a)
// 	}
// 	err = s.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
