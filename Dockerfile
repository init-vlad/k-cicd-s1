FROM alpine:latest

WORKDIR /app

COPY app .

ENTRYPOINT ["./app"]
