package db

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewDB() *sql.DB {
	if e := godotenv.Load(); e != nil {
		fmt.Println(e)
	}

	host := os.Getenv("db-host")
	dbName := os.Getenv("db-name")
	dbPort := os.Getenv("db-port")
	dbUser := os.Getenv("db-user")
	dbPass := os.Getenv("db-password")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, dbPort, dbUser, dbName, dbPass)
	connConfig, err := pgx.ParseConfig(dbUri)
	if err != nil {
		fmt.Println(err)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	conn, err := sql.Open("pgx", connStr)
	if err = conn.Ping(); err != nil {
		log.Panic(err)
	} else {
		fmt.Println("\033[32m","DB connection is established to PostgreSQL!")
	}

	return conn
}