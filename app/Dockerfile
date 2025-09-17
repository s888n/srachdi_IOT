FROM cgr.dev/chainguard/go AS builder
WORKDIR /app

COPY main.go .
COPY static ./static

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

FROM cgr.dev/chainguard/static
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

# ENV POD_NAME="Pod-123" NODE_NAME="Srachdi" APP_NAME="App 1" 

EXPOSE 8888

ENTRYPOINT [ "./server" ]