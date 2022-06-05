package main

import (
	"fmt"
)

// 父子协程通过data和done通信，data传递数据，done标识是否完成了计算
// 输出顺序不定
func main() {
	data := make(chan int)
	done := make(chan bool)

	go func() {
		// 用range可以判断channel是否关闭，若未关闭是出不了range的
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recvive over")
		done <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)

	fmt.Println("send over.")
	<-done
}
