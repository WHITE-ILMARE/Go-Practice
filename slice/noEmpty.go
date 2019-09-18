// 利用slice在原数组上排除空字符串
package main

import "fmt"

func main() {
	origin := []string{"1", "2", "", "4"}
	fmt.Println(noempty(origin))
}

func noempty(strings []string) []string {
	fmt.Println("IN ARRAY:")
	for index, item := range strings {
		fmt.Printf("item %d = %X, value = %s\n", index, &strings[index], item)
	}
	result := strings[:0]
	for _, s := range strings {
		if (s != "") {
			result = append(result, s) // Slice只是引用，操作的还是原数组
		}
	}
	fmt.Println("IN ARRAY:")
	for index, item := range strings {
		fmt.Printf("item %d = %X, value = %s\n", index, &strings[index], item)
	}
	fmt.Println("IN SLICE:")
	for index, item := range result {
		// 此处如果打印&item,则地址都是一样的，都为item变量的地址
		fmt.Printf("item %d = %X, value = %s\n", index, &result[index], item)
	}
	return result
}
