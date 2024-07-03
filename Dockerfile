FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o web-scraper-adv main.go

ENTRYPOINT ["/app/web-scraper-adv"]