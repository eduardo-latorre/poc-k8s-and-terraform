FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o app .

FROM golang:1.19 AS deploy

WORKDIR /app

COPY --from=builder app .

CMD ["./app"]