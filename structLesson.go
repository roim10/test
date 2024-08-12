package main

import "fmt"

type car struct {
	make       string
	model      string
	height     int
	width      int
	backwheel  wheel
	frontwheel wheel
}

type wheel struct {
	material string
	radius   geometry
	height   geometry
	width    geometry
}
type geometry struct {
	height int
	radius int
	width  int
}

func main() {
	toyta := car{}
	fmt.Println("name:")
	toyta.make = "toyta"
	toyta.backwheel.radius.height = 10
	fmt.Println(toyta)
}
