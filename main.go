package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var db *sql.DB
var rdb *redis.Client

func main() {
	// Database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/{short}", redirectHandler) // Note: This is simplified; use gorilla/mux for real routing

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement URL shortening logic
	fmt.Fprintf(w, "URL Shortener API - /shorten endpoint")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement redirect logic
	http.Redirect(w, r, "https://example.com", http.StatusMovedPermanently)
}
