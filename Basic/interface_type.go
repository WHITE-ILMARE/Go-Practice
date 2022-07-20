package main

// 测试interface{}动态类型的demo
import (
	"fmt"
)

func main() {
	var stu interface{} = Student{}
	var stup interface{} = &Student{}
	var stup2 interface{} = (*Student)(nil)
	var init interface{}
	var inti interface{} = 1
	judgeType(stu)
	judgeType(stup)
	judgeType(stup2)
	judgeType(init)
	judgeType(inti)
}

func judgeType(v interface{}) {
	fmt.Printf("%p, %v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p, %v\n", &v, v)
		fmt.Printf("nil type[%T], %v\n", v, v)
	case Student:
		fmt.Printf("%p, %v\n", &v, v)
		fmt.Printf("Student type[%T], %v\n", v, v)
	case *Student:
		fmt.Printf("%p, %v\n", &v, v)
		fmt.Printf("*Student type[%T], %v\n", v, v)
	default:
		fmt.Printf("%p, %v\n", &v, v)
		fmt.Printf("unknown type[%T], %v\n", v, v)
	}

}

type Student struct {
	Name string
	Age  int
}
