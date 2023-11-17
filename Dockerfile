FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod tidy

ADD ./ /app

RUN go generate ./...

RUN go build -o index

ENTRYPOINT ["/app/index"]
