FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download && go mod verify

COPY . .

RUN go build -v -o main .

EXPOSE 3000

CMD ["/app/main"]
