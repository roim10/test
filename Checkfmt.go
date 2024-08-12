package main

import (
	"fmt"
)

func main() {

	var MyList MyListNumbers
	fmt.Println(MyList.count())
}
// создаем структуру с именем MyListNumbers
type MyListNumbers struct {
// создаем две переменных типа string с именами i и j 
	i string
	j string
}
// создаем метод count который возращает тип строк, а принимает переменную с типом структуры MyListNumbers
func (MyList MyListNumbers) count() string {
	
	MyList.i = "1"
	MyList.j = "2"
	return fmt.Sprintf("enter numbers %s %s", MyList.i, MyList.j)
}
