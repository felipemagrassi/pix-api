FROM golang:1.22.4

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/api cmd/api/main.go

EXPOSE 8080

ENTRYPOINT ["/app/api"]
