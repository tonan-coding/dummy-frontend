FROM golang:1.24.2-alpine

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /dummy-frontend

CMD ["/dummy-frontend"]

