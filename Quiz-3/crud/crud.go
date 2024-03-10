package crud

import (
	"database/sql"
	"net/http"
	"strconv"

	"regexp"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

//CRUD Category

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var db *sql.DB

var categories []Category

//GET ALL CATEGORIES

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

// POST -- CREATE CATEGORIES
func CreateCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories = append(categories, category)
	c.JSON(http.StatusCreated, category)
}

// PUT -- UPDATE CATEGORIES
func UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var updatedCategory Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	found := false
	for i, category := range categories {
		if category.ID == id {
			categories[i] = updatedCategory
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DELETE CATEGORIES
func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	found := false
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

// GET ALL BOOKS BY  CATEGORY ID
func GetBooksByCategoryID(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	rows, err := db.Query("SELECT * FROM books WHERE category_id = $1", categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

//CRUD BOOKS

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryID  int       `json:"category_id"`
}

func GetBooks(c *gin.Context) {
	// Implementasi untuk mendapatkan semua buku
	var books []Book
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// CREATE BOOK
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//VALIDASI REGEX
	if matched, _ := regexp.MatchString(`^(http|https)://`, book.ImageURL); !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL format"})
		return
	}

	// VALIDASI RELEASE YEAR
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year should be between 1980 and 2021"})
		return
	}

	//KONVERSI DAN NON-INPUTAN USER
	book.Thickness = calculateThickness(book.TotalPage)

	// TAMBAH TIMESTAMP
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	// INSERT DATA BUKU BARU
	_, err := db.Exec("INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CreatedAt, book.UpdatedAt, book.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

// UPADTE BUKU
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if matched, _ := regexp.MatchString(`^(http|https)://`, book.ImageURL); !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL format"})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year should be between 1980 and 2021"})
		return
	}

	book.Thickness = calculateThickness(book.TotalPage)

	_, err = db.Exec("UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, updated_at=$8, category_id=$9 WHERE id=$10",
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, time.Now(), book.CategoryID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DELETE BUKU
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// DELETE BUKU DARI DATABASE
	_, err = db.Exec("DELETE FROM books WHERE id=$1", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// FUNGSI MENGHITUNG KETEBALAN BUKU
func calculateThickness(totalPage int) string {
	var thickness string
	switch {
	case totalPage <= 100:
		thickness = "tipis"
	case totalPage > 100 && totalPage <= 200:
		thickness = "sedang"
	default:
		thickness = "tebal"
	}
	return thickness
}
