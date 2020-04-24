FROM golang:latest

WORKDIR /

ENTRYPOINT mkdir -p go-gin-todoServer

COPY . ./go-gin-todoServer

WORKDIR /go-gin-todoServer

ENTRYPOINT ["go", "run", "src/main.go"]

EXPOSE 8080