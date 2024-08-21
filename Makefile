migrate:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go"
	docker compose down

dev-cert:
	sudo apt install mkcert -y
	mkcert -install
	mkcert -cert-file ./nginx/default.crt -key-file ./nginx/default.key localhost

dev-mock:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go mock"
	docker compose down

dev-port-clean:
	sudo kill -9 $(ps aux | grep postgres | awk '{print $2}')
