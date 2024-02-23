package main

import (
	"fmt"
)

func main() {
	//jawaban soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print(i, " - Berkualitas\n")
		} else {
			fmt.Print(i, " - Santai\n")
		}
		if i%3 == 0 && i%2 != 0 {
			fmt.Print("I Love Coding\n")
		}
	}
	//jawaban soal 2
	tinggi := 7

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	//jawaban soal 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	hasil := ""

	for i, kata := range kalimat {
		if i == len(kalimat)-1 {
			hasil += kata
		} else {
			hasil += kata + " "
		}
	}

	fmt.Println("[", hasil, "]")

	//jawaban soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, veggie := range sayuran {
		fmt.Printf("%d. %s\n", i+1, veggie)
	}
	//jawaban soal 5
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := 1
	for key, value := range satuan {
		fmt.Printf("%s = %d , ", key, value)
		volume *= value
	}
	fmt.Printf("volume balok = %d\n", volume)

}
