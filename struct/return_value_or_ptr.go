// 本例对应《Go语言圣经》4.4节 提到的函数返回值为结构体时更新操作无法通过编译的问题进行测试
// 按照书中所说，返回指针时可以通过编译
package main

type Employee struct {
	Name, Address string
}

func NewEmployeeValue() Employee {
	return Employee{Name: "value return ?", Address: "v"}
}

func NewEmployeePtr() *Employee {
	return new(Employee)
}

func main() {
	// 这种写法会标红，无法通过编译
	//NewEmployeeValue().Name = "changed"
	// 这种就可以，因为返回了指针，则表示其有对应的内存地址了
	NewEmployeePtr().Name = "new val"
}
