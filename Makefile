migrate:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go"
	docker compose down

swagger:
	npm install -g swagger-cli
	swagger-cli bundle openapi/all.yaml --outfile openapi/generated.yaml --type yaml

redocly:
	npm install -g @redocly/cli
	redocly build-docs openapi/generated.yaml -o openapi/index.html

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

dev-psql:
	psql -h 127.0.0.1 -p 5432 -U postgres -d sechack365
