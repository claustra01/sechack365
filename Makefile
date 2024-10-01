lint-api:
	@if ! command -v $$(go env GOPATH)/bin/golangci-lint > /dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/golangci-lint run ./...

migrate:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go"
	docker compose down

swagger:
	@if ! command -v swagger-cli > /dev/null 2>&1; then \
		npm install -g swagger-cli; \
	fi
	swagger-cli bundle openapi/all.yaml --outfile openapi/generated.yaml --type yaml

redocly:
	@if ! command -v redocly > /dev/null 2>&1; then \
		npm install -g @redocly/cli; \
	fi
	redocly build-docs openapi/generated.yaml -o openapi/index.html

dev-cert:
	@if ! command -v mkcert > /dev/null 2>&1; then \
		sudo apt install mkcert -y; \
	fi
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
