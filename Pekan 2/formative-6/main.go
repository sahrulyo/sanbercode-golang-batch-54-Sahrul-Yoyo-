package main

import "fmt"

var dataFilm = []map[string]string{} //VARIABEL UNTUK SOAL 4

// JAWABAN SOAL 1 ------------------------>
func updateLingkaran(luasLingkaran *float64, kelilingLingkaran *float64, jariJari float64) {
	const pi float64 = 3.14
	*luasLingkaran = pi * jariJari * jariJari
	*kelilingLingkaran = 2 * pi * jariJari
}
func main() {
	var luasLingkaran float64
	var kelilingLingkaran float64
	var jariJari float64 = 5

	updateLingkaran(&luasLingkaran, &kelilingLingkaran, jariJari)
	fmt.Println("Luas Lingkaran:", luasLingkaran)         //Luas Lingkaran: 78.5
	fmt.Println("Keliling Lingkaran:", kelilingLingkaran) //Keliling Lingkaran: 31.400000000000002

	// JAWABAN SOAL 2  ------------------------>
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}
func introduce(sentence *string, name string, gender string, profession string, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, profession, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, profession, age)
	}

	// JAWABAN SOAL 3  ------------------------>
	var buah = []string{}

	tambahData(&buah, "Jeruk")
	tambahData(&buah, "Semangka")
	tambahData(&buah, "Mangga")
	tambahData(&buah, "Strawberry")
	tambahData(&buah, "Durian")
	tambahData(&buah, "Manggis")
	tambahData(&buah, "Alpukat")

	tampilkanBuah(&buah)
}
func tambahData(buah *[]string, item string) {
	*buah = append(*buah, item) // menambahkan data ke variabel buah
}
func tampilkanBuah(buah *[]string) {
	for i, item := range *buah {
		fmt.Printf("%d. %s\n", i+1, item)
	} //menampilkan isi variabel buah dengan nomor urut di depannya

	//JAWABAN SOAL 4 ------------------------>
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	tampilkanDataFilm(dataFilm)
}

func tambahDataFilm(judul string, durasi string, genre string, tahun string, dataFilm *[]map[string]string) {
	film := make(map[string]string)
	film["judul"] = judul
	film["durasi"] = durasi
	film["genre"] = genre
	film["tahun"] = tahun
	*dataFilm = append(*dataFilm, film)
}
func tampilkanDataFilm(dataFilm []map[string]string) {
	for i, film := range dataFilm {
		fmt.Printf("%d. Judul: %s\n", i+1, film["judul"])
		fmt.Printf("   Durasi: %s\n", film["durasi"])
		fmt.Printf("   Genre: %s\n", film["genre"])
		fmt.Printf("   Tahun: %s\n", film["tahun"])
	}
	//OUTPUT NO. 4
	/*1. Judul: LOTR
		Durasi: 2 jam
		Genre: action
		Tahun: 1999
	 2. Judul: avenger
		Durasi: 2 jam
		Genre: action
		Tahun: 2019
	 3. Judul: spiderman
		Durasi: 2 jam
		Genre: action
		Tahun: 2004
	 4. Judul: juon
		Durasi: 2 jam
		Genre: horror
		Tahun: 2004
	 5. Judul: LOTR
		Durasi: 2 jam
		Genre: action
		Tahun: 1999
	 6. Judul: avenger
		Durasi: 2 jam
		Genre: action
		Tahun: 2019
	 7. Judul: spiderman
		Durasi: 2 jam
		Genre: action
		Tahun: 2004
	 8. Judul: juon
		Durasi: 2 jam
		Genre: horror
		Tahun: 2004 */
}
