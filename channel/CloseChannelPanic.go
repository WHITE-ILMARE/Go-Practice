package main

import "time"

// 多次关闭同一个channel会不会panic
// 奇怪的是，如果把g1放在g2后面，g1不会执行

func main() {
	ch := make(chan int, 1)
	done := make(chan struct{}, 1)

	// g1
	go func() {
		println("g1-1")
		<-time.After(1 * time.Second)
		println("close1")
		ch <- 1
		println("g1-2")
		close(ch)
		println("g1-3")
	}()

	// g2
	func() {
		println("g2-1")
		<-time.After(2 * time.Second)
		println("close2")
		close(ch)
		println("g2-2")
		close(done)
		println("g2-3")
	}()

	<-done
	println("done received")
}
