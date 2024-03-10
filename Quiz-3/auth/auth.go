package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	})
}

func Auth() {
	r := gin.Default()
	r.Use(BasicAuth())

	// ROUTE POST
	r.POST("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "POST request is authenticated"})
	})

	// ROUTE PUT
	r.PUT("/editor", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PUT request is authenticated"})
	})

	// ROUTE DELETE
	r.DELETE("/supereditor", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DELETE request is authenticated"})
	})

	r.Run(":8000")
}
