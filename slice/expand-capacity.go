package main

import "fmt"

func main() {
	slice := make([]int, 1, )
	for i := 0; i < 9; i++ {
		slice = append(slice, 1)
		fmt.Printf("cap(slice1)=%d, &(slice1)=%X\n", cap(slice), &slice[0])
	}
	fmt.Println("可以看到，发生扩容时切片引用的地址发生了变化，即创建了新的底层数组")
}