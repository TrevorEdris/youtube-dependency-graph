FROM golang:1.17-alpine

RUN apk update && apk add musl-dev gcc build-base

RUN go get github.com/cosmtrek/air

WORKDIR /app

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
