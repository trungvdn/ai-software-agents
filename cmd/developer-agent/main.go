package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/trungvdn/ai-software-agents/agents/developer"
	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/internal/config"
	"github.com/trungvdn/ai-software-agents/internal/database"
	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/tools"
	"github.com/trungvdn/ai-software-agents/storage/repositories"
)

func main() {
	// Load configuration
	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	if cfg.OllamaBaseURL == "" {
		log.Fatal("OLLAMA_BASE_URL environment variable is required")
	}
	if cfg.OllamaModel == "" {
		log.Fatal("OLLAMA_MODEL environment variable is required")
	}

	// Connect to database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository
	reflectionRepo := repositories.NewReflectionRepository(db)

	historicalBugRepo := repositories.NewHistoricalBugRepository(db)

	// Create embedder
	embedder := embedding.NewOllamaEmbedder(cfg.OllamaBaseURL, cfg.OllamaModel)

	// Reflection retrievers
	reflectionRetriever := reflection.NewReflectionRetriever(
		reflectionRepo,
		embedder,
	)

	// Historical bug retrievers
	historicalBugRetriever := historicalbug.NewHistoricalBugRetriever(
		historicalBugRepo,
		embedder,
	)

	knowledgeRetriever := developer.NewDefaultKnowledgeRetriever(
		reflectionRetriever,
		historicalBugRetriever,
	)

	ollamaClient, err := llm.NewOllamaClient(llm.OllamaConfig{
		Endpoint: cfg.OllamaBaseURL,
		Model:    cfg.OllamaChatModel,
	})
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v", err)
	}

	testdataPath := filepath.Join("testdata")
	// Convert to absolute path to handle different working directories
	absTestdataPath, err := filepath.Abs(testdataPath)
	if err != nil {
		log.Fatalf("Failed to resolve testdata path: %v", err)
	}
	log.Printf("Using testdata path: %s", absTestdataPath)

	codeRetriever := developer.NewDefaultCodeRetriever(
		tools.NewSearchSymbolTool(absTestdataPath),
		tools.NewReadFileTool(absTestdataPath),
	)

	promptBuilder := developer.NewDefaultPromptBuilder()

	diffGenerator := developer.NewDefaultPatchGenerator()

	// Developer Agent
	developerAgent := developer.NewDeveloperAgent(
		knowledgeRetriever,
		codeRetriever,
		promptBuilder,
		diffGenerator,
		ollamaClient,
	)

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	response, err := developerAgent.Execute(ctx, "Fix nil pointer in UserService")
	if err != nil {
		log.Fatalf("Error fixing bug: %v", err)
	}

	log.Printf("Developer Agent Response: %s", response.Analysis.SuggestedFix)
}
