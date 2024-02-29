package main

import (
	"fmt"
	"math"
)

//jawaban soal 1 ----------------------------------------->1

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2*p.panjang + 2*p.lebar
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2*math.Pi*t.jariJari*t.tinggi + 2*math.Pi*t.jariJari*t.jariJari
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2*float64(b.panjang*b.lebar) + 2*float64(b.panjang*b.tinggi) + 2*float64(b.lebar*b.tinggi)
}

func soal1() {
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

// jawaban soal 2 ----------------------------------------->2
type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

type displayInfo interface {
	display() string
}

func (p phone) display() string {
	info := fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.name, p.brand, p.year, p.colors)
	return info
}

func soal2() {
	myPhone := phone{
		name:   "iPhone 13",
		brand:  "Apple",
		year:   2021,
		colors: []string{"Black", "White", "Blue"},
	}

	fmt.Println(myPhone.display())

	//OUTPUT
	/*Name: iPhone 13
	Brand: Apple
	Year: 2021
	Colors: [Black White Blue]*/
}

// jawaban soal 3 ----------------------------------------->3

func soal3() {
	fmt.Println(displayLuasPersegi(4, true))
	fmt.Println(displayLuasPersegi(8, false))
	fmt.Println(displayLuasPersegi(0, true))
	fmt.Println(displayLuasPersegi(0, false))
}

func displayLuasPersegi(sisi int, verbose bool) interface{} {
	if sisi == 0 {
		if verbose {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}

	luas := sisi * sisi
	if verbose {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}

// jawaban soal 4 ----------------------------------------->4
func main() {
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	/* Konversi var menjadi tipe slice []int
	pake type assertion */
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	output := fmt.Sprintf("%v %d + %d + %d + %d = %d", prefix, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
	fmt.Println(output)
	//OUTPUT hasil penjumlahan dari  6 + 8 + 12 + 14 = 40
}
