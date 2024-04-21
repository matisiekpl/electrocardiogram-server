FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN apk add --no-cache git gcc musl-dev
RUN CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CGO_ENABLED=1 GOOS=linux go build -o /main cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 6000
EXPOSE 6001
COPY server.crt /server.crt
COPY server.key /server.key
ENTRYPOINT ["/main"]