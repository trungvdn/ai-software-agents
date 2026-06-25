package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/trungvdn/ai-software-agents/internal/agents/ba"
	"github.com/trungvdn/ai-software-agents/internal/config"
	"github.com/trungvdn/ai-software-agents/internal/database"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure/mcp"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_epic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_requirement"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_requirement_package"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_story"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/publish_requirement"
	llm_infa "github.com/trungvdn/ai-software-agents/internal/requirement/infrastructure/llm"
	"github.com/trungvdn/ai-software-agents/shared/llm"
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

	baAgent := ba.NewBAAgent(
		generate_requirement_package.NewGenerateRequirementPackageUseCase(
			generate_requirement.NewGenerateRequirementUseCase(
				llm_infa.NewOllamaRequirementGenerator(ollamaClient),
			),
			generate_epic.NewGenerateEpicUseCase(
				llm_infa.NewOllamaEpicGenerator(ollamaClient),
			),
			generate_story.NewGenerateStoryUseCase(
				llm_infa.NewOllamaStoryGenerator(ollamaClient)),
		),
		publish_requirement.NewPublishRequirementUseCase(
			infrastructure.NewConfluencePublisher(
				infrastructure.NewRequirementFormatter(),
				mcp.NewMCPConfluenceClient(),
			),
		),
	)
	response, err := baAgent.Execute(context.Background(), ba.Request{
		Idea: "I want to build an intelligent information investment system that can analyze market trends, predict stock prices, and provide personalized investment recommendations to users based on their risk tolerance and financial goals.",
	})

	log.Printf("Requirement Aggregate: %+v", response.RequirementAggregate)
}
