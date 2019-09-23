// 函数的基本知识
package main

import "fmt"

func main() {
	var nil_func func (arg int) int
	fmt.Printf(" 函数的零值是nil ? %v\n", nil_func == nil)
}
