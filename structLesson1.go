package main

import (
	"fmt"
)

func main() {
	MyHateSubject := struct {
		name     string
		MyGrade  int
		MyReview string
	}{
		name:     "Hate",
		MyGrade:  5,
		MyReview: "bad",
	}

	fmt.Println(MyHateSubject)
}
