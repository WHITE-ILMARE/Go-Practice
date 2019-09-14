// 初试web相关
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if (!strings.HasPrefix(url, "http://")) {
			url = "http://" + url;
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchL %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP Status Code=%d, Status=%s\n", resp.StatusCode, resp.Status)
		//_, errC := io.Copy(os.Stdout, resp.Body); // 直接将内容复制到标准输出中，避免申请缓冲区(即下文中的b)
		//b, err := ioutil.ReadAll(resp.Body) // 或将内容存入暂存区b
		resp.Body.Close()
		//if errC != nil {
		//	fmt.Fprintf(os.Stderr, "fetch error: %s: %v\n", url, err)
		//	os.Exit(1)
		//	//fmt.Printf("%s", b)
		//}
	}
}