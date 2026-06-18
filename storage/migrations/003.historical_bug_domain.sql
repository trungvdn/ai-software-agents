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

//docker cp storage/migrations/001_init.sql ai_agents_db:/tmp/init.sq
//docker exec -it ai_agents_db psql -U postgres -d ai_agents -f /tmp/003.sql

