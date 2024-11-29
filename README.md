# Stress Test CLI

Este projeto é uma ferramenta de linha de comando (CLI) desenvolvida em Go para realizar testes de carga em um serviço web. 

O sistema permite que você especifique uma URL de serviço, o número total de requisições a serem feitas e o número de requisições simultâneas. 

Ele gera um relatório com o desempenho do serviço, incluindo o status das requisições HTTP.

## Requisitos

- Go 1.23 ou superior
- Docker (se preferir executar dentro de um contêiner)

## Como Executar

### **Executar Localmente (sem Docker)**

Se você deseja executar a ferramenta diretamente no seu ambiente local (sem Docker), siga os passos abaixo.

#### **Baixe o repositório**
Clone o repositório para sua máquina local:

```bash
https://github.com/lucastg/stress-test.git
cd stress-test
```

### **Instale as dependências (opcional)**
```bash
go mod tidy
```

### **Execute o programa**
```bash
go run main.go --url=http://google.com --requests=100 --concurrency=10
```
> **_Obs:_**  Você pode alterar os parâmetros para personalizar os testes, como:
>* A URL de destino (--url).
>* O número total de requisições (--requests).
>* O número de requisições simultâneas (--concurrency).
>* Caso queira alterar a URL de destino para um serviço diferente, basta passar um novo valor para o parâmetro --url.

### **Exemplo de Saída:**
```
========== Relatório de Testes de Carga ==========
URL Testada: http://google.com
Tempo Total de Execução: 10.20s
Total de Requests Enviados: 100

Distribuição dos Status HTTP:
  Status 200: 90 requests
  Status 404: 5 requests
  Status 500: 5 requests

Total de Requests Bem-Sucedidos (Status 200): 90
==================================================
```

# **Executar com Docker**
Você também pode executar o sistema dentro de um contêiner Docker. Para isso, siga os passos abaixo.

## **Construa a imagem Docker**
```bash
docker build -t stress-test -f docker/Dockerfile .
```
## **Construa a imagem Docker**
```bash
go run main.go --url=http://google.com --requests=100 --concurrency=10
```
> **_Obs:_**  Você pode alterar os parâmetros para personalizar os testes, como:
>* A URL de destino (--url).
>* O número total de requisições (--requests).
>* O número de requisições simultâneas (--concurrency).
>* Caso queira alterar a URL de destino para um serviço diferente, basta passar um novo valor para o parâmetro --url.

# **Licença**
Este projeto está licenciado sob a [Licença MIT](LICENSE).
