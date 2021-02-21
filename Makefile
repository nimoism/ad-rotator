.ONESHELL:

BIN := './bin/ad-rotator'
CONFIG ?= './configs/config.toml'

install-deps:
	@(which goose > /dev/null) || go get github.com/pressly/goose/cmd/goose

generate:
	@protoc ./api/banners.proto --go_out=./internal/api/grpc/pb --go-grpc_out=./internal/api/grpc/pb
	@protoc ./api/slots.proto --go_out=./internal/api/grpc/pb --go-grpc_out=./internal/api/grpc/pb
	@protoc ./api/user_group.proto --go_out=./internal/api/grpc/pb --go-grpc_out=./internal/api/grpc/pb

lint: install-deps
	@golangci-lint run ./... || true

test: install-deps
	@go test -v -race -count=100 ./...

build: install-deps
	@go build -v -o $(BIN) ./main.go

migrate: install-deps
	@[ "$$DSN" ] || (echo "\$$DSN isn't defined" && exit 1)
	@goose -dir migrations postgres $$DSN up

run: build
	@$(BIN) --config=$(CONFIG)

docker-build:
	@docker build -t ad-rotator -f deployments/Dockerfile .

up:
	@docker-compose -f deployments/docker-compose.yml -p adr up

down:
	@docker-compose -f deployments/docker-compose.yml -p adr down

uptest:
	@docker-compose -f deployments/test/docker-compose.yml -p adr-test up -d --build
	docker-compose -f deployments/test/docker-compose.yml -p adr-test exec adr /bin/true || exit 1
	docker-compose -f deployments/test/docker-compose.yml -p adr-test exec adr go run main.go migrate --config=/opt/configs/config.compose.test.toml
	docker-compose -f deployments/test/docker-compose.yml -p adr-test exec adr go test -v -count=1 ./...
	make downtest
	exit $$code

downtest:
	@docker-compose -f deployments/test/docker-compose.yml -p adr-test down || true


.PHONY: install-deps lint test build migrate run docker-build up uptest downtest
