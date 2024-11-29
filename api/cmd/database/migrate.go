package main

import (
	"database/sql"
	"fmt"
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

	if len(os.Args) != 2 {
		panic("invalid argument")
	}
	switch os.Args[1] {
	case "drop":
		drop(conn)
	case "migrate":
		migrate(conn)
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
	// migrate
	migrateSql, err := os.ReadFile("cmd/database/schema.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(migrateSql)); err != nil {
		panic(err)
	}
	// triggers
	triggerSql, err := os.ReadFile("cmd/database/trigger.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}
	if _, err = conn.Exec(string(triggerSql)); err != nil {
		panic(err)
	}
	fmt.Println("migration successed")
}
