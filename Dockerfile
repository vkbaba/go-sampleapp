FROM golang:1.14.3 as builder
WORKDIR /src
COPY . .
RUN go build -o /hello .

FROM alpine:latest
COPY --from=builder /hello /
ENTRYPOINT ["./hello"]
