FROM golang:1.19-alpine

WORKDIR /go/src/bank-transaction

ADD . .

RUN go build -o /bank-transaction cmd/main.go

ENTRYPOINT ["/bank-transaction"]