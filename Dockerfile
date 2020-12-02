FROM golang  AS build
EXPOSE 8080  
WORKDIR /app
 
COPY . /app
RUN go get -d ./...
RUN go build -o server .
CMD ["./server"]
