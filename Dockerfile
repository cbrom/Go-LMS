FROM golang:1.14

WORKDIR /Go-LMS

COPY go.mod ./

RUN go mod tidy

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main
