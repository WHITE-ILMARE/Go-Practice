package main

import (
	"example.com/go-practice/src/selfctx"
	"fmt"
	"sync"
	"time"
)

// 在context中，通过关闭done这个channel传递关闭信号
// 我很疑惑close(channel)可以被select监听到吗，做实验如下
// 结果deadlock了，说明close的channel不会向外发送消息

// 补充：后来查资料发现要求子goroutine中select时要带default，不然就会造成上面的死锁
// 而且closedchannel可以声明如下：
var wg sync.WaitGroup

func main() {
	ctx := selfctx.NewSimpleContex()
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 4)
	ctx.Cancel()
	wg.Wait()
}

func work(ctx selfctx.SimpleContex) {

LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("退出worker")
			break LOOP
		default:
		}
	}
	//ticker := time.NewTicker(time.Second)
	//count := 0
	//for _ = range ticker.C {
	//	fmt.Printf("worker is working: %d\n", count)
	//	count++
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("关闭上下文，退出worker")
	//		break
	//	default: // 必须有default，因为没有channel的输入，只有接受和关闭，不设置default会导致deadlock
	//	}
	//}
	wg.Done()
}
