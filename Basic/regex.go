package main

import (
	"fmt"
	"regexp"
)

func main() {
	testIp()
}

func testIp() {
	ip := "127.0.0.1"
	// 双反斜杠才能转义成功
	matched, _ := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, ip)
	fmt.Println(matched)
}
