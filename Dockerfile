FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]