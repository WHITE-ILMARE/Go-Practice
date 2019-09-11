package main;

import "fmt";

func main() {
	var a []int; // a是切片
	var pointer = &a; // pointer -> a
	var num = 7;
	fmt.Println(a);
	fmt.Print(pointer);
	fmt.Print(&a);
	fmt.Scan(&num);
}