package main

import (
	"fmt"
	"strconv"
)

//DEKLARASI DENGAN VAR

/*func main() {
	var text = "Hello, world!"
	fmt.Println(text)

	text = "Hello World diubah nich"
	fmt.Println(text)
} */

//DEKLARASI DENGAN VAR DAN TIPE DATA

/*func main() {
	//contoh 1
	var text1 string
	text1 = "ini adalah contoh teks 1"
	fmt.Println(text1)

	//contoh 1
	var text2 string = "ini adalah contoh teks 2"
	text2 = "ini text 1 diubah"
	fmt.Println(text2)

}*/

// CONSTANT
//jenis variabel yang nilainya tidak bisa diubah.

/*func main() {
	const judul = "Ini Judul"
	fmt.Println(judul)
	// jika kode di bawah ini di uncomment maka akan error
	//judul = "judul diubah"

}*/

func main() {
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // true

}
