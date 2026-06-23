# Responsibility Developer Agent

- Implement Feature

- Retrieve Memory

- Analyze Bug

- Create Fix Strategy

- Search Symbol

- Read File

- Grep

- Build Code Context

- Generate Patch

- Generate Reflection

# Architecture

Developer Agent

├── Knowledge Retriever ✅
│
├── Code Retriever ✅
│
├── Prompt Builder ✅
│
├── LLM Reasoning ✅
│
├── Analysis Parser ✅
│
├── Patch Generator ❌
│
└── Reflection Writer ❌

# Week 1
1. Thêm Code Retriever
2. Thêm Tool use
3. Chuyển BugFix thành Developer Agent
4. Tạo Patch Generator
5. Tạo Code Patch

# Week 2
1. Refactor Execute() thành DevelopmentTask
2. Thêm TaskType
3. Thêm RequirementAnalyzer
4. Cho CodeRetriever nhận TaskContext
5. Hỗ trợ Feature Workflow
6. Test Generation