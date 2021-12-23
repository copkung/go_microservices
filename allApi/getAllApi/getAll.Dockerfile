# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /demo

ADD /templates ./templates

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /allApi/getAllApi/getAllData.go ./

RUN go build -o /getAll

EXPOSE 8080

CMD [ "/getAll" ]

