
CREATE TABLE IF NOT EXISTS codebase
(
    id UUID PRIMARY KEY,
    file_path TEXT NOT NULL,
    content TEXT NOT NULL,
    embedding VECTOR(768),
    language TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);