# 🔥 AI Code Reviewer

Sistema de análise de código com IA que integra React, Go e n8n para fornecer code reviews automatizados com score de qualidade, sugestões de melhoria e geração de relatórios.

## 🎯 Visão Geral

O **AI Code Reviewer** é uma aplicação full-stack que permite:

- 📝 **Colar código** via interface React
- 🤖 **Análise automática** por IA (OpenAI/Hugging Face)
- 📊 **Score de qualidade** e sugestões detalhadas
- 🔄 **Automação n8n** para salvar análises e gerar relatórios
- 🐳 **100% Dockerizado** para fácil execução local

### 🗣️ Linguagens Suportadas

O sistema suporta análise de código em **11 linguagens**:

- **JavaScript** / **TypeScript**
- **Vue.js** (com diretrizes específicas do framework)
- **PHP** / **Laravel** (com padrões Laravel e PSR)
- **Go**
- **Python**
- **Java**
- **Rust**
- **C++** / **C**

Cada linguagem recebe análise especializada considerando suas convenções, padrões e melhores práticas específicas.

## 🏗️ Arquitetura

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐
│   React     │──────│   Go API     │──────│     n8n     │
│  Frontend   │      │   Backend    │      │  Workflows  │
└─────────────┘      └──────────────┘      └─────────────┘
                            │
                            │
                     ┌──────────────┐
                     │  PostgreSQL  │
                     │   (n8n DB)   │
                     └──────────────┘
```

### Stack Tecnológica

- **Frontend**: React 18 + TypeScript + Vite
- **Backend**: Go 1.21+ (Gin framework)
- **Workflow**: n8n (automação e webhooks)
- **IA**: OpenAI API (ou Hugging Face como alternativa)
- **Banco de Dados**: PostgreSQL (via n8n)
- **Containerização**: Docker + Docker Compose

## 🚀 Início Rápido

**📖 Para um guia passo a passo completo, veja [QUICKSTART.md](./QUICKSTART.md)**

### Pré-requisitos

- Docker e Docker Compose instalados
- Chave de API da OpenAI (ou Hugging Face) - **opcional** para testes locais (usa mock sem API key)

### Executar o Projeto

```bash
# 1. Configure as variáveis de ambiente (opcional)
# Crie .env com OPENAI_API_KEY se tiver (senão usa mock)

# 2. Suba todos os serviços
make up
# ou: docker compose up -d

# 3. Acesse as aplicações:
# - Frontend React: http://localhost:3000
# - Backend Go API: http://localhost:8080
# - n8n Dashboard: http://localhost:5678
# - Health Check: http://localhost:8080/api/v1/health
```

### Comandos Úteis

```bash
make help      # Ver todos os comandos disponíveis
make logs      # Ver logs de todos os serviços
make down      # Parar todos os serviços
make build     # Rebuild todas as imagens
```

## 📁 Estrutura do Projeto

```
test-GO/
├── frontend/              # React + TypeScript + Vite
│   ├── src/
│   │   ├── components/   # Componentes React
│   │   ├── services/     # API clients
│   │   └── types/        # TypeScript types
│   └── Dockerfile
│
├── backend/              # Go API
│   ├── cmd/
│   │   └── api/         # Entry point
│   ├── internal/
│   │   ├── handlers/    # HTTP handlers
│   │   ├── services/    # Business logic
│   │   ├── models/      # Data models
│   │   └── ai/          # IA integration
│   └── Dockerfile
│
├── n8n/                 # n8n workflows
│   └── workflows/       # JSON workflows
│
├── docker-compose.yml    # Orquestração de serviços
├── .env.example         # Template de variáveis
└── README.md           # Este arquivo
```

## 🔧 Configuração

### Variáveis de Ambiente

Veja `.env.example` para todas as variáveis disponíveis:

- `OPENAI_API_KEY`: Chave da API OpenAI (opcional - pode usar mock)
- `N8N_BASIC_AUTH_USER`: Usuário do n8n (padrão: admin)
- `N8N_BASIC_AUTH_PASSWORD`: Senha do n8n (padrão: admin)
- `POSTGRES_USER`: Usuário do PostgreSQL
- `POSTGRES_PASSWORD`: Senha do PostgreSQL

## 📚 Documentação

- [Arquitetura Detalhada](./docs/architecture.md)
- [Guia de Desenvolvimento](./docs/development.md)
- [API Documentation](./docs/api.md)
- [n8n Workflows](./docs/n8n-workflows.md)

## 🧪 Testes

```bash
# Testes do backend Go
cd backend
go test ./...

# Testes do frontend React
cd frontend
npm test
```

## 📝 Roadmap

- [x] Estrutura base do projeto
- [x] Integração com OpenAI API (com fallback para mock)
- [x] Interface React para code review
- [x] Workflow n8n para salvar análises
- [ ] Geração de relatórios PDF
- [ ] Integração com GitHub PRs
- [ ] Histórico de análises
- [ ] Dashboard de métricas
- [ ] Testes automatizados (unit + integration)
- [ ] CI/CD pipeline

## 🤝 Contribuindo

Este é um projeto de portfólio. Sinta-se livre para fazer fork e melhorias!

## 📄 Licença

MIT License
