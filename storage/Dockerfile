FROM golang:latest

COPY ./ ./
RUN go build -v main ./cmd
CMD [".cmd/main"]