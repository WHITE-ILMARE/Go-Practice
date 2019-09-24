// goroutine基础例子
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	fmt.Printf("fibonacci 45 = %d\n", fibonacci(45))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func spinner(delay time.Duration) {
	for {
		for _, char := range `-\|/` {
			fmt.Printf("\r%c", char)
			time.Sleep(delay)
		}
	}
}
