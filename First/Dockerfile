FROM golang:latest
WORKDIR /app
RUN mkdir -p /app/tmp
RUN ulimit -n 2000
RUN go get github.com/google/uuid
RUN go get github.com/minio/minio-go
COPY . .
RUN go build -o main .
CMD ["./main"]
