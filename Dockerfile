# Etapa 1: Construção do binário
FROM golang:1.20 as builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie os arquivos do projeto
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compile o binário como estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Etapa 2: Imagem final
FROM scratch

# Diretório de trabalho
WORKDIR /app

# Copie o binário gerado na etapa anterior
COPY --from=builder /app/main .

# Comando padrão para executar o binário
CMD ["./main"]
