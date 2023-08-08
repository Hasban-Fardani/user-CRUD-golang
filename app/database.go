package app

import (
	"crypto/tls"
	"database/sql"

	"github.com/go-sql-driver/mysql"

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
	password := getEnv("DB_PASSWORD", "")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	database := getEnv("DB_DATABASE", "user_crud_go")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=tidb", user, password, host, port, database)

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-west-2.prod.aws.tidbcloud.com",
	})

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
