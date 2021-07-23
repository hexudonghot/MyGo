package test

import "fmt"

func main() {
	/**
	  array------start
	*/
	var a [5]int
	//fmt.Println(a)
	a[4] = 6
	//fmt.Println(a)

	//fmt.Println(len(a))
	b := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(b)

	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}
	/**
	  array------end
	*/
	s := make([]string, 3)
	fmt.Println(s)

}
