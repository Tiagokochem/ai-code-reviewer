# 🚀 Quick Start - AI Code Reviewer

Guia rápido para começar a usar o projeto em 5 minutos.

## 1. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```bash
# OpenAI API (opcional - funciona sem para desenvolvimento com mock)
OPENAI_API_KEY=

# Portas (padrões funcionam)
BACKEND_PORT=8080
FRONTEND_PORT=3000
N8N_PORT=5678

# n8n
N8N_BASIC_AUTH_USER=admin
N8N_BASIC_AUTH_PASSWORD=admin

# PostgreSQL
POSTGRES_DB=n8n
POSTGRES_USER=n8n
POSTGRES_PASSWORD=n8n_password
```

**Nota**: Se não tiver `OPENAI_API_KEY`, o sistema usará análises mockadas para desenvolvimento.

## 2. Subir os Serviços

```bash
# Opção 1: Usando Makefile (recomendado)
make up

# Opção 2: Docker Compose direto
docker compose up -d
```

## 3. Acessar as Aplicações

Após alguns segundos, os serviços estarão disponíveis:

- **Frontend React**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **n8n Dashboard**: http://localhost:5678
- **Health Check**: http://localhost:8080/api/v1/health

## 4. Testar o Sistema

1. Abra http://localhost:3000 no navegador
2. Cole algum código (exemplo abaixo)
3. Selecione a linguagem
4. Clique em "Analisar Código"
5. Veja o resultado com score e sugestões

### Exemplo de Código para Testar

**JavaScript:**
```javascript
function calculateTotal(items) {
  let total = 0
  for (let i = 0; i < items.length; i++) {
    total += items[i].price
  }
  return total
}
```

**PHP/Laravel:**
```php
<?php

class OrderService
{
    public function calculateTotal($items)
    {
        $total = 0;
        foreach ($items as $item) {
            $total += $item->price;
        }
        return $total;
    }
}
```

**Vue.js:**
```vue
<template>
  <div>
    <p>Total: {{ total }}</p>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['items'])

const total = computed(() => {
  return props.items.reduce((sum, item) => sum + item.price, 0)
})
</script>
```

## 5. Configurar n8n (Opcional)

1. Acesse http://localhost:5678
2. Login: `admin` / `admin`
3. Crie um novo workflow
4. Adicione um nó **Webhook**:
   - Method: POST
   - Path: `/code-review`
5. Copie a URL do webhook
6. Atualize `N8N_WEBHOOK_URL` no `.env` do backend (ou use o padrão)

## Comandos Úteis

```bash
# Ver logs
make logs

# Parar serviços
make down

# Rebuild tudo
make build

# Ver ajuda completa
make help
```

## Troubleshooting

### Porta já em uso

Se alguma porta estiver ocupada, altere no `.env`:

```bash
BACKEND_PORT=8081
FRONTEND_PORT=3001
```

### Erro ao iniciar backend

```bash
# Ver logs específicos
make logs-backend

# Rebuild backend
docker compose build backend
docker compose up -d backend
```

### Frontend não conecta ao backend

Verifique se `VITE_API_URL` no frontend está correto (padrão: `http://localhost:8080`)

## Próximos Passos

- Leia [docs/architecture.md](./docs/architecture.md) para entender a arquitetura
- Veja [docs/development.md](./docs/development.md) para desenvolvimento local
- Configure workflows no n8n seguindo [docs/n8n-workflows.md](./docs/n8n-workflows.md)
