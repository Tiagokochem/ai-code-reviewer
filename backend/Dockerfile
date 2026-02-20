# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

# Copiar binário do builder
COPY --from=builder /app/main .

# Expor porta
EXPOSE 8080

# Comando padrão (pode ser sobrescrito no docker-compose)
CMD ["./main"]
