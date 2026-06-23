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

	requirementPromptBuilder := developer.NewRequirementPromptBuilder()
	executeFeature(ctx, ollamaClient, requirementPromptBuilder, codeRetriever, "Implement a new feature to allow users to reset their passwords")
	executeBug(ctx, developerAgent, "Fix nill pointer in UserService")
	executeTest(ctx, developerAgent, "Add unit tests for UserService")

	// log.Printf("Developer Agent Response: %s", response.Analysis.SuggestedFix)
}

func executeFeature(ctx context.Context,
	client llm.Client,
	requirementPromptBuilder *developer.RequirementPromptBuilder,
	codeRetriever developer.CodeRetriever,
	featureDescription string) {
	task := &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeFeature,
		Description: featureDescription,
	}
	requirementAnalyzer := developer.NewDefaultRequirementAnalyzer(
		client,
		requirementPromptBuilder,
	)
	requirementAnalysis, err := requirementAnalyzer.Analyze(ctx, task)
	if err != nil {
		log.Printf("Error analyzing requirement: %v", err)
		return
	}
	log.Printf("Requirement Analysis: %+v", requirementAnalysis)

	codeContext, err := codeRetriever.Retrieve(&developer.RetrievalQuery{
		Query: featureDescription,
	})
	log.Printf("Retrieved %d files for the feature description", len(codeContext.Files))
	if err != nil {
		log.Printf("Error retrieving code context: %v", err)
		return
	}
	// Prompt builder
	// Implement plan

}

func executeBug(ctx context.Context, agent *developer.DeveloperAgent, bugDescription string) {
	_, err := agent.Execute(ctx, &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeBugFix,
		Description: bugDescription,
	})
	if err != nil {
		log.Printf("Error executing feature: %v", err)
		return
	}

}

func executeTest(ctx context.Context, agent *developer.DeveloperAgent, testDescription string) {
	_, err := agent.Execute(ctx, &developer_domain.DevelopmentTask{
		Type:        developer_domain.TaskTypeTest,
		Description: testDescription,
	})

	if err != nil {
		log.Printf("Error executing feature: %v", err)
		return
	}
}
