package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
}
type circle struct {
	radius float64
}

func (s square) area() int {
	return s.length * s.length
}
func (s square) perimeter() int {
	return s.length * 4
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

type shape interface {
	area() float64
	perimeter() float64
}

func info(ccl shape) {

	fmt.Println(ccl.area())
	fmt.Println(ccl.perimeter())
}
func main() {
	fmt.Println("enter square length")
	var length int
	fmt.Scan(&length)
	s := square{
		length: length,
	}
	fmt.Println("enter circle radius")
	var radius float64
	fmt.Scan(&radius)
	fmt.Println(s.perimeter())
	fmt.Println(s.area())
	info(circle{radius: radius})
}
