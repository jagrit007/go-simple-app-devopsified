FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o APP

FROM scratch

COPY --from=builder /app/APP .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./APP"]

