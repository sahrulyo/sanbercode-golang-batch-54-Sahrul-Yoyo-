package endpoint

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Jawaban2() {
	r := gin.Default()

	r.GET("/bangun-datar/segitiga-sama-sisi", func(c *gin.Context) {
		alasStr := c.Query("alas")
		tinggiStr := c.Query("tinggi")
		hitung := c.Query("hitung")

		alas, errAlas := strconv.ParseFloat(alasStr, 64)
		tinggi, errTinggi := strconv.ParseFloat(tinggiStr, 64)

		if errAlas != nil || errTinggi != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter alas dan tinggi harus berupa angka"})
			return
		}

		switch hitung {
		case "luas":
			result := 0.5 * alas * tinggi
			c.JSON(http.StatusOK, gin.H{"result": result})
		case "keliling":
			// Rumus keliling segitiga sama sisi adalah 3 * sisi.
			result := 3 * alas
			c.JSON(http.StatusOK, gin.H{"result": result})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung harus luas atau keliling"})
		}
	})

	r.GET("/bangun-datar/persegi", func(c *gin.Context) {
		sisiStr := c.Query("sisi")
		hitung := c.Query("hitung")

		sisi, err := strconv.ParseFloat(sisiStr, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter sisi harus berupa angka"})
			return
		}

		switch hitung {
		case "luas":
			result := sisi * sisi
			c.JSON(http.StatusOK, gin.H{"result": result})
		case "keliling":
			result := 4 * sisi
			c.JSON(http.StatusOK, gin.H{"result": result})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung harus luas atau keliling"})
		}
	})

	r.GET("/bangun-datar/persegi-panjang", func(c *gin.Context) {
		panjangStr := c.Query("panjang")
		lebarStr := c.Query("lebar")
		hitung := c.Query("hitung")

		panjang, errPanjang := strconv.ParseFloat(panjangStr, 64)
		lebar, errLebar := strconv.ParseFloat(lebarStr, 64)

		if errPanjang != nil || errLebar != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter panjang dan lebar harus berupa angka"})
			return
		}

		switch hitung {
		case "luas":
			result := panjang * lebar
			c.JSON(http.StatusOK, gin.H{"result": result})
		case "keliling":
			result := 2 * (panjang + lebar)
			c.JSON(http.StatusOK, gin.H{"result": result})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung harus luas atau keliling"})
		}
	})

	r.GET("/bangun-datar/lingkaran", func(c *gin.Context) {
		jariJariStr := c.Query("jariJari")
		hitung := c.Query("hitung")

		jariJari, err := strconv.ParseFloat(jariJariStr, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter jari-jari harus berupa angka"})
			return
		}

		switch hitung {
		case "luas":
			result := math.Pi * jariJari * jariJari
			c.JSON(http.StatusOK, gin.H{"result": result})
		case "keliling":
			result := 2 * math.Pi * jariJari
			c.JSON(http.StatusOK, gin.H{"result": result})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung harus luas atau keliling"})
		}
	})

	r.Run(":8002")
}
