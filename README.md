# Isabela & Mateus 💝

Site single-page romântico para Isabela e Mateus, feito com Go e amor.

## Funcionalidades

- Contador de tempo juntos (atualizado em tempo real)
- Galeria de fotos com lightbox
- Declarações com animação ao scroll
- Design romântico minimalista, responsivo (otimizado para celular)
- Segurança reforçada (CSP, HSTS, anti-XSS, anti-clickjacking)

## Requisitos

- Go 1.22+
- Fly.io CLI (`flyctl`) para deploy

## Rodar localmente

```bash
go run main.go
```

Acessar **http://localhost:8080**

## Como customizar

### Nomes e data de início

Edite a struct `Dados` em `main.go`:

```go
Nome1:      "Isabela",
Nome2:      "Mateus",
DataInicio: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
```

### Declarações

Edite o slice `Declaracoes` em `main.go`:

```go
Declaracoes: []Declaracao{
    {
        Titulo: "Seu Título",
        Texto:  "Seu texto aqui...",
    },
},
```

### Fotos

1. Coloque suas imagens (JPG/PNG) em `static/img/`
2. Atualize o slice `Fotos` em `main.go`:

```go
Fotos: []Foto{
    {Src: "/static/img/sua-foto.jpg", Alt: "Descrição"},
},
```

## Deploy no Fly.io

### 1. Instalar o Fly.io CLI

```bash
curl -fsSL https://fly.io/install.sh | sh
```

### 2. Fazer login

```bash
flyctl auth login
```

### 3. Criar o app

```bash
flyctl apps create sitename-isabela-mateus
```

### 4. Fazer deploy

```bash
flyctl deploy
```

### 5. Acessar

Sua URL será: https://sitename-isabela-mateus.fly.dev

## Deploy alternativo no Render

Também pode hospedar gratuitamente no [Render](https://render.com):

1. Conecte o repositório GitHub
2. Crie um Web Service apontando para este repositório
3. Build command: `go build -o server main.go`
4. Start command: `./server`

## Segurança

- Content Security Policy (CSP) rigorosa
- Strict-Transport-Security (HSTS) com HTTPS forçado
- X-Frame-Options: DENY (anti-clickjacking)
- X-Content-Type-Options: nosniff
- Referrer-Policy: no-referrer
- html/template escapa automaticamente todo conteúdo (anti-XSS)
- Navegação de diretórios bloqueada
- Sem cookies, sem sessões, sem rastreamento