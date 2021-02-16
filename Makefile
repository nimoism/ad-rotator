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
	@go test -v -count=1 ./...

build: install-deps
	@go build -v -o $(BIN) ./cmd/

migrate: install-deps
	@[ "$$DSN" ] || (echo "\$$DSN isn't defined" && exit 1)
	@goose -dir migrations postgres $$DSN up

run: build
	@$(BIN) --config=$(CONFIG)

docker-build:
	@docker build -t ad-rotator -f deployments/Dockerfile .

up:
	@

.PHONY: install-deps lint test build migrate  run
