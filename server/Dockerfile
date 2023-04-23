FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY src/ .
RUN go mod download


RUN go build -o server

EXPOSE 8000

CMD ["./server"]