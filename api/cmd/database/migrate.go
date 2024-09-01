package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	drop(conn)
	migrate(conn)

	if len(os.Args) > 1 && os.Args[1] == "mock" {
		insertMock(conn)
	}
}

func drop(conn *sql.DB) {
	migrateSql, err := ioutil.ReadFile("cmd/database/drop.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	fmt.Println("database dropped")
}

func migrate(conn *sql.DB) {
	migrateSql, err := ioutil.ReadFile("cmd/database/migrate.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	fmt.Println("migration success")
}

func insertMock(conn *sql.DB) {
	migrateSql, err := ioutil.ReadFile("cmd/database/mock.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	fmt.Println("mock inserted")
}
