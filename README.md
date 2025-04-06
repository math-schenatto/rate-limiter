# Rate Limiter em Go 🚦

Este é um projeto de **Rate Limiter** construído em Go, utilizando Redis para armazenamento das métricas de requisição por IP e token. Ele bloqueia requisições quando o limite é excedido por um determinado tempo.

## ✨ Funcionalidades

- Limita requisições por **IP** e/ou **token** (API_KEY)
- Usa **Redis** para controle de requisições com expiração automática
- Responde com **HTTP 429** quando o limite é excedido
- Suporte a configuração via variáveis de ambiente
- Middleware pronto para integrar com seu servidor HTTP

## 🛠️ Tecnologias

- [Go](https://golang.org/)
- [Redis](https://redis.io/)
- [Docker + Docker Compose](https://docs.docker.com/compose/)

## 📦 Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/math-schenatto/rate-limiter.git
cd rate-limiter
```
## 2. Suba os containers

```
docker compose up --build

```
Esse comando inicia:
- Um container com o app Go
- Um container com o Redis na porta 6379


O app fica disponível em: http://localhost:8080


## Testando com curl

```
curl -H "API_KEY: abc123" http://localhost:8080/
```
Se você ainda não atingiu o limite, receberá:
```
✅ Hello! You're not rate-limited.

```

Ao exceder:
```
❌ you have reached the maximum number of requests or actions allowed within a certain time frame

```

## Testando com hey

Voce deve instalar o hey:
```
go install github.com/rakyll/hey@latest

```

Envie 200 requisições, 1 por vez:
```
hey -n 200 -c 1 -q 200 -H "API_KEY: abc123" http://localhost:8080/

```

Você verá algo assim:
```
Status code distribution:
  [200]	100 responses
  [429]	100 responses
```