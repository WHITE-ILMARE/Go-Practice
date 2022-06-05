// 进一步理解channel未满时的非阻塞状态
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 通知本次填值结束，不应再从此channel中读值
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	go func() { fmt.Println("loop start") }()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("loop end")
}
