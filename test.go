package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum1(numbers []int, criteria func(int) bool) int {

	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}
func main() {

	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEvens := sum1(slice, isEven) // сумма четных чисел
	fmt.Println(sumOfEvens)           // -2

	sumOfPositives := sum1(slice, isPositive) // сумма положительных чисел
	fmt.Println(sumOfPositives)               // 37
}
