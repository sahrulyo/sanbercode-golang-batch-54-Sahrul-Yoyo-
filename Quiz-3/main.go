package main

import (
	"database/sql"
	"fmt"
	"log"
	"quiz-3/auth"
	"quiz-3/crud"
	"quiz-3/endpoint"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

// jawaban soal no 2 detailnya ada di endpoint.go ----------------------------------------------->2
func main() {
	endpoint.Jawaban2()

	//jawaban soal no 3 detailnya ada di crud.go ----------------------------------------------->3
	var err error
	// Inisialisasi koneksi ke database
	db, err = sql.Open("postgres", "host=localhost port=5432 user=user password=Ulyasar10389# dbname=user sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	router := gin.Default()

	//ROUTES CATEGORIES
	router.GET("/categories", crud.GetCategories)
	router.POST("/categories", crud.CreateCategory)
	router.PUT("/categories/:id", crud.UpdateCategory)
	router.DELETE("/categories/:id", crud.DeleteCategory)
	router.GET("/categories/:id/books", crud.GetBooksByCategoryID)

	//ROUTES BOOKS
	router.GET("/books", crud.GetBooks)
	router.POST("/books", crud.CreateBook)
	router.PUT("/books/:id", crud.UpdateBook)
	router.DELETE("/books/:id", crud.DeleteBook)

	//START SERVER
	err = router.Run(":8005")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	//TRIGER
	if _, err := db.Exec(`
        CREATE TABLE books (
            id          SERIAL PRIMARY KEY,
            Title       VARCHAR(255) NOT NULL,
            Description TEXT,
            ImageURL    VARCHAR(255),
            ReleaseYear INT,
            Price       DECIMAL(10, 2), 
            TotalPage   INT,
            Thickness   VARCHAR(50),
            CreatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UpdatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            CategoryID  INT,
            FOREIGN KEY (CategoryID) REFERENCES categories(id) 
        )`); err != nil {
		log.Fatal(err)
	}

	// FUNC TRIGERNYA
	if _, err := db.Exec(`
        CREATE OR REPLACE FUNCTION update_timestamp()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW."UpdatedAt" = CURRENT_TIMESTAMP;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;

        CREATE TRIGGER update_books_timestamp
        BEFORE UPDATE ON books
        FOR EACH ROW
        EXECUTE FUNCTION update_timestamp();
    `); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tabel books dan fungsi trigger berhasil dibuat.")

	//jawaban soal 4 detailnya ada di auth.go -------------------------------------------------------------->4

	auth.Auth()

}
