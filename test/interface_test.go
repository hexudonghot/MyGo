package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	ju := juxing{10, 10}
	r := inter_circle{5.5}
	measure(ju)
	measure(r)
}

type geometry interface {
	inter_area() float64
	inter_perim() float64
}

type juxing struct {
	width, height float64
}

type inter_circle struct {
	radius float64
}

func (r juxing) inter_area() float64 {
	return r.height * r.width
}
func (r inter_circle) inter_area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r juxing) inter_perim() float64 {
	return 2 * (r.width + r.height)
}

func (r inter_circle) inter_perim() float64 {
	return 2 * math.Pi * r.radius * r.radius
}

func measure(g geometry) {
	fmt.Println("________________________")
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println("________________________")
	fmt.Println(g.inter_perim())
	fmt.Println(g.inter_perim())
}
