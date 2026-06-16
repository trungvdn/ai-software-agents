#!/bin/bash
# Script to pull embedding models into Ollama
# This ensures models are available immediately after container startup

set -e

echo "Waiting for Ollama to be ready..."
for i in {1..30}; do
    if curl -s http://localhost:11434/api/tags > /dev/null; then
        echo "Ollama is ready!"
        break
    fi
    echo "Attempt $i/30: Ollama not ready yet, waiting..."
    sleep 2
done

echo "Pulling embedding models..."

# Pull nomic-embed-text (768 dims) - Primary model
echo "Pulling nomic-embed-text..."
ollama pull nomic-embed-text

# Pull all-minilm (384 dims) - Lightweight alternative
echo "Pulling all-minilm..."
ollama pull all-minilm

# Pull mistral-embed (1024 dims) - Higher quality
echo "Pulling mistral-embed..."
ollama pull mistral-embed

echo "All embedding models pulled successfully!"
echo ""
echo "Available models:"
curl -s http://localhost:11434/api/tags | jq '.models[].name'
