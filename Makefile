# backend static analysis
lint-api:
	@if ! command -v $$(go env GOPATH)/bin/golangci-lint > /dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/golangci-lint run ./...

# db migration
migrate:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go"
	docker compose down

# generate openapi docs
redocly:
	@if ! command -v redocly > /dev/null 2>&1; then \
		npm install -g @redocly/cli; \
	fi
	redocly bundle openapi/all.yaml -o openapi/generated.yaml
	redocly build-docs openapi/generated.yaml -o openapi/index.html

# generate schema code from openapi
oapi-codegen:
	@if ! command -v $$(go env GOPATH)/bin/oapi-codegen > /dev/null 2>&1; then \
		go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/oapi-codegen -generate types -package openapi -o ./pkg/openapi/types.gen.go ../openapi/all.yaml

# generate certificate for localhost
dev-cert:
	@if ! command -v mkcert > /dev/null 2>&1; then \
		sudo apt install mkcert -y; \
	fi
	mkcert -install
	mkcert -cert-file ./nginx/default.crt -key-file ./nginx/default.key localhost

# db migration with mock data
dev-mock:
	docker compose -f compose.dev.yml up -d --build
	docker compose exec api sh -c "go run cmd/database/migrate.go mock"
	docker compose down

# kill postgres process
dev-port-clean:
	sudo kill -9 $$(ps aux | grep postgres | awk '{print $2}')

# connect to postgres on shell
dev-psql:
	psql -h 127.0.0.1 -p 5432 -U postgres -d sechack365
