FROM golang:1.22.1-alpine

WORKDIR /backend

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

EXPOSE 5000

CMD [ "go", "run", "main.go" ]
