package selfctx

import (
	"fmt"
	"sync"
)

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type SimpleContex struct {
	mu   sync.Mutex
	done chan struct{} // 双向channel
}

func NewSimpleContex() SimpleContex {
	return SimpleContex{}
}

// 返回的是单向只读channel,lazy create
func (c *SimpleContex) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	c.mu.Unlock()
	return c.done
}

func (c *SimpleContex) Cancel() {
	c.mu.Lock()
	if c.done == nil {
		fmt.Printf("in cancel, done is nil\n")
		c.done = closedchan
	} else {
		fmt.Println("in cancel, done is not nil, going to close it")
		close(c.done)
	}
	c.mu.Unlock()
}
