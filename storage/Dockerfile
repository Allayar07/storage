FROM golang:latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./ ./

RUN go build -o storage ./cmd/main.go

CMD [".storage"]