package main

import (
	"fmt"
)

func main() {
	var array [8]string
	fmt.Println(array)
	otherArray()
}
func otherArray() {
	array := [8]string{}
	i := len(array)
	fmt.Println(i)
}
