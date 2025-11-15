# Estágio de Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia os arquivos de módulo e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Constrói o executável
# O nome do executável é upskilling-server, conforme o arquivo principal
RUN go build -o /upskilling-server upskilling-server.go
RUN go build -o /app/cmd/start/setup ./cmd/start/setup.go

RUN chmod +x /upskilling-server
RUN chmod +x /app/cmd/start/setup

# Estágio de Execução
FROM alpine:latest

WORKDIR /root/

# Instala dependências necessárias (ex: certificados SSL)
RUN apk --no-cache add ca-certificates

# Copia o executável do estágio de build
COPY --from=builder /upskilling-server .

# Copia o script de setup (seeder) e o esquema SQL
COPY --from=builder /app/cmd/start/setup /setup
COPY --from=builder /app/db/schema.sql /db/schema.sql

# Copia o arquivo de variáveis de ambiente de exemplo para referência
COPY --from=builder /app/.env.example ./.env.example

COPY docs ./docs

# Porta de exposição
EXPOSE 8080

# Comando padrão para rodar a aplicação
CMD ["./upskilling-server"]
