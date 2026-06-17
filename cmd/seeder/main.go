package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/internal/config"
	"github.com/trungvdn/ai-software-agents/internal/database"
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
	reflectionRepo := repositories.NewReflectionRepository(db)

	// Create embedder
	embedder := embedding.NewOllamaEmbedder(cfg.OllamaBaseURL, cfg.OllamaModel)
	if embedder == nil {
		log.Fatal("Failed to initialize embedder: Ollama configuration is required")
	}

	// embedAndSaveReflection(context.Background(), reflectionRepo, embedder)
	retrieveReflection(context.Background(), reflectionRepo, embedder)

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
		"validate objects",
		"database transaction",
		"handle rollback",
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

		// Increase usage count for the retrieved reflections
		var reflectionIDs []string
		for _, refl := range results {
			reflectionIDs = append(reflectionIDs, refl.ID)
		}
		retriever.repo.IncrementUsageCount(ctx, reflectionIDs)

	}
}
