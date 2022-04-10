package main

import (
	"example.com/go-practice/src/selfctx"
	"fmt"
	"sync"
	"time"
)

// 在context中，通过关闭done这个channel传递关闭信号
// 结果deadlock了，说明close的channel不会向外发送消息

// 补充：后来查资料发现要求子goroutine中select时要带default，不然就会造成上面的死锁
// 而且closedchannel可以声明如下：
var wg sync.WaitGroup

func main() {
	ctx := selfctx.NewSimpleContex()
	wg.Add(1)
	go work(&ctx)
	time.Sleep(time.Second * 4)
	ctx.Cancel()
	wg.Wait()
	fmt.Println("main等到了wg，退出")
}

func work(ctx *selfctx.SimpleContex) {
	ticker := time.NewTicker(time.Second)
	count := 0
	defer wg.Done()
	for _ = range ticker.C {
		fmt.Printf("worker is working: %d\n", count)
		count++
		select {
		case <-ctx.Done():
			fmt.Println("关闭上下文，退出worker")
			return
		default: // 必须有default，因为没有channel的输入，只有接受和关闭，不设置default会导致deadlock
		}
	}
}
