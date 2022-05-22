// 测试不导出的匿名成员是否能访问
// 结论：同包下可以访问
package main

import "fmt"

type point struct {
	x int
	Y int
}

type Circle struct {
	point
	Radius int
}

func main() {
	c := Circle{
		point{3, 5},
		4,
	}
	fmt.Printf("c = %#v\n", c)
	fmt.Printf("c.x = %d, c.point.x=%d\n", c.x, c.point.x)
	fmt.Printf("c.Y = %d\n", c.Y)
}
