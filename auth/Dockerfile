FROM golang:1.22-alpine as builder
WORKDIR /app
COPY auth/go.mod auth/go.sum ./
COPY auth/ ./
COPY common ../common
WORKDIR /app/cmd
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth

FROM alpine as runtime
COPY --from=builder /auth /auth
ENTRYPOINT ["/auth"]
