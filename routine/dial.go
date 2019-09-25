package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go connectToHost("localhost:8000")
	time.Sleep(5 * time.Second)
}

func connectToHost(host string) {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Printf("status = %v\n", status)
		time.Sleep(500 * time.Millisecond)
	}
}
