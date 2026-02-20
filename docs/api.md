# API Documentation

## Base URL

```
http://localhost:8080/api/v1
```

## Endpoints

### POST /review

Analisa código usando IA e retorna score, sugestões e issues.

**Request:**

```json
{
  "code": "function example() {\n  return 'hello';\n}",
  "language": "javascript",
  "context": "Optional: contexto adicional sobre o código"
}
```

**Response (200 OK):**

```json
{
  "score": 75,
  "suggestions": [
    "Adicione tratamento de erros",
    "Use nomes mais descritivos"
  ],
  "issues": [
    {
      "type": "warning",
      "severity": "medium",
      "line": 10,
      "message": "Função muito longa",
      "suggestion": "Dividir em funções menores"
    }
  ],
  "summary": "Código funcional com algumas oportunidades de melhoria",
  "review_id": "review-1234567890"
}
```

**Erros:**

- `400 Bad Request`: Dados inválidos ou linguagem não suportada
- `500 Internal Server Error`: Erro ao processar análise

**Linguagens Suportadas:**

- `go`
- `javascript`
- `typescript`
- `vue` (Vue.js)
- `php`
- `laravel` (Laravel/PHP)
- `python`
- `java`
- `rust`
- `cpp`
- `c`

### GET /health

Verifica se o serviço está funcionando.

**Response (200 OK):**

```json
{
  "status": "ok",
  "service": "ai-code-reviewer",
  "version": "1.0.0"
}
```

## Exemplo de Uso

### cURL

```bash
curl -X POST http://localhost:8080/api/v1/review \
  -H "Content-Type: application/json" \
  -d '{
    "code": "function add(a, b) { return a + b; }",
    "language": "javascript"
  }'
```

### JavaScript (Fetch)

```javascript
const response = await fetch('http://localhost:8080/api/v1/review', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    code: 'function add(a, b) { return a + b; }',
    language: 'javascript',
  }),
});

const review = await response.json();
console.log(review);
```

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func main() {
    reqBody := map[string]string{
        "code":     "func Add(a, b int) int { return a + b }",
        "language": "go",
    }
    
    jsonData, _ := json.Marshal(reqBody)
    resp, _ := http.Post(
        "http://localhost:8080/api/v1/review",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    defer resp.Body.Close()
}
```
