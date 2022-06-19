package main

// 测试defer，panic，recover特性
// panic后会立即执行所有deferred func，LIFO顺序
// recover()只能出现在deferred func中
// 相当于deferred func可以给panic兜底
import "fmt"

func main() {
	testDeferParamEvaluate()
	result := testDeferModifyResult(5)
	fmt.Println("result = ", result)

	f()
	fmt.Println("return normally from f")
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

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from f ", r)
		}
	}()
	fmt.Println("call g 0")
	g(0)
	fmt.Println("return normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("g panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Printf("defer in g %d\n", i)
	fmt.Println("normal exec in g ", i)
	g(i + 1)
}
