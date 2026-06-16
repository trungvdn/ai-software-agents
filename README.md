# AI Software Agents

## Summary
This project is designed to learn and master AI Agents by building a multi-agent system for autonomous software development. The system leverages AI agents with memory, RAG (Retrieval-Augmented Generation), context builders, and embeddings to automatically analyze, debug, test, and improve code.

**Goal**: Create an intelligent system that can independently develop, maintain, and evolve software applications with minimal human intervention.

**Future Application**: The learnings will be applied to build **InvestPilot**, an intelligent information investment system powered by coordinated AI agents to analyze market data, generate insights, and recommend investment strategies.

## Project Phases

- **Phase 1**: Build BugFix Agent (using Memory, RAG, Context Builder, Embedding)
- **Phase 2**: Build Coder Agent, Reviewer Agent, Tester Agent, Planner Agent
- **Phase 3**: Build BA Agent, Merge Agent, Judge Agent, DevOps Agent
- **Phase 4**: Collaborate all agents in coordinated workflow
- **Phase 5**: Distributed system architecture


## Current Progress: Phase 1 Development

### Sprint 1

Implemented:
- ReadFileTool
- SearchCodeTool
- BugFix Agent skeleton

### Sprint 2

- PostgreSQL + pgvector
- Reflection Store
- Historical Bug Store

### Sprint 3

- Retriever
- Hybrid Search (BM25 + Vector)

### Sprint 4

- Episode Memory
- Reflection Generator
- Memory Promotion

### Sprint 5

- OpenAI/Ollama Integration
- Full Agent Loop
