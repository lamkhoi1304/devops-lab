package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "POSTGRES_USER",
		Password: "POSTGRES_PASSWORD",
		Addr:     "POSTGRES_HOST:POSTGRES_PORT",
		Database: "POSTGRES_DB",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Kube\n")
	})

	Connect()
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}
