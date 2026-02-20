package main

import (
	"log"
	"os"

	"github.com/test-go/ai-code-reviewer/internal/api"
	"github.com/test-go/ai-code-reviewer/internal/config"
)

func main() {
	// Carregar configuração
	cfg := config.Load()

	// Criar servidor API
	server := api.NewServer(cfg)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor iniciando na porta %s", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
	}
}
