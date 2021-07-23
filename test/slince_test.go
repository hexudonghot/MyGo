package test

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	fmt.Println(s)
	s = append(s, "c")
	s = append(s, "d", "e", "f")
	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	l := s[:2]
	fmt.Println(l)
}
