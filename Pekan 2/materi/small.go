package main

import "fmt"

// struct
type Gorengan struct {
	nama      string
	bahan     int
	caraMasak []string
}

// function
func main() {
	var angka int = 8
	angka += 10
	fmt.Println(angka)

	var telur string = "rebus"
	telur = "mateng"
	fmt.Println("makan telur", telur)

	bubur := "Sarapan"
	ayam := "daging"
	fmt.Println("pagi-pgi", bubur)
	fmt.Println("makan", ayam)

	//struct
	gorengan1 := Gorengan{
		nama:      "Bala-bala",
		bahan:     15,
		caraMasak: []string{"iris sayuran", "campur terigu", "goreng"},
	}
	fmt.Println("nama gorengan", gorengan1.nama)
	fmt.Println("jumlah bahan", gorengan1.bahan)
	fmt.Println("cara memasak")
	//slice
	for i, langkah := range gorengan1.caraMasak {
		fmt.Println(i+1, langkah)
	}

	//interface

}
