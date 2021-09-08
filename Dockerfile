FROM golang:1.17

WORKDIR $GOPATH/src/aveplen/avito_qa_test
COPY . .
RUN go mod download
RUN go build -o avito_qa ./cmd/main.go

CMD ["./avito_qa", "-f", "message.json"]