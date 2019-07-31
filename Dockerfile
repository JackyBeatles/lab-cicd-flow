FROM golang:1.11-alpine AS build-go

WORKDIR /usr/local/go/src/simple-http-go

COPY . .

RUN go build -o http-serv

FROM alpine:3.10

WORKDIR /app

COPY --from=build-go /usr/local/go/src/simple-http-go/http-serv ./

EXPOSE 9000

ENTRYPOINT ["./http-serv"]