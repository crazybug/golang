package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Circle) area() float64 {
	return r.radius * r.radius * math.Pi
}

type MySlice []string

func (s MySlice) toString() string {
	var x string = ""
	for _, v := range s {
		x += v
	}
	return x
}

func main() {
	r1 := Rectangle{12, 5}
	c1 := Circle{10}

	fmt.Println("Area of r1 is :", r1.area())
	fmt.Println("Area of c1 is :", c1.area())

	myslice := MySlice{"a", "b", "c", "d"}
	fmt.Println(myslice)
	fmt.Println(myslice.toString())

}
