package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//jawaban soal 1
	var Satu = "Bootcamp"
	var Dua = "Digital"
	var Tiga = "Skill"
	var Empat = "Sanbercode"
	var Lima = "Golang"
	kalimat := Satu + " " + Dua + " " + Tiga + " " + Empat + " " + Lima
	fmt.Println(kalimat)

	//jawaban soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)

	//jawaban soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	var kalimatSaya = kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println(kalimatSaya)

	//jawaban soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	/* short variable declaration */
	var num, err = strconv.Atoi(angkaPertama + angkaKedua + angkaKetiga + angkaKeempat)

	if err == nil {
		fmt.Println(num)
		/*total := strconv.Itoa(num)
		fmt.Println(total)*/
	}

	//jawaban soal 5
	lirikLagu := "halo halo bandung"
	angka := 2021

	lirikLagu = strings.Replace(lirikLagu, "halo", "Hi", 2)
	lirikLagu = strings.Replace(lirikLagu, "", "bandung", 0)
	lirikLagu += " - " + strconv.Itoa(angka)
	fmt.Println(lirikLagu)

}
