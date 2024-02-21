package main

import (
	"fmt"
	"strconv"
)

func main() {

	//jawaban soal 1 ------------------------------------------>
	var (
		panjangPersegiPanjang = "8"
		lebarPersegiPanjang   = "5"
		alasSegitiga          = "6"
		tinggiSegitiga        = "7"
	)
	total := 0
	digits := []string{panjangPersegiPanjang, lebarPersegiPanjang, alasSegitiga, tinggiSegitiga}
	for _, digits := range digits {
		num, err := strconv.Atoi(digits)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		total += num
	}
	fmt.Println("Total:", total)

	// persegi panjang
	panjang, err := strconv.Atoi(panjangPersegiPanjang)
	if err != nil {
		fmt.Println("Error konversi panjang:", err)
		return
	}

	lebar, err := strconv.Atoi(lebarPersegiPanjang)
	if err != nil {
		fmt.Println("Error konversi lebar:", err)
		return
	}

	// Hitung luas dan keliling persegi panjang
	luasPersegiPanjang := panjang * lebar
	kelilingPersegiPanjang := 2 * (panjang + lebar)

	//segitiga
	alas, err := strconv.Atoi(alasSegitiga)
	if err != nil {
		fmt.Println("Error konversi alas:", err)
		return
	}

	tinggi, err := strconv.Atoi(tinggiSegitiga)
	if err != nil {
		fmt.Println("Error konversi tinggi:", err)
		return
	}

	// Hitung luas segitiga
	luasSegitiga := (alas * tinggi) / 2

	// Tampilkan hasil perhitungan keseluruhan
	fmt.Println("Luas persegi panjang:", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas segitiga:", luasSegitiga)

	//soal no 2 ------------------------------------------>
	var nilaiJohn = 80
	var nilaiDoe = 50

	// Penentuan indeks nilai untuk nilaiJohn
	fmt.Print("Indeks nilai John: ")
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// Penentuan indeks nilai untuk nilaiDoe
	fmt.Print("Indeks nilai Doe: ")
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	//jawaban soal 3 ------------------------------------------>
	var tanggal = 3
	var bulan = 3
	var tahun = 1989

	fmt.Print("Tanggal lahir saya: ")
	switch bulan {
	case 1:
		fmt.Print(tanggal, " Januari ", tahun)
	case 2:
		fmt.Print(tanggal, " Februari ", tahun)
	case 3:
		fmt.Print(tanggal, " Maret ", tahun)
	case 4:
		fmt.Print(tanggal, " April ", tahun)
	case 5:
		fmt.Print(tanggal, " Mei ", tahun)
	case 6:
		fmt.Print(tanggal, " Juni ", tahun)
	case 7:
		fmt.Print(tanggal, " Juli ", tahun)
	case 8:
		fmt.Print(tanggal, " Agustus ", tahun)
	case 9:
		fmt.Print(tanggal, " September ", tahun)
	case 10:
		fmt.Print(tanggal, " Oktober ", tahun)
	case 11:
		fmt.Print(tanggal, " November ", tahun)
	case 12:
		fmt.Print(tanggal, " Desember ", tahun)
	default:
		fmt.Println("Bulan tidak valid")
		return
	}

	//jawaban soal 4 ------------------------------------------>

	var tahunKelahiran = 1989

	fmt.Print("Anda termasuk dalam generasi: ")
	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		fmt.Println("Generasi X")
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		fmt.Println("Generasi Y (Millennials)")
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("Generasi tidak diketahui")
	}
}
