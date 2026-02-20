package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test-go/ai-code-reviewer/internal/models"
	"github.com/test-go/ai-code-reviewer/internal/services"
)

type CodeReviewHandler struct {
	reviewService *services.ReviewService
}

func NewCodeReviewHandler(reviewService *services.ReviewService) *CodeReviewHandler {
	return &CodeReviewHandler{
		reviewService: reviewService,
	}
}

// ReviewCode @Summary Analisa código usando IA
// @Description Recebe código e retorna análise com score, sugestões e issues
// @Tags code-review
// @Accept json
// @Produce json
// @Param request body models.CodeReviewRequest true "Código para análise"
// @Success 200 {object} models.CodeReviewResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/review [post]
func (h *CodeReviewHandler) ReviewCode(c *gin.Context) {
	var req models.CodeReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Validar linguagem suportada
	supportedLanguages := map[string]bool{
		"go": true, "javascript": true, "typescript": true, "python": true,
		"java": true, "rust": true, "cpp": true, "c": true,
		"php": true, "laravel": true, "vue": true,
	}

	if !supportedLanguages[req.Language] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Linguagem não suportada"})
		return
	}

	review, err := h.reviewService.ReviewCode(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao analisar código: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// HealthCheck @Summary Health check endpoint
// @Description Verifica se o serviço está funcionando
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/health [get]
func (h *CodeReviewHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "ai-code-reviewer",
		"version": "1.0.0",
	})
}
