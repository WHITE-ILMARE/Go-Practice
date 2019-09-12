package main

import (
	"fmt"
)


func main() {
	// 数组的构造
	var basic_arr = [...]int{1,2,3,4,5};
	// Slice的构造
	slice1 := basic_arr[1:3];
	slice2 := basic_arr[0:];
	slice3 := []int{};
	// 值为nil的Slice
	var slice_nil []int;
	slices := [4][]int{slice1, slice2, slice3, slice_nil};
	fmt.Println(slices[0]);
	fmt.Println(slices[1]);
	fmt.Println(slices[2]);
	fmt.Println(slices[3]);
	for i, v := range slices {
		if (v == nil) {
			fmt.Printf("%dth item is nil\n", i + 1);
		}
		// Slice的len和cap属性
		fmt.Printf("cap(%d) = %d, len(%d) = %d\n", i, cap(v), i, len(v));
	}
}
