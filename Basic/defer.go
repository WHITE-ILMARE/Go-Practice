package main

// 测试defer特性
import "fmt"

func main() {
	testDeferParamEvaluate()
	result := testDeferModifyResult(5)
	fmt.Println("result = ", result)
}

// 测试defer参数值何时计算
func testDeferParamEvaluate() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func testDeferModifyResult(x int) (result int) {
	defer func() { result += x }()
	result = x + x
	return
}
