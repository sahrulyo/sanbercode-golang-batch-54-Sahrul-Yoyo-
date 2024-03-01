package libraries

import (
	"fmt"
	"math"
)

func SayHello() {
	fmt.Println("hello")
}

//soal no. 1 ------------------------------------------->

func Soal1() {
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
	
}


