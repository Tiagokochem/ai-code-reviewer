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

type N8NService struct {
	config     *config.Config
	httpClient *http.Client
}

func NewN8NService(cfg *config.Config) *N8NService {
	return &N8NService{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SaveReview envia análise para n8n via webhook
func (s *N8NService) SaveReview(review *models.CodeReviewResponse, code string) error {
	payload := map[string]interface{}{
		"review_id":  review.ReviewID,
		"score":      review.Score,
		"suggestions": review.Suggestions,
		"issues":     review.Issues,
		"summary":    review.Summary,
		"code":       code,
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao serializar payload: %w", err)
	}

	req, err := http.NewRequest("POST", s.config.N8NWebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		// Não falhar se n8n não estiver disponível (modo desenvolvimento)
		return nil // Log error mas não retorna erro
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("erro ao salvar no n8n (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
