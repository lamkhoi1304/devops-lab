package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello env\n")
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}
