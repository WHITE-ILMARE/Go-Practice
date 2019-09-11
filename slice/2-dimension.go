package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]int {
	var result [][]int;
	item := make([]int, 0, dx);
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			//fmt.Printf("i=%d,j=%d\n", i, j);
			item = append(item, i + j);
		}
		result = append(result, item);
		item = make([]int, 0, dx);
	}
	return result;
}

func main() {
	dx, dy := 3, 3;
	for line := 0; line < dy; line++ {
		fmt.Println(Pic(dx, dy)[line]);
	}
}

