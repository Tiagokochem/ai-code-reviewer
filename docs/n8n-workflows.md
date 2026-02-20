# n8n Workflows - AI Code Reviewer

## Configuração Inicial

1. Acesse o n8n: http://localhost:5678
2. Login: admin / admin (configurável no .env)
3. Crie um novo workflow

## Workflow: Salvar Code Review

### Trigger: Webhook

1. Adicione um nó **Webhook**
2. Configure:
   - **Method**: POST
   - **Path**: `/code-review`
   - **Response Mode**: Respond When Last Node Finishes
3. Copie a URL do webhook (ex: `http://localhost:5678/webhook/code-review`)

### Ação: Salvar no Banco

1. Adicione um nó **Postgres** (ou **Set** para mock)
2. Configure para salvar:
   - `review_id`
   - `score`
   - `suggestions` (JSON)
   - `issues` (JSON)
   - `summary`
   - `code`
   - `timestamp`

### Ação: Gerar Relatório (Futuro)

1. Adicione nó para gerar PDF
2. Ou enviar email com relatório
3. Ou criar issue no GitHub

## Exemplo de Payload Recebido

```json
{
  "review_id": "review-1234567890",
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
  "summary": "Código funcional com melhorias sugeridas",
  "code": "function example() { ... }",
  "timestamp": "2026-02-20T18:00:00Z"
}
```

## Workflows Futuros

- **GitHub PR Integration**: Receber webhook do GitHub e analisar PR
- **Relatório PDF**: Gerar PDF com análise completa
- **Notificações**: Enviar email/Slack com resultados
- **Histórico**: Dashboard com histórico de análises
