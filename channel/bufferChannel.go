// 缓冲channel
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c) // 这里读取channel，释放了一个空间，所以下一步可以继续填
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
