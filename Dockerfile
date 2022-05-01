FROM golang:1.18-alpine3.15
ENV GO111MODULE=on

WORKDIR $GOPATH/App-Runner-Go

RUN apk update && apk add git

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server.go

CMD ["./server"]