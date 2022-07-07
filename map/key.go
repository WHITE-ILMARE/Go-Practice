package main

// 实验struct作为map的key和*struct作为map的key时的区别
import "fmt"

type S struct {
	A int
}

func main() {
	s1 := S{3}
	s2 := S{3}

	m1 := make(map[S]int)
	m2 := make(map[*S]int)

	m1[s1] = 3
	m2[&s1] = 4

	fmt.Printf("s1 == s2? : %v\n", s1 == s2)
	fmt.Println(m1[s2])
	fmt.Println(m2[&s2])
}
