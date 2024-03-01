package main

import (
	"fmt"
	"formative-9/libraries"
)

func Soal1() {
	libraries.Soal1()
	segitiga := segitigaSamaSisi{alas: 5, tinggi: 4}
	persegi := persegiPanjang{panjang: 6, lebar: 3}

	var bangunDatar hitungBangunDatar
	bangunDatar = segitiga
	fmt.Println("Luas Segitiga:", bangunDatar.luas())
	fmt.Println("Keliling Segitiga:", bangunDatar.keliling())

	bangunDatar = persegi
	fmt.Println("Luas Persegi:", bangunDatar.luas())
	fmt.Println("Keliling Persegi:", bangunDatar.keliling())

	tabung := tabung{jariJari: 3, tinggi: 7}
	balok := balok{panjang: 4, lebar: 5, tinggi: 6}

	var bangunRuang hitungBangunRuang
	bangunRuang = tabung
	fmt.Println("Volume Tabung:", bangunRuang.volume())
	fmt.Println("Luas Permukaan Tabung:", bangunRuang.luasPermukaan())

	bangunRuang = balok
	fmt.Println("Volume Balok:", bangunRuang.volume())
	fmt.Println("Luas Permukaan Balok:", bangunRuang.luasPermukaan())

}

func main() {
	libraries.SayHello()

}
