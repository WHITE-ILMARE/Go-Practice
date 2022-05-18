/**
这个例子演示了加锁后程序按预期运行
*/
package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

//var Cnt int  Cnt不再重复声明
var mu sync.Mutex

func MuAdd(iter int) {
	mu.Lock()
	for i := 0; i < iter; i++ {
		Cnt++
	}
	mu.Unlock()
}

func main() {
	f, _ := os.Create("mutex/trace.dat")
	trace.Start(f)
	defer trace.Stop()
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			MuAdd(100000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(Cnt)
}
