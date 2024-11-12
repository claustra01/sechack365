package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		panic("invalid argument")
	}
	switch os.Args[1] {
	case "drop":
		drop(conn)
	case "migrate":
		migrate(conn)
	case "mock":
		mock(conn)
	}
}

func drop(conn *sql.DB) {
	migrateSql, err := os.ReadFile("cmd/database/drop.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	fmt.Println("database dropped")
}

func migrate(conn *sql.DB) {
	migrateSql, err := os.ReadFile("cmd/database/migrate.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	fmt.Println("migration successed")
}

func mock(conn *sql.DB) {
	mockSql, err := os.ReadFile("cmd/database/mock.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	replacedMockSql := strings.ReplaceAll(string(mockSql), "$hashed_password$", hex.EncodeToString(sha256.New().Sum([]byte("password"))))
	if _, err = conn.Exec(replacedMockSql); err != nil {
		panic(err)
	}
	fmt.Println("mock data inserted")
}
