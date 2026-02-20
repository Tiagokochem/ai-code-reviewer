FROM node:20-alpine

WORKDIR /app

# Copiar arquivos de dependências
COPY package.json ./

# Instalar dependências (npm install funciona com ou sem package-lock.json)
RUN npm install

# Copiar código fonte
COPY . .

# Expor porta
EXPOSE 3000

# Comando padrão (pode ser sobrescrito no docker-compose)
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
