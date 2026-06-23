package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/trungvdn/ai-software-agents/agents/developer"
	developer_domain "github.com/trungvdn/ai-software-agents/domain/developer"
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

	diffGenerator := developer.NewDefaultDiffGenerator()

	patchApplier := developer.NewDefaultPatchApplier(
		absTestdataPath,
	)

	// Developer Agent
	developerAgent := developer.NewDeveloperAgent(
		knowledgeRetriever,
		codeRetriever,
		promptBuilder,
		diffGenerator,
		patchApplier,
		ollamaClient,
	)

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	developerAgent.ExecuteFeature(ctx, &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeFeature,
		Description: "Implement a new feature that allows users to reset their password via email.",
	})
	developerAgent.ExecuteBugFix(ctx, &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeBugFix,
		Description: "Fix the bug where the application crashes when the user inputs an empty string.",
	})
	developerAgent.ExecuteTest(ctx, &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeTest,
		Description: "Write a test case to verify that the password reset feature works correctly.",
	})
}
