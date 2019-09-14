// 获取命令行指定的文件中重复行的列表
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 此处若用var，则声明的就是个nil了
	result := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, result)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Errorf("open file%s filed: %v\n", file, err)
				continue
			}
			countLines(f, result)
			f.Close()
		}
		for text, count := range result {
			if count > 1 {
				fmt.Printf("%d\t%s\n", count, text)
			}
		}
	}
}

func countLines(f *os.File, result map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		result[input.Text()]++
	}
}
