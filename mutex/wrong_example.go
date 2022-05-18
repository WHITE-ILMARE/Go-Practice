/**
这个例子演示了两个线程同时访问共享变量Cnt，各自对Cnt加10000次，
最后Cnt的结果不确定
*/
package main

var Cnt int

func Add(iter int) {
	for i := 0; i < iter; i++ {
		Cnt++
	}
}

//func main() {
//	f, _ := os.Create("./trace.dat")
//	trace.Start(f)
//	defer trace.Stop()
//	wg := &sync.WaitGroup{}
//	for i := 0; i < 2; i++ {
//		wg.Add(1)
//		go func() {
//			Add(100000)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println(Cnt)
//}
