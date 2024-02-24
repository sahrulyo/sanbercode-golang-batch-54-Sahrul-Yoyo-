//latihan

package main

import "fmt"

//func printHello() {
//fmt.Println("Hello Bro!")

/*func bregedewDew() {
	fmt.Println("what's Up Bregedew")
}

func main() {
	bregedewDew()
}
*/
//function parameter
/*func printAngka(angka1 int, angka2 int) {
	fmt.Println("angka pertama", angka1)
	fmt.Println("angka kedua", angka2)
}

func main() {
	printAngka(1, 2)
}
*/

//function return value
/*func introduction(nama string) string {
	return "Hello my name is " + nama
}
func main() {
	lamgsung panggil didalam print
	fmt.Println(introduction("Bregedew"))

	//menggunakan variabel
	result := introduction("doe")
	fmt.Println(result)

	//function  as value
	secondResult := introduction
	fmt.Println(secondResult("Jack"))
} */

// Function Return Multiple Value
/*func introduction(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Name Is " + lastName
	return introFirstName, introLastName
}
func main() {
	firstName, lastName := introduction("John", "Doe")
	fmt.Println(firstName, lastName)

	firstName2, _ := introduction("Jane", "Smith")
	fmt.Println(firstName2)
} */

// contoh 1
/*func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

// contoh 2
func tampilkanKataDanAngka() (firstWord, secondWord string, number int) {
	firstWord = "Halo"
	secondWord = "Dunia"
	number = 10
	return
}

func main() {
	hasil := tambahAngka(4, 5)
	fmt.Println(hasil)

	kataPertama, kataKedua, angka := tampilkanKataDanAngka()
	fmt.Println(kataPertama, kataKedua, angka)
}
*/
//VARIADIC FUNCTION
/*func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Println(total)
} */
/*func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	var numbers = []int{2, 6, 7, 8, 9, 10}
	var total = sum(numbers...)
	fmt.Println(total)
}
*/
//Function Dengan Parameter Biasa & Variadic

/*
	func yourHobbies(name string, hobbies ...string) {
		fmt.Println("Hello, my name is", name)
		fmt.Println("My hobbies are: ")
		for _, hobby := range hobbies {
			fmt.Println(hobby)
		}
	}

	func main() {
		yourHobbies("John", "Gaming", "Hiking", "Reading")
		fmt.Println()

		var hobbies = []string{"Sleeping", "Eating"}
		yourHobbies("Doe", hobbies...)
	}
*/
/*func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Println("data :", numbers)
	fmt.Println("min :", min)
	fmt.Println("max :", max)
}
*/

//kode tidak berjalan
/*func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()
	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
} */

/*func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()
	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}
*/
// function yang menggunakan function sebagai parameternya
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// function yang digunakan sebagai parameter
func spamFilter(name string) string {
	if name == "Kasar" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("John", spamFilter)
	sayHelloWithFilter("Kasar", spamFilter)
}
