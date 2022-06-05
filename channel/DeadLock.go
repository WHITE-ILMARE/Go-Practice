package main

// 实验channel死锁的情况
func main() {
	// 非缓冲型channel，只有生产/消费者或生产、消费者在同一个goroutine中时会阻塞
	// 非缓冲型channel可以看做"同步模式"，带缓冲的看作"异步模式"
	// 同步模式下，发送者和接受者要同时准备就绪，数据才能在二者间传输，否则，任意一方先发送或接受都会被阻塞直至另一方准备好
	ch := make(chan int)
	ch <- 1

	<-ch
}
