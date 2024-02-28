/*Ketika kita berbicara tentang bahasa pemrograman, ada sesuatu yang disebut "class". Ini adalah semacam wadah di mana Anda bisa menaruh variabel dan fungsi bersama-sama. Tetapi, dalam bahasa Go, istilah "class" tidak digunakan. Sebaliknya, mereka menggunakan sesuatu yang disebut "struct".

Struct ini sebenarnya mirip dengan class dalam bahasa pemrograman lain. Ini adalah tempat di mana Anda bisa menaruh variabel dan fungsi bersama-sama. Tetapi bedanya, dalam Go, Anda tidak menggunakan kata kunci "class", melainkan "struct".

Jadi, dalam bahasa Go, Anda membuat struct untuk mengelompokkan variabel dan fungsi bersama-sama. Ini memungkinkan Anda untuk membuat sesuatu yang berperilaku mirip dengan class di bahasa pemrograman lain. Misalnya, Anda bisa membuat struct yang menyimpan data tentang seseorang, seperti nama, umur, dan alamat, serta fungsi-fungsi yang beroperasi pada data itu. Jadi, secara sederhana, struct di Go adalah cara untuk mengatur data dan operasi-operasi yang berhubungan dengan data tersebut.*/

package main

import "fmt"

// Deklarasi struct bernama Person
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Jabatan string
	gaji    int
}

// embedded ---------------------------------------->
type Biodata struct {
	Id          string
	CountryCode int
}

type Role struct {
	role     string
	whatsApp int
	Biodata
}

// anonymous struct --------------------------------->
type jalema struct {
	ngaran string
	umur   int
}

// nested struct
type student struct {
	person struct {
		name string
		age  int
	}
	grade int
}

// method
type students struct {
	names string
	grade int
}

func main() {
	// Membuat variabel bertipe struct Person
	var person1 Person

	// Mengisi nilai bidang atau field dari struct
	person1.Name = "John"
	person1.Age = 30

	// Menampilkan nilai dari struct
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)

	var employee1 Employee
	employee1.Jabatan = "Manager Area"
	employee1.gaji = 50000
	fmt.Println("Jabatan:", employee1.Jabatan)
	fmt.Println("Gaji $", employee1.gaji)

	//struct literals
	//cara pertama
	var john = Person{}
	john.Name = "Wick"
	john.Age = 45

	fmt.Println("Nama John", john.Name)
	fmt.Println("Usia", john.Age)

	//cara kedua
	var doe = Person{"doe", 50000}
	fmt.Println(doe.Name)

	var tom = Person{Name: "Tom", Age: 45}
	fmt.Println(tom)

	//Embedded Struct ------------------------>

	var sahrul = Biodata{}
	sahrul.Id = "SAHRUL7355"
	sahrul.CountryCode = 62
	var yoyo = Role{}
	yoyo.role = "Developer"
	yoyo.whatsApp = 6281213163186

	fmt.Println("id", sahrul.Id)
	fmt.Println("Kode negara:", sahrul.CountryCode)
	fmt.Println(yoyo.role)
	fmt.Println(yoyo.whatsApp)

	//Anonymous Struct------------------------------->
	var arsam = struct {
		jalema
		Age int
	}{}
	arsam.jalema = jalema{"wick", 21}
	arsam.Age = 2

	fmt.Println("name Arsam :", arsam.jalema.ngaran)
	fmt.Println("age   :", arsam.jalema.umur)
	fmt.Println("grade :", arsam.Age)

	// anonymous struct tanpa pengisian property
	/*var rosid = struct {
		jalema
		Age int
	}{}

	// anonymous struct dengan pengisian property
	var herun = struct {
		jalema
		Age int
	}{
		jalema: jalema{"arsam", 21},
		Age:    2,
	} */
}

//method
/*Cara menerapkan method sedikit berbeda dibanding penggunaan fungsi. Ketika deklarasi, ditentukan juga siapa pemilik method tersebut. Contohnya bisa dilihat pada kode berikut: */

func (s students) sayHello() {
	fmt.Println("halo", s.names)
}

/*func main() {
    var john = student{"john wick", 21}
    john.sayHello()
}*/
