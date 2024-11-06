package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "user=postgres dbname=restaurant_db sslmode=disable password=02061996"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database:", err)
	}
	fmt.Println("Koneksi ke database berhasil.")
	return db
}
