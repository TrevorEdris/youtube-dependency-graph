FROM golang:1.17-alpine as builder

RUN apk update && apk add musl-dev gcc build-base

WORKDIR /go/src/github.com/TrevorEdris/youtube-dependency-graph

COPY . .
RUN go test -v ./...
RUN go build -ldflags "-linkmode external -extldflags \"-static\" -s -w $LDFLAGS" -o ydg cmd/ydg/main.go

# Copy the binary from the "builder" docker target into a scratch container
# to vastly reduce the overall size of the image
FROM scratch AS final

ENTRYPOINT ["/ydg"]
COPY --from=builder /go/src/github.com/TrevorEdris/youtube-dependency-graph/ydg /ydg
