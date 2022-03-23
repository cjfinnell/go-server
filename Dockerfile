FROM golang:1.18-alpine3.15 as builder
WORKDIR /build
RUN apk update --no-cache && apk add git make --no-cache
COPY . .
RUN go mod vendor && make build

FROM alpine:3.15 as release
WORKDIR /app
RUN apk update --no-cache
COPY --from=builder /build/go-server /app/go-server
EXPOSE 8080
ENTRYPOINT ["/app/go-server"]
CMD ["start"]
