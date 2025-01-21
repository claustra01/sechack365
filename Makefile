# start
up:
	docker compose -f compose.prod.yaml up -d --build


# stop
down:
	docker compose -f compose.prod.yaml down


# setup environment
env:
	sudo apt update -y
	sudo apt upgrade -y
	# check files
	if [ ! -f .env ]; then \
		echo ".env file not found"; \
		exit 1; \
	fi
	if [ ! -f nginx/default.crt ]; then \
		echo "nginx/default.crt file not found"; \
		exit 1; \
	fi
	if [ ! -f nginx/default.key ]; then \
		echo "nginx/default.key file not found"; \
		exit 1; \
	fi
	# install docker
	sudo apt -y install ca-certificates curl gnupg
	sudo install -m 0755 -d /etc/apt/keyrings
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
	sudo chmod a+r /etc/apt/keyrings/docker.gpg
	echo \
	"deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
	"$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
	sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
	sudo apt -y update
	sudo apt -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
	sudo docker run hello-world
	sudo usermod -aG docker $(whoami)
	newgrp docker


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


# backend static analysis
lint-api:
	@if ! command -v $$(go env GOPATH)/bin/golangci-lint > /dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	cd api && $$(go env GOPATH)/bin/golangci-lint run ./... --timeout 5m


# backend test
test-api:
	cd api && go test -v ./...


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


# connect to postgres on shell
dev-psql:
	psql -h 127.0.0.1 -p 5432 -U postgres -d sechack365


# kill postgres process
dev-port-clean:
	sudo kill -9 $$(ps aux | grep postgres | awk '{print $2}')
