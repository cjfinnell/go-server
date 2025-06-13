ARG GO_VER=1.24
ARG AL_VER=3.22

FROM golang:${GO_VER}-alpine${AL_VER} AS builder
WORKDIR /build
RUN apk update --no-cache && apk add git make --no-cache
COPY . .
RUN go mod vendor && make build

FROM alpine:${AL_VER} AS release
WORKDIR /app
RUN apk update --no-cache
COPY --from=builder /build/go-server /app/go-server
EXPOSE 8080
ENTRYPOINT ["/app/go-server"]
CMD ["start"]
