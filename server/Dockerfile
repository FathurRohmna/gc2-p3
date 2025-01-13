FROM golang:1.22-alpine as builder 

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o grpc-server main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/grpc-server /usr/local/bin/grpc-server

EXPOSE 50051
EXPOSE 50052

CMD ["grpc-server"]
