package main

import (
	"os"

	"github.com/claustra01/sechack365/pkg/database"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	conn, err := database.Connect(connStr)
	if err != nil {
		panic(err)
	}
	migrate(conn)
}

func migrate(conn *database.SqlHandler) {
	conn.Pg.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255),
		username VARCHAR(255) NOT NULL,
		encrypted_password VARCHAR(255) NOT NULL,
		display_name VARCHAR(255) NOT NULL,
		profile VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
	);`)
}
