FROM golang  AS build

EXPOSE 8080  
WORKDIR /app

COPY . /app
RUN go get -d ./...
RUN go build -o server .

FROM scratch AS bin
COPY --from=build /app .
CMD ["./server"]