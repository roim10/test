package main

import "fmt"

func main() {
	var array [5]int
	var slice []int

	fmt.Println(array)
	fmt.Println(slice)

	array[4] = 10
	fmt.Println(array)
	slice = append(slice, 5)
	fmt.Println(slice)

	array1 := []int{}
	fmt.Println(array1)
	fmt.Println(array1 == nil)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
