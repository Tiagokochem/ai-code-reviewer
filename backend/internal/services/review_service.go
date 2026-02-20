package services

import (
	"github.com/test-go/ai-code-reviewer/internal/models"
)

type ReviewService struct {
	aiService  *AIService
	n8nService *N8NService
}

func NewReviewService(aiService *AIService, n8nService *N8NService) *ReviewService {
	return &ReviewService{
		aiService:  aiService,
		n8nService: n8nService,
	}
}

// ReviewCode executa análise completa: IA + salvamento no n8n
func (s *ReviewService) ReviewCode(req *models.CodeReviewRequest) (*models.CodeReviewResponse, error) {
	// 1. Analisar código com IA
	review, err := s.aiService.ReviewCode(req)
	if err != nil {
		return nil, err
	}

	// 2. Salvar no n8n (assíncrono - não bloqueia resposta)
	go func() {
		if err := s.n8nService.SaveReview(review, req.Code); err != nil {
			// Log error (em produção usar logger)
			_ = err
		}
	}()

	return review, nil
}
