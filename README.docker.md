# Container Docker para aios-core

Ambiente com Node.js 20 para instalar e usar o **aios-core** no projeto **test-go** sem precisar instalar Node/npm na máquina.

## Pré-requisito

- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/install/) instalados.

## Uso

### 1. Instalar o aios-core no projeto (test-go)

Primeiro passo: instalar o aios-core no projeto existente.

**Opção A – Entrar no container e rodar:**

```bash
docker compose run --rm aios-core bash
```

Dentro do container (o diretório `/app` é o projeto test-go):

```bash
npx aios-core install
```

**Opção B – Um comando direto:**

```bash
docker compose run --rm aios-core npx aios-core install
```

Os arquivos ficam em `/app` no container (volume com o diretório atual), então tudo que o install criar será salvo no seu disco.

### 2. Criar novo projeto com outro nome (opcional)

Se no futuro quiser criar um projeto do zero com o nome `test-go`:

```bash
docker compose run --rm aios-core npx aios-core init test-go
# ou versão específica
docker compose run --rm aios-core npx aios-core@latest init test-go
```

### 3. Reconstruir a imagem (após mudar o Dockerfile)

```bash
docker compose build --no-cache
```

## Resumo dos arquivos

- **Dockerfile** – imagem base Node.js 20 (slim).
- **docker-compose.yml** – serviço `aios-core` com volume `.:/app`.
- **.dockerignore** – reduz contexto de build.
