# Use uma imagem oficial do Go como base
FROM golang:1.21

# Cria diretório de trabalho
WORKDIR /app

# Copia os arquivos para dentro do container
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compila o binário
RUN go build -o main .

# Expõe a porta da API
EXPOSE 8080

# Comando para rodar o binário
CMD ["./main"]

