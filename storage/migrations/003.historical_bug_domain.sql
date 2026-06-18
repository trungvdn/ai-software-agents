CREATE TABLE IF NOT EXISTS historical_bugs (
    id UUID PRIMARY KEY,
    title TEXT,
    root_cause TEXT,
    fix_summary TEXT,
    severity TEXT,
    usage_count INT,
    embedding vector(768),
    created_at TIMESTAMP
);



