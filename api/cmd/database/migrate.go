package main

import (
	"fmt"
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
	_, err := conn.Pg.Exec(`
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS ap_identity;
		CREATE TABLE users (
			id VARCHAR(255),
			user_id VARCHAR(255) NOT NULL,
			encrypted_password VARCHAR(255) NOT NULL,
			display_name VARCHAR(255) NOT NULL,
			profile VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);
		CREATE TABLE ap_identity (
			user_id VARCHAR(255) NOT NULL,
			public_key VARCHAR(255) NOT NULL,
			secret_key VARCHAR(255) NOT NULL,
			PRIMARY KEY (user_id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("migration success")
}
