// 获取命令行输入的参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	// for循环的range写法，类似于js的forEach
	// Go对于语法比JS更严格，发明了blank identifier(_)，体验很好
	for _, val := range os.Args {
		fmt.Print(val + ",")
	}
	fmt.Print("\n")
	// 留坑待填，比较join和低效循环方式的运行效率
	fmt.Println(strings.Join(os.Args[1:], ","))
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ","
	}
	fmt.Println(s)
	// 练习分割函数，以=为分隔符，要求输入的命令行参数中有=
	//for ind, val := range os.Args {
	//	if ind == 0 {
	//		continue
	//	}
	//	fmt.Printf("%s = %s\n", strings.Split(val, "=")[0], strings.Split(val, "=")[1])
	//}
}
