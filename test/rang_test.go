package test

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(i, sum)
	}

	kvs := map[string]string{"a": "1", "b": "2", "c": "3"}
	for _, v := range kvs {
		fmt.Printf("%s\n", v)
	}

	for _, v := range "abcdef" {
		fmt.Println(v)
	}

}
