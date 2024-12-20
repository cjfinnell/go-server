FROM golang:1.22-alpine3.18 as builder
WORKDIR /build
RUN apk update --no-cache && apk add git make --no-cache
COPY . .
RUN go mod vendor && make build

FROM alpine:3.21 as release
WORKDIR /app
RUN apk update --no-cache
COPY --from=builder /build/go-server /app/go-server
EXPOSE 8080
ENTRYPOINT ["/app/go-server"]
CMD ["start"]
