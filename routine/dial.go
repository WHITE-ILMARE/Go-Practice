// 每300ms读取tcp连接内容，持续5s
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
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
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
