ARG GO_VER=1.24

FROM golang:${GO_VER}-alpine AS builder
WORKDIR /build
RUN apk update --no-cache && apk add git make --no-cache
COPY . .
RUN go mod vendor && make build

FROM scratch AS release
COPY --from=builder /build/go-server /go-server
EXPOSE 8080
ENV PATH=/
ENTRYPOINT ["/go-server"]
CMD ["start"]
