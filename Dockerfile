# Imagem oficial Node.js 20 (requisito aios-core: Node >= 18, npm >= 9)
FROM node:20-slim

# Evita perguntas interativas e logs excessivos
ENV DEBIAN_FRONTEND=noninteractive

# Diretório de trabalho no container
WORKDIR /app

# Garante npm/npx atualizados
RUN npm install -g npm@latest

# Mantém o container rodando para você usar o shell
# Uso: docker compose run --rm aios-core bash
# Ou: docker compose run --rm aios-core npx aios-core init meu-projeto
CMD ["bash"]
