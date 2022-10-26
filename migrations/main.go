package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/7201-12/ELS/migrations/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	time.Sleep(2 * time.Second)

	connStr := "postgres://postgres:postgres@localhost:5430/els?sslmode=disable"
	migrationsConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("open migrations db connection")
	}
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("set dialect")
	}
	if err := goose.Up(migrationsConn, "."); err != nil {
		fmt.Println(err)
		log.Fatalf("migrate up. %s", connStr)
	}
	_ = migrationsConn.Close()
}
