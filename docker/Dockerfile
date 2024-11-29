# Use uma versão recente do Golang
FROM golang:1.23-alpine

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos go.mod e go.sum para o container
COPY go.mod ./

# Baixe as dependências do Go
RUN go mod tidy

# Copie o restante do código fonte para o container
COPY . .

# Compile o programa Go
RUN go build -o stress-test .

# Defina o comando padrão para executar a aplicação
ENTRYPOINT ["./stress-test"]
