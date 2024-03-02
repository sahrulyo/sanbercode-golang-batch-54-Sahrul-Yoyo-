package main

import (
	"fmt"
	"formative-9/libraries"
)

func main() {
	libraries.SayHello()
	libraries.Soal1()
	libraries.Soal3()
	libraries.Soal4()

}
func Soal2() {
	myPhone := libraries.Phone{
		Name:   "iPhone 13",
		Brand:  "Apple",
		Year:   2021,
		Colors: []string{"Black", "White", "Blue"},
	}

	fmt.Println(myPhone.Display())
}
