package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menangani permintaan untuk file HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Mengirimkan file HTML sebagai respons
		http.ServeFile(w, r, "index.html")
	})

	// Menangani permintaan untuk hasil pencarian
	http.HandleFunc("/hasil-pencarian", func(w http.ResponseWriter, r *http.Request) {
		// Memproses data yang diterima dari formulir
		r.ParseForm()
		punyaBanyakPacar := r.Form.Get("punyaBanyakPacar")
		pasanganSetia := r.Form.Get("pasanganSetia")

		// Menampilkan hasil berdasarkan data yang diterima
		if punyaBanyakPacar == "ya" {
			fmt.Fprintf(w, "Selamat! Anda terdaftar sebagai cowok komersial.")
		} else {
			fmt.Fprintf(w, "Anda telah mendaftar sebagai pencari jodoh.")
			if pasanganSetia != "" {
				fmt.Fprintf(w, " Namun, cukup dirimu satu.")
			} else {
				fmt.Fprintf(w, " Dan sedang ingin sendiri.")
			}
		}
	})

	// Menjalankan server pada port 8080
	http.ListenAndServe(":8080", nil)
}
