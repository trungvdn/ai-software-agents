# AI Software Agents

## Summary
This project is designed to learn and master AI Agents by building a multi-agent system for autonomous software development. The system leverages AI agents with memory, RAG (Retrieval-Augmented Generation), context builders, and embeddings to automatically analyze, debug, test, and improve code.

**Goal**: Create an intelligent system that can independently develop, maintain, and evolve software applications with minimal human intervention.

**Future Application**: The learnings will be applied to build **InvestPilot**, an intelligent information investment system powered by coordinated AI agents to analyze market data, generate insights, and recommend investment strategies.

# Final Vision

```
User:
"I want to build an intelligent information investment system"

            ↓

BA Agent
    analyze requirements

            ↓

Planner Agent
    divide tasks

            ↓

Scheduler Agent
    distribute tasks

            ↓

Coder Agents (N)
    implement

            ↓

Tester Agents (N)
    test

            ↓

Reviewer Agents (N)
    review

            ↓

Judge Agent
    resolve conflict

            ↓

DevOps Agent
    deploy

            ↓

Running Software
```

---

# PHASE 1 — Foundation

Objective:

```
Agent learns from experience
```

---

## Sprint 1

Basic Tool

```
ReadFileTool ✅
SearchCodeTool ✅
BugFix Agent skeleton ✅
```

---

## Sprint 2

Store data

```
PostgreSQL + pgvector ✅
Reflection Store ✅
Historical Bug Store
```

---

## Sprint 3

```
Docker + pgvector ✅
Embedding Provider ✅
Reflection Repository ✅
Semantic Search ✅
SearchResult ✅
Typed Metadata ✅
Reflection Retriever ✅
Simple ReRanker ✅
Usage Tracking ✅
Cosine Similarity Search ✅
Context Builder
```

---

Deliverable

```
Reflection Memory System
```

---

# PHASE 2 — Single Agent

Objective:

```
A complete Agent
```

---

## Sprint 4

BugFix Agent v1 

```
Bug 
↓
Reflection 
↓
LLM 
↓
Suggestion
```

---

## Sprint 5

Historical Bug Domain 

```
Case Study Memory
```

---

## Sprint 6

Code Change Planner

```
Bug
↓
Affected Files
↓
Change Plan
```

---

## Sprint 7

Coder Agent v1

```
Task
↓
Code Suggestion
```

---

Deliverable

```
Junior Coder Agent
```

---

# Phase 3 — Memory Lifecycle

Learning how memory lifecycle works.

---

# Phase 4 — Advanced Coder Agent

Coder Agent v2

Adds:

```
Code Retrieval
Historical Bug
Coding Standard
Architecture Decision
```

---

Pipeline:

```
Task
↓
Retrieve Knowledge
↓
Generate Code
↓
Generate Unit Test
```

---

Deliverable

```
Middle Coder Agent
```

---

# Phase 5 — Senior Coder Agent

Coder Agent v3

Adds:

```
Multi File Change
Dependency Analysis
Refactoring
```

---

Example:

```
Add RBAC
```

---

Agent knows:

```
User Service
Role Service
Middleware
Database Migration
Tests
```

All need to be modified.

---

Deliverable

```
Senior Coder Agent
```

---

# PHASE 3 — Learning System

Objective:

```
Agent learns

---

## Sprint 8

Memory Lifecycle

```
Episode
↓
Reflection
↓
Coding Standard
```

---

## Sprint 9

Coding Standard Domain

```
Rules
Policies
Best Practices
```

---

## Sprint 10

Memory Promotion Pipeline

```
Promote
Decay
Prune
```

---

Deliverable

```
Self-Learning Agent
```

---

# PHASE 4 — Engineering Agent

Objective:

```
Agent truly codes

---

## Sprint 11

Code Retrieval

```
Repository
Service
Controller
```

---

## Sprint 12

Architecture Decision Domain

```
ADR
Design Decision
```

---

## Sprint 13

Code Change Planner

```
Bug
↓
Impact Analysis
↓
Files To Change
```

---

Deliverable

```
Software Engineer Agent
```

---

# PHASE 5 — Team Simulation

Objective:

```
Multiple Agents working together

---

## Sprint 14

Planner Agent

```
Requirement
↓
Tasks
```

---

## Sprint 15

Coder Agent

```
Task
↓
Code
```

---

## Sprint 16

Tester Agent

```
Task
↓
Test Cases
```

---

## Sprint 17

Reviewer Agent

```
Code Review
```

---

## Sprint 18

Judge Agent

```
Conflict Resolution

Coder
Tester
Reviewer
```

---

Deliverable

```
AI Scrum Team
```

---

# PHASE 6 — Workflow Engine

Objective:

```
Agent Orchestration

---

## Sprint 19

Task Graph

```
DAG
Dependencies
```

---

## Sprint 20

Scheduler Agent

```
Assign Tasks
```

---

## Sprint 21

Workflow State Machine

```
Todo
Doing
Review
Done
```

---

Deliverable

```
Agent Orchestration Platform
```

---

# PHASE 7 — AI Software Company

Objective:

```
Build software from Requirements

---

## Sprint 22

BA Agent

```
User Idea
↓
Requirements
```

---

## Sprint 23

Project Manager Agent

```
Roadmap
Milestones
```

---

## Sprint 24

Multi-Agent Collaboration

```
BA
Planner
Scheduler
Coder
Tester
Reviewer
Judge
```

---

Deliverable

```
AI Software Company v1
```

---

# PHASE 8 — Production Scale

Objective:

```
Scale multiple projects

---

## Sprint 25

Shared Memory

```
Cross Project Learning
```

---

## Sprint 26

Knowledge Graph

```
Architecture
Code
Bug
Decision
```

---

## Sprint 27

Hybrid Search

```
BM25
Vector
Graph
```



## Sprint 27

Cross Encoder ReRanker



## Sprint 28

Human In The Loop



Deliverable

```
AI Software Company v2
```
