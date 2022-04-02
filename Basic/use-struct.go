package main

import (
	"example.com/go-practice/src/struct_practice"
	"fmt"
)

func main() {
	// 这样初始化是过不了编译的
	//s := struct_practice.Test{1, 2}
	s := new(struct_practice.Test)
	s.A = 1
	s.B = 1
	fmt.Println(s)
}
