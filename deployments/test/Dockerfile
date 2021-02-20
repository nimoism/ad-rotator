FROM golang:1.15

WORKDIR /opt
COPY . .
RUN go mod download

ENTRYPOINT ["go"]
CMD ["run", "main.go", "--config=configs/config.compose.test.toml"]
