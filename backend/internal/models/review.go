package models

// CodeReviewRequest representa a requisição de análise de código
type CodeReviewRequest struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
	Context  string `json:"context,omitempty"` // Contexto adicional (ex: PR description)
}

// CodeReviewResponse representa a resposta da análise
type CodeReviewResponse struct {
	Score       int      `json:"score"`       // Score de 0-100
	Suggestions []string `json:"suggestions"` // Lista de sugestões
	Issues      []Issue  `json:"issues"`      // Problemas encontrados
	Summary     string   `json:"summary"`     // Resumo da análise
	ReviewID    string   `json:"review_id"`   // ID único da análise
}

// Issue representa um problema encontrado no código
type Issue struct {
	Type        string `json:"type"`        // error, warning, info
	Severity    string `json:"severity"`    // high, medium, low
	Line        int    `json:"line"`        // Linha do problema
	Message     string `json:"message"`     // Mensagem descritiva
	Suggestion  string `json:"suggestion"`   // Sugestão de correção
}
