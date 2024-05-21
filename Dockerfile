FROM golang:latest
WORKDIR  /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /bin/hello ./hello/hello.go
CMD ["hello"]