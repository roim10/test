package main

import "fmt"

func main() {

	fmt.Println(doctor{
		human:      human{},
		speciality: "doctor",
	})
}

type human struct {
	name string
	age  int
}
type doctor struct {
	human
	speciality string
}
