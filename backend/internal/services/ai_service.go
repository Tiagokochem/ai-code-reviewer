package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/test-go/ai-code-reviewer/internal/config"
	"github.com/test-go/ai-code-reviewer/internal/models"
)

type AIService struct {
	config     *config.Config
	httpClient *http.Client
}

func NewAIService(cfg *config.Config) *AIService {
	return &AIService{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ReviewCode analisa código usando IA
func (s *AIService) ReviewCode(req *models.CodeReviewRequest) (*models.CodeReviewResponse, error) {
	// Se não tiver API key, retorna análise mockada para desenvolvimento
	if s.config.OpenAIAPIKey == "" {
		return s.mockReview(req), nil
	}

	return s.reviewWithOpenAI(req)
}

func (s *AIService) reviewWithOpenAI(req *models.CodeReviewRequest) (*models.CodeReviewResponse, error) {
	// Construir prompt para análise de código
	prompt := s.buildPrompt(req)

	// Chamar OpenAI API
	payload := map[string]interface{}{
		"model": s.config.OpenAIModel,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Você é um code reviewer experiente. Analise o código fornecido e retorne um JSON com: score (0-100), suggestions (array de strings), issues (array com type, severity, line, message, suggestion), summary (string).",
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": 0.3,
		"max_tokens": 2000,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar payload: %w", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.config.OpenAIAPIKey)

	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar OpenAI API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro da API OpenAI (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse da resposta (simplificado - em produção usar estrutura completa)
	var openAIResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, fmt.Errorf("erro ao parsear resposta: %w", err)
	}

	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("resposta vazia da API")
	}

	// Parse do JSON retornado pela IA
	content := openAIResp.Choices[0].Message.Content
	var reviewResp models.CodeReviewResponse
	if err := json.Unmarshal([]byte(content), &reviewResp); err != nil {
		// Se não conseguir parsear, criar resposta básica
		reviewResp = models.CodeReviewResponse{
			Score:       75,
			Suggestions: []string{content},
			Summary:     content,
		}
	}

	reviewResp.ReviewID = generateReviewID()

	return &reviewResp, nil
}

func (s *AIService) buildPrompt(req *models.CodeReviewRequest) string {
	// Adicionar contexto específico para linguagens com frameworks
	languageContext := s.getLanguageContext(req.Language)
	
	prompt := fmt.Sprintf(`
Analise o seguinte código %s%s e forneça um code review detalhado:

%s

%s

%s

Retorne um JSON válido com:
- score: número de 0 a 100
- suggestions: array de strings com sugestões de melhoria
- issues: array de objetos com {type, severity, line, message, suggestion}
- summary: resumo textual da análise
`, req.Language, languageContext, req.Code, req.Context, s.getReviewGuidelines(req.Language))

	return prompt
}

func (s *AIService) getLanguageContext(language string) string {
	contexts := map[string]string{
		"laravel": " (framework Laravel/PHP - considere padrões Laravel, Eloquent, Blade, Service Providers, etc.)",
		"vue":     " (framework Vue.js - considere Composition API, Options API, Vuex/Pinia, diretivas Vue, etc.)",
		"php":     " (PHP - considere PSR standards, namespaces, type hints, etc.)",
	}
	if ctx, ok := contexts[language]; ok {
		return ctx
	}
	return ""
}

func (s *AIService) getReviewGuidelines(language string) string {
	guidelines := map[string]string{
		"laravel": `
Diretrizes específicas para Laravel:
- Verifique uso adequado de Eloquent ORM
- Valide seguimento de convenções Laravel (naming, estrutura de pastas)
- Analise uso de Service Providers, Middleware, Requests
- Verifique segurança (CSRF, SQL injection, XSS)
- Considere boas práticas de Blade templates`,
		"vue": `
Diretrizes específicas para Vue.js:
- Verifique uso correto de reatividade (ref, reactive, computed)
- Analise estrutura de componentes e props
- Verifique uso adequado de lifecycle hooks
- Considere performance (v-if vs v-show, lazy loading)
- Analise uso de diretivas e eventos`,
		"php": `
Diretrizes específicas para PHP:
- Verifique seguimento de PSR standards (PSR-1, PSR-12)
- Analise uso de type hints e return types
- Verifique tratamento de exceções
- Considere segurança (prepared statements, validação de input)
- Analise uso de namespaces e autoloading`,
	}
	if guide, ok := guidelines[language]; ok {
		return guide
	}
	return ""
}

func (s *AIService) mockReview(req *models.CodeReviewRequest) *models.CodeReviewResponse {
	// Análise mockada para desenvolvimento sem API key
	return &models.CodeReviewResponse{
		Score: 75,
		Suggestions: []string{
			"Considere adicionar tratamento de erros",
			"Variáveis poderiam ter nomes mais descritivos",
			"Adicione comentários para funções complexas",
		},
		Issues: []models.Issue{
			{
				Type:     "warning",
				Severity: "medium",
				Line:     10,
				Message:  "Função muito longa, considere dividir em funções menores",
				Suggestion: "Extrair lógica em funções auxiliares",
			},
		},
		Summary:     "Código funcional com algumas oportunidades de melhoria em legibilidade e manutenibilidade.",
		ReviewID:    generateReviewID(),
	}
}

func generateReviewID() string {
	return fmt.Sprintf("review-%d", time.Now().Unix())
}
