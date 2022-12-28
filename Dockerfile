FROM golang:1.19-alpine AS builder
COPY ./* /app/
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates \
    && cd /app \
    && CGO_ENABLED=0 go build -o bin ./...

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/bin /app
EXPOSE 10051
ENTRYPOINT ["/app"]
