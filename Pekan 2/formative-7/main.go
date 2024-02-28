package main

import "fmt"

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilmPtr *[]movie) {
	film := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilmPtr = append(*dataFilmPtr, film)
}

/*
	func main() {
//jawaban soal 1-------------------------------------> 1

		buahs := []buah{
			{"Nanas", "Kuning", false, 9000},
			{"Jeruk", "Oranye", true, 8000},
			{"Semangka", "Hijau & Merah", true, 10000},
			{"Pisang", "Kuning", false, 5000}, // Membuat slice untuk menyimpan objek buah
		}

		for _, buah := range buahs {
			fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %t, Harga: %d\n", buah.nama, buah.warna, buah.adaBijinya, buah.harga)
		} // Iterasi melalui setiap objek buah dan mencetak informasinya

//jawaban soal 2-------------------------------------> 2
		seg := segitiga{alas: 4, tinggi: 5}
		per := persegi{sisi: 4}
		perpan := persegiPanjang{panjang: 6, lebar: 4}

		fmt.Println("Luas Segitiga:", (seg.alas*seg.tinggi)/2)
		fmt.Println("Luas Persegi:", per.sisi*per.sisi)
		fmt.Println("Luas Persegi Panjang:", perpan.panjang*perpan.lebar)
	}
*/

//jawaban soal 3-------------------------------------> 3
/*func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
}
func main() {
	// Membuat objek phone
	p := phone{
		name:  "Model X",
		brand: "Brand A",
		year:  2022,
		colors: []string{
			"Black",
			"White",
		},
	}

	p.tambahWarna("Blue")

	fmt.Println("Name:", p.name)
	fmt.Println("Brand:", p.brand)
	fmt.Println("Year:", p.year)
	fmt.Println("Colors:", p.colors)
}*/

//jawaban soal 4---------------------------------------> 4

func main() {
	// Menambahkan data film menggunakan fungsi tambahDataFilm
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// Menampilkan data film
	fmt.Println("Data Film:")
	for _, film := range dataFilm {
		fmt.Printf("Title: %s, Duration: %d minutes, Genre: %s, Year: %d\n", film.title, film.duration, film.genre, film.year)
	}
	//OUTPUT
	/*Data Film:
	Title: LOTR, Duration: 120 minutes, Genre: action, Year: 1999
	Title: avenger, Duration: 120 minutes, Genre: action, Year: 2019
	Title: spiderman, Duration: 120 minutes, Genre: action, Year: 2004
	Title: juon, Duration: 120 minutes, Genre: horror, Year: 2004 */
}
