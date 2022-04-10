package selfctx

import "sync"

var closedchan = make(chan struct{})
var wg sync.WaitGroup

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
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *SimpleContex) Cancel() {
	c.mu.Lock()
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	c.mu.Unlock()
}
