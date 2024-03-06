package main

import (
	"encoding/json"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var username = "user"
var password = "pass"

func main() {
	http.HandleFunc("/post_nilai", BasicAuth(PostNilaiMahasiswa, username, password))
	http.HandleFunc("/get_nilai", GetNilaiMahasiswa) // soal no 2
	http.ListenAndServe(":8080", nil)
}

// jawaban soal no 1 ------------------------------------------------------->
func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Cek Basic Auth
	username, password, ok := r.BasicAuth()
	if !ok || username != "user" || password != "pass" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode JSON
	var newNilai NilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&newNilai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi Nilai
	if newNilai.Nilai > 100 {
		http.Error(w, "Nilai tidak valid (maksimal 100)", http.StatusBadRequest)
		return
	}

	// Tentukan Indeks Nilai
	if newNilai.Nilai >= 80 {
		newNilai.IndeksNilai = "A"
	} else if newNilai.Nilai >= 70 {
		newNilai.IndeksNilai = "B"
	} else if newNilai.Nilai >= 60 {
		newNilai.IndeksNilai = "C"
	} else if newNilai.Nilai >= 50 {
		newNilai.IndeksNilai = "D"
	} else {
		newNilai.IndeksNilai = "E"
	}

	// Generate ID
	if len(nilaiNilaiMahasiswa) == 0 {
		newNilai.ID = 1
	} else {
		newNilai.ID = nilaiNilaiMahasiswa[len(nilaiNilaiMahasiswa)-1].ID + 1
	}

	// Add to slice
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)

	// Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNilai)
}

// jawaban soal no 2 ---------------------------------------------->
func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check Basic Auth
	username, password, ok := r.BasicAuth()
	if !ok || username != "user" || password != "pass" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

// BasicAuth middleware
func BasicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
