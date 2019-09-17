// 使用new声明变量
// 编译器会自动选择在栈上还是在堆上分配局部变量的空间，而不是由用var还是new声明决定的
package main

import "fmt"

var global *int

func main() {
	p := new(int)
	fmt.Printf("p=%v, *p=%v \n", p, *p)
	fmt.Println(delta(3, 1))
	escapeVar()
}

// new并不是关键字,若是重新声明了new，则内部无法使用new函数
func delta(old, new int) int {
	return new - old
}

// x在函数返回后依然存在，地址应该被分配在堆上
func escapeVar() {
	x := 1
	global = &x
	fmt.Println(global)
}