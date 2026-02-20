.PHONY: help up down build logs clean test

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

up: ## Sobe todos os serviços
	docker compose up -d

down: ## Para todos os serviços
	docker compose down

build: ## Rebuild todas as imagens
	docker compose build --no-cache

logs: ## Mostra logs de todos os serviços
	docker compose logs -f

logs-backend: ## Mostra logs do backend
	docker compose logs -f backend

logs-frontend: ## Mostra logs do frontend
	docker compose logs -f frontend

logs-n8n: ## Mostra logs do n8n
	docker compose logs -f n8n

clean: ## Remove containers, volumes e imagens
	docker compose down -v --rmi all

test-backend: ## Roda testes do backend
	cd backend && go test ./...

test-frontend: ## Roda testes do frontend
	cd frontend && npm test

restart: ## Reinicia todos os serviços
	docker compose restart

status: ## Mostra status dos serviços
	docker compose ps

shell-backend: ## Abre shell no container do backend
	docker compose exec backend sh

shell-frontend: ## Abre shell no container do frontend
	docker compose exec frontend sh

install-backend-deps: ## Instala dependências do backend localmente
	cd backend && go mod download

install-frontend-deps: ## Instala dependências do frontend localmente
	cd frontend && npm install
