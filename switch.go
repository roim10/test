package main

import "fmt"

func main() {
	fmt.Println("enter number from 0 to 9")
	var num int
	fmt.Scan(&num)
	if num > 9 || num < 0 {
		fmt.Println("wrong number")
	}
	switch num {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	}
}
