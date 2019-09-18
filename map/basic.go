// map的基础知识
package main

func main() {
	map1 := map[string]int {
		"firstName": 1,
		"lastName": 2,
	}
	addr := &map1["firstName"] // 不能对map中的项取址，因为map可能会因为扩容而重新分配内存空间，导致之前的地址失效
}
