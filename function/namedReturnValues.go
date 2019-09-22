// 命名函数返回值demo
package main

import (
	"fmt"
	html2 "golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		images_out, err_out := CountImages(url)
		if (err_out != nil) {
			// 有错误这里会打印，函数内部的Errorf不会打印
			fmt.Printf("out err: %v\n", err_out)
			continue
		}
		for index, img := range images_out {
			fmt.Printf("%s image %d: %s\n", url, index, img)
		}
	}
}

// 大写的函数名，表示导出的函数
func CountImages(url string) (images []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("getting %s error: %v\n", url, err)
		return
	}
	// 对err再次赋值可以使用:=，前提是左值不能只有err一个变量
	doc, err := html2.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Errorf("parsing html error: %v\n", err)
		return
	}
	images = countImages(nil, doc)
	return
}

// 具体处理node节点的函数
func countImages(init []string, node *html2.Node) (images []string) {
	if node.Type == html2.ElementNode && node.Data == "img" {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				init = append(init, attr.Val)
			}
		}
	}
	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		init = countImages(init, nextNode)
	}
	images = init
	return
}
