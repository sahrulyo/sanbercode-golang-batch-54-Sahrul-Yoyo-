package main

import (
	"fmt"
)

// jawaban soal 1 --------------------------------------------> di comment jarena "main" double declaration ya..
/*func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

func main() {
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
	fmt.Println("Volume Balok:", volume)
} */

//jawaban soal 2

/*func introduce(nama, gender, pekerjaan, usia string) string {
	var title string
	if gender == "laki-laki" {
		title = "Pak"
	} else {
		title = "Bu"
	}
	// Membentuk kalimat hasil
	intro := fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, nama, pekerjaan, usia)

	return intro
}

func main() {
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)
}*/

// jawaban soal 3
/*func buahFavorit(nama string, buah ...string) string {
	buahStr := strings.Join(buah, ", ")
	result := fmt.Sprintf("Halo nama saya %s dan buah favorit saya adalah \"%s\"", nama, buahStr)

	return result
}

func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	buahFavoritJohn := buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)
} */

// jawaban soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm() func(string, string, string, string) {
	return func(judul, durasi, genre, tahun string) {
		film := map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}
}

func main() {
	tambahData := tambahDataFilm()

	tambahData("LOTR", "2 jam", "action", "1999")
	tambahData("avenger", "2 jam", "action", "2019")
	tambahData("spiderman", "2 jam", "action", "2004")
	tambahData("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
