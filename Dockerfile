FROM golang:1.14

WORKDIR /Go-LMS

COPY go.mod ./

RUN go mod tidy

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN go get github.com/joho/godotenv/cmd/godotenv

CMD CompileDaemon --build="go build cmd/graphql/main.go" --command="godotenv -f .env ./main"
