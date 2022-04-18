FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

ADD certs/ca/cert.pem /etc/ssl/certs/
ADD certs/ca/key.pem /etc/ssl/certs/

RUN go build -o /server

FROM alpine

WORKDIR /

COPY --from=builder /etc/ssl/certs/cert.pem /etc/ssl/certs/

COPY --from=builder /etc/ssl/certs/key.pem /etc/ssl/certs/

COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT [ "./server" ]