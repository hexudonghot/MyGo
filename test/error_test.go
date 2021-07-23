package test

import (
	"errors"
	"fmt"
	"testing"
)

//错误处理
func TestError(t *testing.T) {
	result, err := f1(10)
	fmt.Println(result, err)
	result, err = f1(-10)
	fmt.Println(result, err)
	// f2
	result, err = f2(10)
	fmt.Println(result, err)

	result, err = f2(-10)
	fmt.Println(result, err)

}

func f1(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("参数错误")
	}
	return 2 * num, nil
}

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d -> %s", e.arg, e.msg)
}

func f2(num int) (int, error) {
	if num < 0 {
		return -1, &argError{num, "参数不能为负值"}
	}
	return 2 * num, nil
}
