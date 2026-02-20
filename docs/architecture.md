# Arquitetura do AI Code Reviewer

## Visão Geral

O sistema é composto por 4 serviços principais:

1. **Frontend React** - Interface do usuário
2. **Backend Go** - API REST para processamento
3. **n8n** - Automação e workflows
4. **PostgreSQL** - Banco de dados do n8n

## Fluxo de Dados

```
┌─────────────┐
│   Usuário   │
└──────┬──────┘
       │ 1. Cola código
       ▼
┌─────────────┐
│   React     │ POST /api/v1/review
│  Frontend   │─────────────────────┐
└─────────────┘                      │
                                     ▼
                            ┌──────────────┐
                            │   Go API     │
                            │   Backend    │
                            └──────┬───────┘
                                   │
                    ┌──────────────┼──────────────┐
                    │              │              │
                    ▼              ▼              ▼
            ┌──────────┐   ┌──────────┐   ┌──────────┐
            │   OpenAI │   │   n8n    │   │ Response │
            │    API   │   │ Webhook  │   │   JSON   │
            └──────────┘   └────┬─────┘   └────┬─────┘
                                │              │
                                │              │
                                ▼              │
                         ┌──────────────┐     │
                         │  PostgreSQL  │     │
                         │   (n8n DB)   │     │
                         └──────────────┘     │
                                              │
                                              ▼
                                    ┌─────────────┐
                                    │   Usuário   │
                                    │  (Frontend) │
                                    └─────────────┘
```

## Componentes Detalhados

### Frontend (React + TypeScript + Vite)

- **Porta**: 3000
- **Tecnologias**: React 18, TypeScript, Vite, Axios
- **Responsabilidades**:
  - Interface para colar código
  - Seleção de linguagem
  - Exibição de resultados (score, sugestões, issues)
  - Comunicação com backend via REST API

### Backend (Go + Gin)

- **Porta**: 8080
- **Tecnologias**: Go 1.21+, Gin framework
- **Estrutura**:
  ```
  backend/
  ├── cmd/api/          # Entry point
  ├── internal/
  │   ├── handlers/     # HTTP handlers
  │   ├── services/     # Business logic
  │   ├── models/       # Data models
  │   ├── config/       # Configuration
  │   └── api/          # Server setup
  └── Dockerfile
  ```
- **Endpoints**:
  - `POST /api/v1/review` - Analisa código
  - `GET /api/v1/health` - Health check

### n8n Workflow

- **Porta**: 5678
- **Responsabilidades**:
  - Receber webhook do backend com análise
  - Salvar análise no PostgreSQL
  - Gerar relatórios (futuro)
  - Notificações (futuro)

### PostgreSQL

- **Porta**: 5432
- **Uso**: Banco de dados do n8n para persistir workflows e dados

## Integração com IA

### OpenAI API (Padrão)

- **Modelo**: gpt-3.5-turbo (configurável)
- **Fallback**: Se não houver API key, usa análise mockada
- **Prompt**: Especializado para code review com retorno estruturado em JSON

### Alternativas Futuras

- Hugging Face API
- Ollama (local)
- Claude API

## Segurança

- CORS configurado no backend
- Variáveis de ambiente para secrets
- n8n com autenticação básica
- Validação de entrada no backend

## Escalabilidade

- Cada serviço em container separado
- Stateless backend (fácil escalar horizontalmente)
- Banco de dados pode ser migrado para serviço gerenciado
- Frontend pode usar CDN
