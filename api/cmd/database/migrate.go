package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	migrate(conn)

	if len(os.Args) > 1 && os.Args[1] == "mock" {
		insertMock(conn)
	}
}

func migrate(conn *sql.DB) {
	_, err := conn.Exec(`
		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS ap_user_identifiers;

		CREATE TABLE users (
				id VARCHAR(255) PRIMARY KEY,
				username VARCHAR(255) NOT NULL UNIQUE,
				host VARCHAR(255) NOT NULL,
				encrypted_password VARCHAR(255) NOT NULL,
				display_name VARCHAR(255),
				profile TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE ap_user_identifiers (
				user_id VARCHAR(255) PRIMARY KEY,
				public_key TEXT NOT NULL,
				private_key TEXT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("migration success")
}

func insertMock(conn *sql.DB) {
	conn.Exec(`
		INSERT INTO users (id, username, host, encrypted_password, display_name, profile)
			VALUES ('1', 'mock', 'localhost', 'password', 'mock user', 'This is mock user');
		INSERT INTO ap_user_identifiers (user_id, public_key, private_key)
			VALUES ('1', '-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAykPAEF/84PZaGUc3b8GR
5df/COT+G8Mjm5/xw2Eyqo6zsbSTt7RyN4xAfl8i1yILUvCM0LkTIKjw+AAXWC+L
gPpZVGESn1JqH/MrpwmVvkauMJtC9/h3DIAUOvVPbSgar4JUM90KmN9iZi2XIajp
wy/pqjWnu5lDIl/DERraEj6nZ5AsZwDWrtqm7JA9kTxVUVMp4i+gSUytiLZRY316
UdwaJPPFSnYJZ2+HsnXJWcheFPE+dxIwERo37z11UhD9C7fAZePHkJBK2xjhZBUJ
EDJJqMNZXtSWvYjJCDrL3/a5eaX0NcBDR/jwXZcni6D+JZn3GnADlsO0dh8dS9Zb
AwIDAQAB
-----END PUBLIC KEY-----
', '-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAykPAEF/84PZaGUc3b8GR5df/COT+G8Mjm5/xw2Eyqo6zsbST
t7RyN4xAfl8i1yILUvCM0LkTIKjw+AAXWC+LgPpZVGESn1JqH/MrpwmVvkauMJtC
9/h3DIAUOvVPbSgar4JUM90KmN9iZi2XIajpwy/pqjWnu5lDIl/DERraEj6nZ5As
ZwDWrtqm7JA9kTxVUVMp4i+gSUytiLZRY316UdwaJPPFSnYJZ2+HsnXJWcheFPE+
dxIwERo37z11UhD9C7fAZePHkJBK2xjhZBUJEDJJqMNZXtSWvYjJCDrL3/a5eaX0
NcBDR/jwXZcni6D+JZn3GnADlsO0dh8dS9ZbAwIDAQABAoIBAQC12iIV1ud6r6Ok
NJaQIR524zOGoLQi69jY8/4fJwWxuSmwrWVedpt2e+AEfq7Jc+9we5xvkOa0p5A3
uYVDoUOxC+VC6yAeJLAL18s4nHKIp+22//E/F6KZl5IYzDPENZmAkRH5q1P6zGUg
7v6BoefCuRJCGYmcLpjgj+7HMzg7y4W96SDQ4xZgWpTj+iicgE2FzYRaQo150bb3
u1grUO3AoP+i05piXUf57a5fRT+z+4IYb7TaCQtCGM1fp6tSVkFGuH26j0zNLvmi
D9VvmlmTGgUpuSl2x6GI0v24GiD63lEM5XRSkjFrSn+osww4gr2AH9DFgt6kIkQi
gjBUpgUpAoGBANxkZOPhUeCNYYJpM/HkbEuPDktMFkjFwHAD8aYHmei6JCOvyQuH
XL6gFa8BNx/Ry3qr0RtmGQtEw8xYVbNKf3go2msiK5PjzkZ6s2WH6ufgMQxCpUlU
+huwTymO5pZxbfzeNWMujCgk6jLpI+me49uo5xiIh4hii0yDmQau39RXAoGBAOrx
ljplzYQo7uPuiMK0/VpNzlr4aiCiiq+TrX/woLINsKdj3Y4Tpm5X17N40E//+14L
BoqXgX/ZNeTnOqQXLR/goMrvVU2NyA6U/IYGTMlCKtEoRjO8Aoc/3TSZcxpqrvhV
C6KKlErVtKq7oCPPm6r1JOdrKOPcwRH+ZFgXlKM1AoGBAKC2+TeQRPPaRaQi8YVQ
zIQhEwxntMxmoJlO1vX7Dwo+S0JW2uX0VPaRqJ5Q5ZDnnVmcV8WCI3srLxkhxYUU
K3ZFXFnJtjuHYRHWQmIkxnFG9J17MCsUs7pjTKcClTZaCxneNNJZzE0t9jcf+ldP
zduOBM/IKAWVzv0B7iKIfaLLAoGAXX6GKfcZMd6YMlxaUCF2MNmFpO32TcZhKj26
bY90Y2bPRc2X/VIUiRSr4d/SBgP4JBR/JefkwNvPdqgNzf7rFiRt2FQlvhcN5b+k
PjGDnROXtmQwi6Xl26yueqAWDg0mU+yEFMrQ+HbSzp6bu6SCbiXf6bfbLdJLgr2Z
cPTxUYECgYBwJmDbNKQTC833OQHW5sM+FwgDCxqutWnW1MyBAIOraVUzUgTZdARx
5b4BkDwCAa6LLZw64ya5+DvALI3DZStNj2PNqKZnRsBFYmpnTqh/rO7T32IeNCzv
Nr0I9tn7GcHnN/yryv7yAJ9WzGqKABIsucjxysnVpJVc30LfDUQ61Q==
-----END RSA PRIVATE KEY-----
');
	`)
	fmt.Println("mock data inserted")
}
