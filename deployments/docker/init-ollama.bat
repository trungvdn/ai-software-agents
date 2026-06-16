@echo off
REM Script to pull embedding models into Ollama on Windows
REM This ensures models are available immediately after container startup

echo Waiting for Ollama to be ready...
setlocal enabledelayedexpansion
for /L %%i in (1,1,30) do (
    curl -s http://localhost:11434/api/tags >nul 2>&1
    if !errorlevel! equ 0 (
        echo Ollama is ready!
        goto :ready
    )
    echo Attempt %%i/30: Ollama not ready yet, waiting...
    timeout /t 2 /nobreak
)

:ready
echo Pulling embedding models...

echo Pulling nomic-embed-text...
docker exec ai_agents_ollama ollama pull nomic-embed-text

echo Pulling all-minilm...
docker exec ai_agents_ollama ollama pull all-minilm

echo Pulling mistral-embed...
docker exec ai_agents_ollama ollama pull mistral-embed

echo All embedding models pulled successfully!
echo.
echo Available models:
docker exec ai_agents_ollama ollama list
