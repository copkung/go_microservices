# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /demo

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD /templates ./templates

COPY /allApi/putApi/putOneData.go ./

RUN go build -o /putOne

EXPOSE 8080

CMD [ "/putOne" ]
