FROM golang:1.27-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY internal ./internal
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
