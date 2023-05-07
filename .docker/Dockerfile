FROM golang:1.20-alpine

RUN apk update && apk add gcc g++ ca-certificates git
WORKDIR /opt

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

CMD ["air"]