FROM golang:1.21 as builder

WORKDIR /ewallet

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM golang:1.21

WORKDIR /ewallet/

COPY config.yml  .

COPY --from=builder /ewallet/main .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
