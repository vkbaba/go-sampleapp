FROM golang:1.14.3 as builder
COPY ./main.go ./
RUN go build -o /hello ./main.go

FROM alpine:latest
COPY --from=builder /hello .
ENTRYPOINT ["./hello"]
