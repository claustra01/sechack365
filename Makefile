# start
up:
	docker compose -f compose.prod.yaml up -d --build

# stop
down:
	docker compose -f compose.prod.yaml down

# backend static analysis
lint-api:
	@if ! command -v $$(go env GOPATH)/bin/golangci-lint > /dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/golangci-lint run ./...
	cd frontend && pnpm lint

# lint all
lint:
	make lint-api
	cd frontend && pnpm lint

# db migration
migrate:
	docker compose up -d
	docker compose exec api sh -c "go run cmd/database/migrate.go migrate"
	docker compose down

# db drop
drop:
	docker compose up -d
	docker compose exec api sh -c "go run cmd/database/migrate.go drop"
	docker compose down

# refresh tables
refresh-tables: drop migrate

# generate openapi docs
redocly:
	@if ! command -v redocly > /dev/null 2>&1; then \
		npm install -g @redocly/cli; \
	fi
	redocly bundle openapi/all.yaml -o openapi/all.gen.yaml
	redocly build-docs openapi/all.gen.yaml -o openapi/index.html

# generate schema code from openapi
oapi-codegen:
	@if ! command -v $$(go env GOPATH)/bin/oapi-codegen > /dev/null 2>&1; then \
		go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/oapi-codegen -generate types,spec -package openapi -o ./pkg/openapi/types.gen.go ../openapi/all.gen.yaml

# generate hooks and types from openapi
orval:
	cd frontend && pnpm orval && pnpm orval:lint

# generate all openapi files
openapi: redocly oapi-codegen orval

# generate sql from dbdocs
schema:
	@if ! command -v dbml2sql > /dev/null 2>&1; then \
		npm install -g @dbml/cli; \
	fi
	dbml2sql dbdocs/schema.dbml --postgres > api/cmd/database/schema.sql
	sed -i "s/timestamp DEFAULT 'CURRENT_TIMESTAMP'/timestamp with time zone DEFAULT CURRENT_TIMESTAMP/g" api/cmd/database/schema.sql

# generate dbdocs token
dbdocs-token:
	@if ! command -v dbdocs > /dev/null 2>&1; then \
		npm install -g dbdocs; \
	fi
	dbdocs login
	dbdocs token -g

# generate certificate for localhost
dev-cert:
	@if ! command -v mkcert > /dev/null 2>&1; then \
		sudo apt install mkcert -y; \
	fi
	mkcert -install
	mkcert -cert-file ./nginx/default.crt -key-file ./nginx/default.key localhost

# kill postgres process
dev-port-clean:
	sudo kill -9 $$(ps aux | grep postgres | awk '{print $2}')

# connect to postgres on shell
dev-psql:
	psql -h 127.0.0.1 -p 5432 -U postgres -d sechack365
