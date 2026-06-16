-- Verify and optimize the vector index for pgvector similarity search

-- Check existing indexes
SELECT indexname, indexdef FROM pg_indexes WHERE tablename = 'reflections';

-- Drop old index if it exists and recreate with proper operators
DROP INDEX IF EXISTS reflections_embedding_idx;

-- Create index with cosine distance operator (L2 distance operator)
-- pgvector uses <-> for L2 distance, <#> for negative inner product, <=> for cosine distance
CREATE INDEX reflections_embedding_idx ON reflections 
USING ivfflat (embedding vector_cosine_ops) 
WITH (lists = 100);

-- Verify the index was created
SELECT indexname, indexdef FROM pg_indexes WHERE tablename = 'reflections';

-- Check the reflections data
SELECT id, content, embedding FROM reflections LIMIT 5;
