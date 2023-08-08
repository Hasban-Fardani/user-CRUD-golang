package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func Connect() *sql.DB {
	var db *sql.DB
	var err error

	user := getEnv("DB_USER", "root")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	database := getEnv("DB_DATABASE", "user_crud_go")
	connStr := fmt.Sprintf("%s@tcp(%s:%s)/%s", user, host, port, database)

	db, err = sql.Open("mysql", connStr)
	err = db.Ping()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Koneksi ke database berhasil")
	}

	return db
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}
}
