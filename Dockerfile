FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /main cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 6000
EXPOSE 6001
ENTRYPOINT ["/main"]