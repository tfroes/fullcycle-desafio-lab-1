FROM golang:1.23.11-alpine as builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o server ./cmd/.

# FROM scratch
FROM alpine:latest
# FROM golang:latest
# WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080 

CMD ["./server"]
