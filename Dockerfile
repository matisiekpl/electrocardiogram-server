FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -o /main cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 6000
EXPOSE 6001
ENTRYPOINT ["/main"]