package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	// DB接続
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to init database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Default().Println("success to connect db!!")

	// HTTPサーバーを起動
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to serve: ", err)
	}
	log.Default().Println("Server started on port: 8080")
}
