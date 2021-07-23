package test

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {

	user1 := user{"hxd", "ddd", 0}
	user2 := &user1
	user2.name = "ydddd"
	fmt.Println(user2.name)
	fmt.Println(user1.name)

	r := rect{10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}

type user struct {
	name     string
	password string
	age      int
}

//为结构体定义专属的方法

type rect struct {
	width, height int
}

//结构体方法定义
func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return (r.height + r.width) * 2
}
