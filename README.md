# AI Software Agents

## Summary
This project is designed to learn and master AI Agents by building a multi-agent system for autonomous software development. The system leverages AI agents with memory, RAG (Retrieval-Augmented Generation), context builders, and embeddings to automatically analyze, debug, test, and improve code.

**Goal**: Create an intelligent system that can independently develop, maintain, and evolve software applications with minimal human intervention.

**Future Application**: The learnings will be applied to build **InvestPilot**, an intelligent information investment system powered by coordinated AI agents to analyze market data, generate insights, and recommend investment strategies.

## Recent Infrastructure Update
| Item | Before Optimization | After Optimization | Improvement |
| --- | --- | --- | --- |
| **Inference Backend** | CPU | RTX 2060 CUDA | ✅ Switched to GPU |
| **Model Offloading** | 0 layer GPU | 24/29 layers on GPU | ✅ Most model layers on GPU |
| **VRAM Usage** | 0 MB | ~4039 MB | ✅ Utilizes GPU memory |
| **GPU Memory Headroom** | N/A | ~1065 MB | Stable headroom available |
| **Concurrent Requests** | 1 | 2 | **2×** |
| **Slots** | 1 (`slot0`) | 2 (`slot0`, `slot1`) | **2×** |
| **Total Context** | 4096 | 8192 | **2×** |
| **Prompt Cache** | Fresh start | 202 → 333 MB | Cache efficiency improving |
| **Prompt Similarity** | Insignificant | 0.43 – 0.63 | Better prompt reuse |
| **Graph Reuse** | Low | 2732 → 4317 | Continuous increase |
| **Prompt Evaluation** | ~4 token/s (cold) | 274–694 token/s (warm) | **≈60–150×** |
| **Generation Speed / Request** | ~25 token/s | ~21–22 token/s | Slight decrease due to parallelism |
| **Total Throughput** | ~25 token/s | ~44 token/s | **≈+76%** |
| **Queue** | None | None | No bottleneck |

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

# Roadmap

## Phase 1 - Foundation
✅ Reflection Memory.<br>
✅ Historical Bug Memory.<br> 
✅ Knowledge Retrieval.<br> 
✅ Context Builder.<br>
✅ Prompt Builder.<br>  

## Phase 2 - Developer Agent v1

✅ Reflection Retriever.<br>
✅ Historical Bug Retriever.<br>
✅ Knowledge Retriever.<br>
✅ Knowledge Context.<br>
✅ Search Symbol Tool.<br>
✅ Read File Tool.<br>
✅ Code Retrieval.<br>
✅ Code Context.<br>
✅ Analysis Generation.<br>
✅ Patch Candidate.<br>
✅ Diff Generator.<br>
✅ Patch Applier.<br>
❌ Feature Development

**Milestone:** Developer Agent can fix bug and apply changes now

## Phase 3 - BA Agent
✅ Generate requirements.<br>
✅ Generate epics.<br>
✅ Generate story.<br>
✅ Orchestrator workflow.<br>
❌ Publish to Confluence<br>
    - Formatted the page<br>
    - Using Remote MCP Confluence (Go SDK: https://github.com/modelcontextprotocol/go-sdk)<br>
    - Authenticated with OAuth 2.0<br>
    - Issue report/discussion: https://community.atlassian.com/forums/Rovo-questions/Atlassian-Remote-MCP-tools-fail-during-execution-despite/qaq-p/3253923#M4876<br>
❌ Human in loop<br>
**Milestone:** BA Agent can create requirement, epics, story

## Phase 4 - Planner Agent

## Phase 5 - BA Agent and Scheduler Agent

## Phase 6 - Reviewer Agent

## Phase 7 - Tester Agent

## Phase 8 - Judge Agent

## Phase 9 - AI Software Company

**Note:** Adjust Code RAG to Code Tool
