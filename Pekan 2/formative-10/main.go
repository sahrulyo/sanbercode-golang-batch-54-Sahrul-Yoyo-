package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// jawaban soal 1------------------------------------------------------------------------>
func Soal1() {
	kalimat := "Golang Backend Development"
	tahun := 2021

	defer printKalimatDanTahun(kalimat, tahun)

	fmt.Println("Program selesai")
}

func printKalimatDanTahun(kalimat string, tahun int) {
	fmt.Printf("Kalimat: %s\n", kalimat)
	fmt.Printf("Tahun: %d\n", tahun)
}

//jawaban soal 2------------------------------------------------------------------------>

func soal2() {
	// Contoh pemanggilan fungsi
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

func kelilingSegitigaSamaSisi(sisi int, tampilkanDetail bool) (string, error) {
	if sisi == 0 {
		if tampilkanDetail {
			return "", fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic:", r)
				}
			}()
			panic(fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi"))
		}
	}

	keliling := sisi * 3
	if tampilkanDetail {
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

// jawaban soal 3------------------------------------------------------------------------>
func soal3() {
	angka := 1
	defer cetakAngka(&angka)

	// call function tambahAngka beberapa kali
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Println("Total angka:", *total)
}

// jawaban soal 4------------------------------------------------------------------------>

//var phones = []string{}

func soal4() {

	tambahData(&phones)

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
}

func tambahData(data *[]string) {
	*data = append(*data, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
}

/*OUTPUT : muncul satu persatu selama per-detik
1. Asus
2. IPhone
3. Oppo
4. Realme
5. Samsung
6. Vivo
7. Xiaomi */

// jawaban soal 5------------------------------------------------------------------------>
var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

func soal5() {

	var wg sync.WaitGroup
	wg.Add(len(phones))

	// Mengurutkan data phones
	sort.Strings(phones)

	for i, phone := range phones {
		go func(index int, phone string) {
			defer wg.Done()
			fmt.Printf("%d. %s\n", index+1, phone)
			time.Sleep(time.Second)
		}(i, phone)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()
}

//jawaban soal 6---------------------------------------------------------------------->

var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

func main() {
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func getMovies(ch chan<- string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}
