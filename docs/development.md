# Guia de Desenvolvimento

## Setup Local

### Pré-requisitos

- Docker e Docker Compose
- (Opcional) Go 1.21+ instalado localmente
- (Opcional) Node.js 20+ instalado localmente

### Executar com Docker

```bash
# Subir todos os serviços
docker compose up -d

# Ver logs
docker compose logs -f

# Parar serviços
docker compose down
```

### Desenvolvimento Local (sem Docker)

#### Backend Go

```bash
cd backend

# Instalar dependências
go mod download

# Rodar servidor
go run cmd/api/main.go

# Testes
go test ./...
```

#### Frontend React

```bash
cd frontend

# Instalar dependências
npm install

# Rodar dev server
npm run dev

# Build
npm run build
```

## Estrutura de Código

### Backend Go

Seguindo padrões Go comuns:

- `cmd/` - Entry points da aplicação
- `internal/` - Código privado da aplicação
- Handlers separados de Services
- Models para estruturas de dados
- Config centralizado

### Frontend React

Seguindo padrões React modernos:

- Componentes funcionais com hooks
- TypeScript para type safety
- Separação de concerns (components, services, types)
- CSS modules ou styled components (futuro)

## Variáveis de Ambiente

Copie `.env.example` para `.env` e configure:

```bash
# OpenAI (opcional - funciona sem para desenvolvimento)
OPENAI_API_KEY=sk-...

# Portas (padrões funcionam)
BACKEND_PORT=8080
FRONTEND_PORT=3000
N8N_PORT=5678
```

## Testes

### Backend

```bash
cd backend
go test ./... -v
```

### Frontend

```bash
cd frontend
npm test
```

## Debugging

### Backend

```bash
# Logs do container
docker compose logs -f backend

# Entrar no container
docker compose exec backend sh
```

### Frontend

```bash
# Logs do container
docker compose logs -f frontend

# Dev tools do navegador
# Abrir http://localhost:3000
```

### n8n

```bash
# Acessar dashboard
# http://localhost:5678

# Criar workflow manualmente ou importar JSON
```

## Comandos Úteis

```bash
# Rebuild de um serviço específico
docker compose build backend
docker compose up -d backend

# Ver status dos serviços
docker compose ps

# Limpar volumes (cuidado - apaga dados)
docker compose down -v
```

## Próximos Passos

1. Adicionar testes unitários
2. Configurar CI/CD
3. Adicionar logging estruturado
4. Implementar métricas (Prometheus)
5. Adicionar documentação Swagger
