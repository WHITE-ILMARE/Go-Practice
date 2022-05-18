package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

func main() {
	var (
		maxWorkers = runtime.GOMAXPROCS(0)          // worker数量
		sem        = NewWeighted(int64(maxWorkers)) // 信号量，size是worker数
		out        = make([]int, 32)                // 任务数
	)
	ctx := context.TODO()
	// 这儿的信号量其实是虚指，n可以随便设置，与goroutinue真实使用的线程数没关系，只是在信号量层面的限制而已
	// 但n不能超过12，因为我的电脑是6核12线程的
	for i := range out {
		if err := sem.Acquire(ctx, 3); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		go func(i int) {
			defer sem.Release(3)
			out[i] = collatzSteps(i + 1)
		}(i)
	}
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}
	fmt.Println(out)
}

func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}
	for ; n > 1; steps++ {
		if steps < 0 {
			panic("too many steps")
		}
		if n%2 == 0 {
			n /= 2
			continue
		}
		const maxInt = int(^uint(0) >> 1)
		if n > (maxInt-1)/3 {
			panic("overflow")
		}
		n = 3*n + 1
	}
	return steps
}
