FROM golang:1.22-alpine

WORKDIR /app

RUN curl -sSf https://atlasgo.sh | sh

RUN go install github.com/air-verse/air@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

COPY go.mod go.sum ./
RUN go mod download

