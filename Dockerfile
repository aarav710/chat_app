# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /chatapp

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /backend

EXPOSE 8080

CMD [ "/backend" ]