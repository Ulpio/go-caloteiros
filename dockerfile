# Use a imagem oficial do Go
FROM golang:1.22.4

# Ative o CGO
ENV CGO_ENABLED=1

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos do projeto para o contêiner
COPY . .

# Baixe as dependências do projeto
RUN go mod download

# Compile o projeto
RUN go build -o main .

# Comando para executar o projeto
CMD ["./main"]