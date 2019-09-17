// 自定义类型
// go的类型系统提供了语义上的清晰
package main

import "fmt"

type int1 int
type int2 int

func main() {
	var var1 int1 = 3
	var var2 int2 = 3
	fmt.Println(var1)
	fmt.Println(var2)
	// 这里直接比较会报错，因为var1 和 var2不是同一个类型，不可互相赋值，也就不可比较
	//fmt.Println(var1 == var2)
	// 进行底层值的比较是可以的
	fmt.Println(var1 == 3)
	// 不改变值的类型转换也是可以的
	fmt.Println(var1 == int1(var2))
}
