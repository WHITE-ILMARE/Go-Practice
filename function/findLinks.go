// 找到输入的地址指向的HTML文件中的所有超链接，主要练习递归函数的用法
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findAndVisit(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"error occur: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Printf("link: %s\n", link)
		}
	}
}

// 将一个节点及其子节点中的a标签中的超链接提取出来存入links中
func BFS_visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		links = BFS_visit(links, nextNode)
	}
	return links
}

// 访问参数中的链接，并返回链接中的超链接地址
func findAndVisit(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if (resp.StatusCode != http.StatusOK) {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s error: %s\n", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return BFS_visit(nil, doc), nil
}
