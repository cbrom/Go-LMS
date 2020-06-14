FROM golang:1.14

WORKDIR /Go-LMS

COPY go.mod ./

RUN go mod tidy

RUN go mod download