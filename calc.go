package main

import "fmt"

func main() {
	var num1, num2, result float32
	var act string
	fmt.Println("Enter first number")
	fmt.Scan(&num1)
	fmt.Println("Enter second number")
	fmt.Scan(&num2)
	fmt.Println("Enter action")
	fmt.Scan(&act)
	switch act {
	case "+":
		result = sum(num1, num2)
		fmt.Println(result)
	case "-":
		result = sub(num1, num2)
		fmt.Println(result)
	case "*":
		result = mul(num1, num2)
		fmt.Println(result)
	case "/":
		result = div(num1, num2)
		fmt.Println(result)
	}
}
func sum(num1, num2 float32) float32 {
	return num1 + num2
}
func sub(num1, num2 float32) float32 {
	return num1 - num2
}
func mul(num1, num2 float32) float32 {
	return num1 * num2
}
func div(num1, num2 float32) float32 {
	return num1 / num2
}
