package main

import (
	"fmt"
	"time"
)

// 在context中，通过关闭done这个channel传递关闭信号
// 我很疑惑close(channel)可以被select监听到吗，做实验如下
// 结果deadlock了，说明close的channel不会向外发送消息

type simple_ctx struct {
	done chan struct{} // 双向channel
}

// 返回的是单向只读channel
func (c *simple_ctx) Done() <-chan struct{} {
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	return d
}

func (c *simple_ctx) cancel() {
	if c.done == nil {
		temp := make(chan struct{})
		close(temp)
		c.done = temp
	} else {
		close(c.done)
	}
}

func main() {
	ctx := simple_ctx{done: nil}
	go func(ctx simple_ctx) {
		time.Sleep(3 * time.Second)
		ctx.cancel()
	}(ctx)
	select {
	case something := <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println(something)
	}
}
