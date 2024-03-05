package main

import (
	"fmt"
	"math"
	"net/http"
)

func hitungVolume(jariJari, tinggi float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) * tinggi
}

func hitungLuasAlas(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2)
}

func hitungKelilingAlas(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func handler(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.0
	tinggi := 10.0

	volume := hitungVolume(jariJari, tinggi)
	luasAlas := hitungLuasAlas(jariJari)
	kelilingAlas := hitungKelilingAlas(jariJari)

	fmt.Fprintf(w, "jariJari: %.2f, tinggi: %.2f, volume: %.2f, luas alas: %.2f, keliling alas: %.2f",
		jariJari, tinggi, volume, luasAlas, kelilingAlas)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
