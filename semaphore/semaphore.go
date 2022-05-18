package main

import (
	"container/list"
	"context"
	"sync"
)

type Weighted struct {
	size    int64
	cur     int64
	mu      sync.Mutex
	waiters list.List
}

func NewWeighted(size int64) Weighted {
	return Weighted{size: size}
}

type Waiter struct {
	n     int64
	ready chan<- struct{}
}

// Acquire P操作
// 细节1：不能defer s.mu.Unlock()，因为此方法会阻塞，阻塞时不应该持有锁
// 细节2：当n>s.size时，这里的处理是unlock并等待外部取消信号，也就是等这个线程的父线程判断超时然后取消，如果父线程不去检测，那这个线程就会阻塞，在测试例子中，设置n为13就会导致所有线程阻塞，也就是死锁
func (s *Weighted) Acquire(ctx context.Context, n int64) error {
	s.mu.Lock()
	if s.size-s.cur >= n && s.waiters.Len() == 0 {
		s.cur += n
		s.mu.Unlock()
		return nil
	}
	if n > s.size {
		s.mu.Unlock()
		<-ctx.Done() // 等待外部取消信号
		return ctx.Err()
	}
	// 暂时资源不足
	ready := make(chan struct{})
	w := Waiter{n: n, ready: ready}
	elem := s.waiters.PushBack(w)
	s.mu.Unlock()
	// 陷入阻塞
	select {
	case <-ctx.Done():
		err := ctx.Err()
		s.mu.Lock()
		select {
		case <-ready:
			err = nil
		default:
			isFront := s.waiters.Front() == elem
			s.waiters.Remove(elem)
			if isFront && s.size > s.cur {
				s.notifyWaiters()
			}
		}
		s.mu.Unlock()
		return err
	case <-ready:
		return nil
	}
}

// notifyWaiters 不是通知一个waiter，而是尽可能多地通知能运行的waiters
// 所以要for循环处理
func (s *Weighted) notifyWaiters() {
	for {
		next := s.waiters.Front()
		if next == nil {
			break
		}
		w := next.Value.(Waiter) // next是Element类型，Value是any类型，也就是interface{}
		if s.size-s.cur < w.n {
			break
		}
		s.cur += w.n
		s.waiters.Remove(next)
		close(w.ready) // 通知已经拿到资源了
	}
}

// TryAcquire 非阻塞获取资源
func (s *Weighted) TryAcquire(n int64) bool {
	s.mu.Lock()
	success := s.size-s.cur >= n && s.waiters.Len() == 0
	if success {
		s.cur += n
	}
	s.mu.Unlock()
	return success
}

// Release V操作
func (s *Weighted) Release(n int64) {
	s.mu.Lock()
	s.cur -= n
	if s.cur < 0 { // n传得不对
		s.mu.Unlock()
		panic("semaphore: released more than held")
	}
	s.notifyWaiters()
	s.mu.Unlock()
}
