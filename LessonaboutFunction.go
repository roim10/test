package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	i := 0
	i = implements(i)
	fmt.Println(i)
}

func implements(i int) int {
	i++
	return i
}
