FROM golang:1.18.3-alpine3.16 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o action main.go

FROM scratch
COPY --from=builder /app/action /bin/action
ENTRYPOINT ["action"]
