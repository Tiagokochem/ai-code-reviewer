# AI Code Reviewer
![1](https://github.com/user-attachments/assets/d9b2d23f-d7d6-467d-a533-d3d5d200aa48)
![2](https://github.com/user-attachments/assets/51e239c3-f7a4-4d74-bb9b-8e9149e43197)



Code review system powered by AI, built with React, Go and n8n. It provides automated reviews with a quality score, improvement suggestions and optional report generation via n8n workflows.

## Overview

**AI Code Reviewer** is a full-stack application that lets you:

- Paste code in a React UI
- Get automated analysis via AI (OpenAI, with mock fallback when no API key is set)
- See a quality score and detailed suggestions
- Send review results to n8n for storage and further automation
- Run everything locally with Docker

### Supported languages

The system supports 11 languages:

- JavaScript / TypeScript
- Vue.js (with framework-specific guidelines)
- PHP / Laravel (Laravel and PSR conventions)
- Go
- Python
- Java
- Rust
- C++ / C

Each language is analyzed with its own conventions and best practices in mind.

## Architecture

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐
│   React     │──────│   Go API     │──────│     n8n     │
│  Frontend   │      │   Backend    │      │  Workflows  │
└─────────────┘      └──────────────┘      └─────────────┘
                            │
                     ┌──────────────┐
                     │  PostgreSQL  │
                     │   (n8n DB)   │
                     └──────────────┘
```

### Tech stack

- **Frontend**: React 18, TypeScript, Vite
- **Backend**: Go 1.21+ (Gin)
- **Workflows**: n8n (webhooks and automation)
- **AI**: OpenAI API (optional; runs with mock data if no key is provided)
- **Database**: PostgreSQL (used by n8n)
- **Containers**: Docker and Docker Compose

## Quick start

For a step-by-step guide, see [QUICKSTART.md](./QUICKSTART.md).

### Prerequisites

- Docker and Docker Compose
- OpenAI API key is optional for local runs (mock analysis is used when no key is set)

### Run the project

```bash
# 1. Optional: set environment variables
# Create .env and add OPENAI_API_KEY if you want real AI analysis

# 2. Start all services
make up
# or: docker compose up -d

# 3. Open in the browser:
# - Frontend: http://localhost:3000
# - Backend API: http://localhost:8080
# - n8n: http://localhost:5678
# - Health: http://localhost:8080/api/v1/health
```

### Useful commands

```bash
make help      # List available commands
make logs      # Follow logs for all services
make down      # Stop all services
make build     # Rebuild images
```

## Project structure

```
test-GO/
├── frontend/              # React + TypeScript + Vite
│   ├── src/
│   │   ├── components/
│   │   ├── services/
│   │   └── types/
│   └── Dockerfile
│
├── backend/               # Go API
│   ├── cmd/api/           # Entry point
│   ├── internal/
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── models/
│   │   └── config/
│   └── Dockerfile
│
├── n8n/
│   └── workflows/
│
├── docker-compose.yml
├── .env.example
└── README.md
```

## Configuration

### Environment variables

See `.env.example` for all options. Main ones:

- `OPENAI_API_KEY`: OpenAI API key (optional; mock is used when empty)
- `N8N_BASIC_AUTH_USER`: n8n login (default: admin)
- `N8N_BASIC_AUTH_PASSWORD`: n8n password (default: admin)
- `POSTGRES_USER` / `POSTGRES_PASSWORD`: PostgreSQL credentials

## Documentation

- [Architecture](./docs/architecture.md)
- [Development guide](./docs/development.md)
- [API](./docs/api.md)
- [n8n workflows](./docs/n8n-workflows.md)

## Tests

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && npm test
```

## Roadmap

- [x] Base project structure
- [x] OpenAI integration with mock fallback
- [x] React code review UI
- [x] n8n workflow for saving reviews
- [ ] PDF report generation
- [ ] GitHub PR integration
- [ ] Review history
- [ ] Metrics dashboard
- [ ] Automated tests (unit and integration)
- [ ] CI/CD

## Contributing

This is a portfolio project. Feel free to fork and adapt.

## License

MIT
