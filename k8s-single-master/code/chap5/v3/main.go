package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var resCount = 0

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		if resCount > 2 {
			c.String(http.StatusInternalServerError, "Application v3 have some internal error has occurred!\n")
			return
		}
		resCount++
		c.String(http.StatusOK, "Hello application v3\n")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
