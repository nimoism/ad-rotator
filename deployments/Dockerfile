FROM golang:1.15

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /opt/
COPY . /opt/
RUN go build -a -o ad-rotator ./main.go

FROM alpine:latest
WORKDIR /opt/
COPY --from=0 /opt/ad-rotator /bin/
COPY Makefile .
COPY migrations /opt/migrations

ENTRYPOINT ["/bin/ad-rotator"]
CMD ["--config=configs/config.toml"]
