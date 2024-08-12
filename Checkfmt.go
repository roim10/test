package main

import (
	"fmt"
)

func main() {

	var MyList MyListNumbers
	fmt.Println(MyList.count())
}

type MyListNumbers struct {
	i string
	j string
}

func (MyList MyListNumbers) count() string {
	MyList.i = "1"
	MyList.j = "2"
	return fmt.Sprintf("enter numbers %s %s", MyList.i, MyList.j)
}
