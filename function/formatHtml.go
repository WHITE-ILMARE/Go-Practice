// 格式化输入抓取到的dom结构，这样递归很容易stack overflow
// 用队列来实现功能相同的函数，见formatHtml_optimized.go
package main

import (
	"golang.org/x/net/html"
	"fmt"
	"log"
	"net/http"
	"os"
)

var indent = 0

func main() {
	for _, url := range os.Args[1:] {
		doc, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			continue
		}
		node, err := html.Parse(doc.Body)
		doc.Body.Close()
		if err != nil {
			log.Fatal(err)
			continue
		}
		eachNode(node, doBefore, doAfter)
	}
}

// doc->DOM节点，prev,post->在遍历doc之前/后的处理程序
func eachNode(doc *html.Node, prev, post func(elm *html.Node)) {
	if prev != nil {
		prev(doc)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		eachNode(doc, prev, post)
	}
	if post != nil {
		post(doc)
	}
}

func doBefore(node *html.Node) {
	if node != nil && node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", indent, "", node.Data)
		indent += 2
	}
}

func doAfter(node *html.Node) {
	if node != nil && node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", indent, "", node.Data)
		indent -= 2
	}
}
