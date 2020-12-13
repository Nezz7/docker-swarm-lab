FROM golang  AS builder

EXPOSE 8080  
WORKDIR /app

COPY . /app
RUN go get -d ./...
RUN go build -o server .

FROM alpine:latest AS bin
COPY --from=builder /app .
CMD ["./server"]