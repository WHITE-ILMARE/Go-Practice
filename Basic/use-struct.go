package main

import (
	"fmt"
)

type innerStruct struct {
	C int
}

type innerAnonymous struct {
	D int
}

type testStruct struct {
	A, B  int
	inner innerStruct
	innerAnonymous
}

func main() {
	// 这样初始化是过不了编译的
	//s := testStruct{1, 2}
	s := new(testStruct)
	s.A = 1
	s.B = 1
	// 过不了编译，只有匿名字段才会由外向内查找属性
	//s.C = 2
	// 通过编译，因为是匿名字段
	s.D = 2
	fmt.Println(s)
}
