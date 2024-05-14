FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server cmd/server/*.go

FROM gcr.io/distroless/base
COPY --from=builder /server /server

CMD ["/server"]
