package test

import (
	"fmt"
	"testing"
)

func TestRecusion(t *testing.T) {
	fmt.Println(sums(10))
}

func sums(num int) int {
	if num == 1 {
		return num
	}
	return sums(num-1) + num
}
