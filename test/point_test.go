package test

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	//指针调用
	//定义一个数字
	i := 10
	fmt.Println("initial: ", i)

	byValue(i)
	fmt.Println("byval:", i)
	// 地址传递给函数byValue
	byPointer(&i)
	fmt.Println("byptr:", i)
	// 打印出 i 的地址
	fmt.Println("pointer:", i)

}

func byPointer(iptr *int) {
	//表示指针指向的数值
	*iptr = 0

}

func byValue(ival int) {
	ival = 0
}
