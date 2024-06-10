FROM golang:1.22.1-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

EXPOSE 5000

CMD [ "go", "run", "main.go" ]
