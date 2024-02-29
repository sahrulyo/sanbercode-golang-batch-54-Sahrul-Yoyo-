//interface
//tipe data abstract murni isinya sebagai kontrak

// contoh :
package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName()+"!")
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var eko Person
	eko.Name = "eko"
	SayHello(eko)
}
