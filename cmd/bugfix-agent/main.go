package main

import (
	"context"
	"log"
	"time"

	"github.com/trungvdn/ai-software-agents/agents/bugfix"
	"github.com/trungvdn/ai-software-agents/agents/coder"
	"github.com/trungvdn/ai-software-agents/agents/planner"
	"github.com/trungvdn/ai-software-agents/domain/codebase"
	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/internal/config"
	"github.com/trungvdn/ai-software-agents/internal/database"
	prompt_context "github.com/trungvdn/ai-software-agents/shared/context"
	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
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

	codeBaseRepo := repositories.NewCodeBaseRepository(db)

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

	// Codebase retriever
	codeBaseRetriever := codebase.NewCodeBaseRetriever(
		codeBaseRepo,
		embedder,
	)

	mergeRetriever := retrieval.NewMergeRetriever(
		reflectionRetriever,
		historicalBugRetriever,
		codeBaseRetriever,
	)

	reRanker := &retrieval.SimpleReRanker{}

	knowledgeContextBuilder := prompt_context.NewKnowledgeContextBuilder()

	ollamaClient, err := llm.NewOllamaClient(llm.OllamaConfig{
		Endpoint: cfg.OllamaBaseURL,
		Model:    cfg.OllamaChatModel,
	})
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v", err)
	}

	// Planner
	planner := planner.NewLLMChangePlanner(ollamaClient)

	// Coder
	coder := coder.NewCoderAgent(
		ollamaClient,
	)

	// Fix bug agent
	fixBugAgent := bugfix.NewBugFixAgent(
		mergeRetriever,
		reRanker,
		knowledgeContextBuilder,
		ollamaClient,
		planner,
		coder,
	)

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	_, err = fixBugAgent.FixBug(ctx, "Fix nil pointer in UserService")
	if err != nil {
		log.Fatalf("Error fixing bug: %v", err)
	}
}
