FROM golang:1.19.3-alpine3.16 as builder

RUN adduser -D deployer

WORKDIR /go/src/beach

COPY src .

RUN go mod download

RUN CGO_ENABLED=0 go build -o build/beach-api cmd/main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/src/beach/build/beach-api .

USER deployer

CMD ["./beach-api"]
