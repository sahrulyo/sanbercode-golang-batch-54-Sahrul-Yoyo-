package main

import (
	"fmt"
	"strconv"
)

func main() {
	// contoh 1
	var text1 string
	text1 = "ini teks 1"
	fmt.Println(text1)

	// contoh 1
	var text2 string = "ini teks 2"
	text2 = "ini teks 2 diubah"
	fmt.Println(text2)

	//KONVERSI STRING KE INT
	angkaString := "3789"
	angka, err := strconv.Atoi(angkaString)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Angka:", angka)

	//KONVERSI STRING KE INT DAN MEMBUAT PERTAMBAHAN MATEMATIKA
	nilai1String := "5"
	nilai2String := "3"

	nilai1, err := strconv.Atoi(nilai1String)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	nilai2, err := strconv.Atoi(nilai2String)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	hasil := nilai1 * nilai2
	fmt.Println("hasil perkalian:", hasil)

	//CONSTANT - DISABLE EDIT
	const judul = "Go Go Language"
	fmt.Println(judul)
	//judul = "judul diubah menjadi Hu Hu Language"

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//CONTOH fmt.PrintLn X fmt.PrintF
	nama := "Alice"
	umur := 30

	// fmt.Printf
	fmt.Printf("Halo, nama saya %s dan saya berumur %d tahun.\n", nama, umur)

	// fmt.Println
	fmt.Println("Halo, nama saya", nama, "dan saya berumur", umur, "tahun.")

	pertambahan := 8 + 2
	fmt.Println("hasil pertambahan ", pertambahan)

	//augmented assignments
	var biji = 10
	biji += 10
	fmt.Println(biji)

	//OPERATOR KETERANGAN
	angka2 := 8
	fmt.Println(angka2) // 8
	angka2++
	fmt.Println(angka2) // 9

	angka3 := 5
	fmt.Println(angka3) // 5
	angka3--
	fmt.Println(angka3) // 4

	//LOGIC   Operator Logika
	var angka5 = 8

	fmt.Println(angka5 > 5) // true

	fmt.Println(angka5 < 5) // false

	fmt.Println(angka5 >= 5) // true

	fmt.Println(angka5 <= 5) // false

	fmt.Println(angka5 == 5) // false

	fmt.Println(angka5 != 5) // true

	var a = true
	var b = false
	var c = true
	var d = false

	fmt.Println(a && c) // true

	fmt.Println(a && b) // false

	fmt.Println(a || b) // true

	fmt.Println(b || d) // false

	fmt.Println(!b && !d) // true

	fmt.Println(!a || b) // false

	//CONDITIONAL IF, ELSE IF, ELSE

	if true {
		fmt.Println("jalankan kode")
	}
	if false {
		fmt.Println("tidak jalankan kode")
	}
	var mood = "happy"
	if mood == "happy" {
		fmt.Println("i am happy")
	}

	var banyakCewek = "Laku"
	if banyakCewek == "Laku" {
		fmt.Println("aku ganteng")
	}

	var udahDinilai = "tidak sibuk"
	if udahDinilai == "tidak sibuk" {
		fmt.Println("sedang libur")
	}

	/*var miniMarketStatus = "close"
	var minuteRemainingToOpen = 5
	if miniMarketStatus == "open" {
		fmt.Println("saya belanja sekarang")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("minimarket buka sebentar lagi saya tungguin")
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi ")
	}*/

	var punyaBanyakPacar = "keren"
	var pasanganSetia = 2

	if punyaBanyakPacar == "keren" {
		fmt.Println("cowok komersial")
	} else if pasanganSetia <= 2 {
		fmt.Println("cukup dirimu satu")
	} else {
		fmt.Println("sedang ingin sendiri")
	}

	//SLICE
	// Membuat slice kosong dengan panjang awal 0
	var angkax []int
	fmt.Println("Slice kosong:", angkax)

	// Menambahkan elemen ke slice menggunakan fungsi append
	angkax = append(angkax, 1)
	angkax = append(angkax, 2, 3, 4)
	fmt.Println("Setelah ditambahkan elemen:", angkax)

	// Mengakses elemen slice
	fmt.Println("Elemen pada indeks ke-0:", angkax[0])
	fmt.Println("Elemen pada indeks ke-2:", angkax[2])

	// Mengubah elemen slice
	angkax[0] = 10
	fmt.Println("Setelah diubah elemen pada indeks ke-0:", angkax)

	// Menghapus elemen dari slice
	angkax = append(angkax[:2], angkax[3:]...)
	fmt.Println("Setelah dihapus elemen pada indeks ke-2:", angkax)

	// Menggunakan slice dalam perulangan
	for i, v := range angkax {
		fmt.Printf("Indeks ke-%d: %d\n", i, v)
	}

	//SWITCH

	// Variabel umur
	umur = 25

	// Switch berdasarkan nilai umur
	switch umur {
	case 0: //input___________________>
		fmt.Println("Bayi")
	case 1, 2, 3, 4:
		fmt.Println("Balita")
	case 5, 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Anak-anak")
	case 13, 14, 15, 16, 17:
		fmt.Println("Remaja")
	default:
		fmt.Println("Dewasa")
	}

	// Switch tanpa kondisi (mirip if-else if-else)
	nilai := 80 //input________________>
	switch {
	case nilai >= 90:
		fmt.Println("Nilai A")
	case nilai >= 80:
		fmt.Println("Nilai B")
	case nilai >= 70:
		fmt.Println("Nilai C")
	case nilai >= 60:
		fmt.Println("Nilai D")
	default:
		fmt.Println("Nilai E")
	}

	//LOOPING
	// looping for
	fmt.Println("contoh pengulangan menggunakan for:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//looping for dan slice
	angkaz := []int{1, 2, 3, 4, 5}
	for index, nilai := range angkaz {
		fmt.Println("Indeks ke-%d: %d/n", index, nilai)
	}
	/*fmt.Println("\nContoh perulangan tak hingga (untuk dipaksa berhenti tekan CTRL+C):")
	for {
		fmt.Println("Loop tak hingga")
	} */

}
