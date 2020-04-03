FROM golang:1.14.1-alpine3.11 as build
ARG VERSION

RUN apk add g++ libc-dev

RUN mkdir /app
WORKDIR /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download

COPY main.go /app/main.go
RUN go build -ldflags "-linkmode external -extldflags -static -X main.version=${VERSION}"

FROM scratch
ARG PORT
COPY --from=build /app/hello-golang /app
EXPOSE $PORT
ENTRYPOINT ["/app"]