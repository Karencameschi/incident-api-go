# Incident API (Go)

API REST em Go para gerenciamento de incidentes de TI, desenvolvida como projeto de estudo.

## OBJETIVO

API desenvolvida em Go para praticar a linguagem, aplicando minha experiência profissional com o tratamento de incidentes de TI e reforçando meu portfólio de desenvolvimento de software.

## O QUE A API FAZ

- `GET /health`: verifica se a API está no ar
- `POST /incidents`: cadastra um incidente (o título é obrigatório)
- `GET /incidents`: lista os incidentes cadastrados
- Logs estruturados em JSON a cada requisição (método, rota, status e tempo de resposta)

Os dados ficam guardados **em memória**: ao reiniciar a API, a lista é zerada.

## TECNOLOGIAS UTILIZADAS

- Go (biblioteca padrão: `net/http`, `encoding/json`, `log/slog`)

## COMO RODAR

Pré-requisito: Go 1.22 ou superior (testado com Go 1.27).

```
git clone https://github.com/Karencameschi/incident-api-go.git
cd incident-api-go
go run .
```

A API sobe em `http://localhost:8080`.

## EXEMPLOS DE USO

Cadastrar um incidente (PowerShell):

```
Invoke-RestMethod -Method Post -Uri http://localhost:8080/incidents -ContentType "application/json" -Body '{"title":"Falha na rede","description":"Queda na regiao sul"}'
```

Listar incidentes:

```
curl http://localhost:8080/incidents
```

Exemplo de resposta:

```json
[
  {
    "id": 1,
    "title": "Falha na rede",
    "description": "Queda na regiao sul",
    "status": "open",
    "createdAt": "2026-11-11T21:03:05-03:00"
  }
]
```

## ESTRUTURA DO PROJETO

- `main.go`: inicia o servidor e define as rotas
- `handlers.go`: funções que respondem a cada rota
- `incident.go`: modelo `Incident` e armazenamento em memória (`Store`)
- `middleware.go`: registro de logs de cada requisição

## PRÓXIMOS PASSOS

- [ ] Buscar incidente por ID
- [ ] Atualizar o status de um incidente
- [ ] Testes automatizados
- [ ] Persistência em banco de dados (PostgreSQL)
