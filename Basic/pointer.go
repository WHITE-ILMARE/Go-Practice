// 指针的基本用法
package main

import "fmt"

func main() {
	var a int;
	a = 1;
	arr := [5]int{1, 3, 5, 7 ,9};
	var pointer *int;
	var pointer_arr *[5]int; // 数组的长度也是其类型标识之一
	pointer = &a;
	fmt.Printf("a=%d\n", a);
	// I saw my PC's memory has 80bits and couldn't find the reason
	fmt.Println(pointer);
	fmt.Println(arr);
	pointer_arr = &arr;
	fmt.Println(pointer_arr);
}