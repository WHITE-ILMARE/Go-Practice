package main

import (
	"fmt"
	"time"
)

// 从close了的channel仍然可以读取数据，都是零值

func main() {
	closedchan := make(chan int)
	close(closedchan)
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		select {
		case val := <-closedchan:
			fmt.Println(val)
		}
	}
}
