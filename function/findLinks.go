// 找到输入的地址指向的HTML文件中的所有超链接，主要练习递归函数的用法
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	fmt.Println("in main")
	doc, err := html.Parse(os.Stdin)
	fmt.Println("hrere")
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err);
		os.Exit(1)
	}
	for index, link := range visit(nil, doc) {
		fmt.Printf("%d link: %s\n", index, link)
	}
}

// 递归地将一个节点及其子节点中的a标签中的超链接提取出来存入links中
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		links = visit(links, nextNode)
	}
	return links
}
