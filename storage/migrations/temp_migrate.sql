DROP TABLE IF EXISTS reflections;

CREATE TABLE reflections
(
    id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    importance_score DOUBLE PRECISION DEFAULT 0,
    usage_count INTEGER DEFAULT 0,
    last_accessed TIMESTAMP NULL,
    embedding VECTOR(768),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX ON reflections USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);
