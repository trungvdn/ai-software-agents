package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/internal/config"
	"github.com/trungvdn/ai-software-agents/internal/database"
	ai_context "github.com/trungvdn/ai-software-agents/shared/context"
	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
	"github.com/trungvdn/ai-software-agents/storage/repositories"
)

type SampleReflection struct {
	Content string `json:"content"`
}

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
	// reflectionRepo := repositories.NewReflectionRepository(db)

	// Create embedder
	embedder := embedding.NewOllamaEmbedder(cfg.OllamaBaseURL, cfg.OllamaModel)
	if embedder == nil {
		log.Fatal("Failed to initialize embedder: Ollama configuration is required")
	}

	// embedAndSaveReflection(context.Background(), reflectionRepo, embedder)
	// retrieveReflection(context.Background(), reflectionRepo, embedder)

	// embedAndSaveHistoricalBug(context.Background(), repositories.NewHistoricalBugRepository(db), embedder)

	retrieveHistoricalBug(context.Background(), repositories.NewHistoricalBugRepository(db), embedder)

}

func embedAndSaveReflection(ctx context.Context, repo reflection.ReflectionRepository, embedder embedding.Embedder) {
	// Load sample reflections from JSON file
	data, err := os.ReadFile("cmd/seeder/sample_reflection.json")
	if err != nil {
		log.Fatalf("Failed to read sample_reflection.json: %v", err)
	}

	var samples []SampleReflection
	if err := json.Unmarshal(data, &samples); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Insert each reflection
	for _, sample := range samples {
		// Step 1: Create reflection
		refl := reflection.Reflection{
			ID:              uuid.New(),
			Content:         sample.Content,
			ImportanceScore: 0.5,
			UsageCount:      0,
			LastAccessed:    nil,
			CreatedAt:       time.Now(),
		}

		// Step 2: Generate embedding
		embeddingVector, err := embedder.Embed(ctx, sample.Content)
		if err != nil {
			log.Fatalf("Failed to generate embedding: %v", err)
		}
		refl.Embedding = embeddingVector

		// Step 3: Save reflection with embedding
		if err := repo.Save(ctx, refl); err != nil {
			log.Fatalf("Failed to save reflection: %v", err)
		}
		log.Printf("Inserted reflection: %s (embedding size: %d)", refl.Content, len(embeddingVector))
	}

	log.Printf("Successfully inserted %d reflections", len(samples))
}

func retrieveReflection(ctx context.Context, repo reflection.ReflectionRepository, embedder embedding.Embedder) {

	retriever := reflection.NewReflectionRetriever(
		repo,
		embedder,
	)

	reRanker := &retrieval.SimpleReRanker{}

	// Test with multiple search queries
	testQueries := []string{
		"Fix nil pointer in UserService",
		// "validate objects",
		// "database transaction",
		// "handle rollback",
	}

	for _, query := range testQueries {
		log.Printf("\nSearching for: '%s'", query)
		results, err := retriever.Retrieve(ctx, query, 3)
		if err != nil {
			log.Fatalf("Failed to retrieve similar reflections: %v", err)
		}
		// Rerank results based on similarity score
		if len(results) == 0 {
			log.Printf("  No results found")
		} else {
			reRanker.ReRank(ctx, query, results)
			for i, refl := range results {
				log.Printf(
					"[%d] score=%.4f source=%s content=%s",
					i+1,
					refl.Score,
					refl.Source,
					refl.Content,
				)
				log.Printf("  Updated usage count for reflection ID: %s", refl.ID)
				log.Printf("  Importance Score: %.2f, Usage Count: %d", refl.Metadata.ImportanceScore, refl.Metadata.UsageCount)

				log.Printf("  ---")
				// Optionally, you can also update the importance score based on some logic here
			}
		}

		// Inject LLM Context (for demonstration, we just print the retrieved reflections)

		// Context builder
		contextBuilder := ai_context.NewReflectionContextBuilder()
		ctxData, _ := contextBuilder.Build(
			ctx,
			results,
		)

		log.Println(ctxData.Content)

		// Increase usage count for the retrieved reflections
		var reflectionIDs []string
		for _, refl := range results {
			reflectionIDs = append(reflectionIDs, refl.ID)
		}

		repo.IncrementUsageCount(ctx, reflectionIDs)
	}
}

func embedAndSaveHistoricalBug(ctx context.Context, repo historicalbug.HistoricalBugRepository, embedder embedding.Embedder) {
	// Example historical bug data
	bugs := []*historicalbug.HistoricalBug{
		{
			ID:         uuid.New(),
			Title:      "Null Pointer Exception in UserService",
			RootCause:  "User object was not properly initialized before accessing its methods.",
			FixSummary: "Added null checks and proper initialization for User object in UserService.",
			Severity:   "High",
			UsageCount: 0,
			CreatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Title:      "Database transaction not rolled back",
			RootCause:  "In case of an error during the transaction, the rollback was not triggered.",
			FixSummary: "Added error handling to ensure that transactions are rolled back on failure.",
			Severity:   "Medium",
			UsageCount: 0,
			CreatedAt:  time.Now(),
		},
		{
			ID:         uuid.New(),
			Title:      "Panic due to nil context",
			RootCause:  "A function expected a non-nil context but received nil, leading to a panic when trying to access context values.",
			FixSummary: "Added checks to ensure that context is not nil before accessing its values.",
			Severity:   "Critical",
			UsageCount: 0,
			CreatedAt:  time.Now(),
		},
	}

	for _, bug := range bugs {
		embeddingVector, err := embedder.Embed(ctx, bug.RootCause+" "+bug.FixSummary)
		if err != nil {
			log.Fatalf("Failed to generate embedding: %v", err)
		}
		bug.Embedding = embeddingVector
	}

	// Save the historical bug with embedding
	if err := repo.Save(ctx, bugs); err != nil {
		log.Fatalf("Failed to save historical bugs: %v", err)
	}
	log.Printf("Successfully inserted %d historical bugs", len(bugs))
}

func retrieveHistoricalBug(ctx context.Context, repo historicalbug.HistoricalBugRepository, embedder embedding.Embedder) {
	// Example query to search for similar historical bugs
	query := "How to fix null pointer exception in user service?"
	embeddingVector, err := embedder.Embed(ctx, query)
	if err != nil {
		log.Fatalf("Failed to generate embedding for query: %v", err)
	}
	fmt.Printf("Embedding vector for query: %v\n", embeddingVector)
	bugs, err := repo.SearchSimilar(ctx, embeddingVector, 3)
	if err != nil {
		log.Fatalf("Failed to search for similar historical bugs: %v", err)
	}
	log.Printf("Search results for query: '%s'", query)
	for i, bug := range bugs {
		log.Printf(
			"[%d] Title: %s, Severity: %s, Root Cause: %s, Fix Summary: %s",
			i+1,
			bug.Title,
			bug.Severity,
			bug.RootCause,
			bug.FixSummary,
		)
	}
}
