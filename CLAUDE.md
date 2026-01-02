# Sandbox

Test repository for evaluating local MLX models on development tasks.

## Structure

```
sandbox/
├── packages/utils/     # TypeScript utilities (pnpm workspace)
├── cmd/api/           # Go API server
└── internal/handler/  # Go HTTP handlers
```

## Commands

```bash
# TypeScript
pnpm install
pnpm test
pnpm lint
pnpm typecheck

# Go
go test ./internal/...
go build ./cmd/api
```

## Test Branches

| Branch | Purpose | Test |
|--------|---------|------|
| `test/f1-simple-commit` | Single file addition (validatePhone) | Commit message generation |
| `test/f5-type-error` | Type error in sleep() return type | Type fix accuracy |
| `test/q1-multi-file-refactor` | Logger module + retry integration | Multi-file understanding |
| `test/q5-logic-bug` | Subtle bug in backoff calculation | Logic bug detection |

---

## MLX Model Evaluation Results

**Last Updated**: 2026-01-02

### Current Configuration

| Tier | Model | Size | Speed | Use For |
|------|-------|------|-------|---------|
| **Local** | Qwen2.5-Coder-14B-Instruct-4bit | 7.7 GB | 50-56 tok/s | All local tasks |
| **Cloud** | Claude (Sonnet) | - | ~17s/call | Security, architecture, complex reasoning |

**Total disk usage**: ~8 GB (down from 25 GB)

### Why Single Model?

We tested Qwen-14B vs Qwen-32B on sophisticated tasks (2026-01-02):

| Test | Qwen-14B | Qwen-32B | Winner |
|------|----------|----------|--------|
| Multi-bug detection | Found off-by-one | Same | Tie |
| Race condition analysis | Clear explanation | More confused | **14B** |
| Rate limiter generation | Correct code | Has off-by-one bug | **14B** |
| Security vulnerability scan | Found 2/5 issues | Found 3/5 issues | 32B (slight) |
| Speed | 52-56 tok/s | 26-28 tok/s | **14B** (2x faster) |

**Conclusion**: 32B's extra parameters didn't improve reasoning. 14B was faster and often better.

### Benchmark Results (Basic Tests)

| Test | Task | DeepSeek | Qwen-14B |
|------|------|----------|----------|
| F1 | Commit message | **FAIL** | PASS |
| F5 | Fix type error | **FAIL** | PASS |
| Q5 | Logic bug detection | MARGINAL | PASS |

### Model Routing

```yaml
local_tier:
  model: mlx-community/Qwen2.5-Coder-14B-Instruct-4bit
  use_for:
    - commit_messages
    - type_fixes
    - code_explanation
    - bug_detection
    - documentation
    - code_generation

cloud_tier:
  model: claude-sonnet
  use_for:
    - security_review (local models miss critical vulns)
    - architecture_decisions
    - multi-repo_analysis
    - when_uncertain
```

### Deprecated Models

| Model | Size | Why Removed |
|-------|------|-------------|
| Qwen2.5-Coder-32B-Instruct-4bit | 17 GB | No better than 14B, 2x slower |
| Llama-3.3-70B-Instruct-8bit | 70 GB | Too slow (0.4 tok/s) |
| DeepSeek-Coder-V2-Lite-Instruct-4bit | 9 GB | Failed F1, F5 tests |
| Llama-3.2-1B/3B-Instruct-4bit | 0.7-1.7 GB | Unreliable on basic tasks |
| Mistral-7B-Instruct-v0.3-4bit | 3.8 GB | Slower than alternatives |

---

## Evaluation Framework

### Test Categories

**Fast Tier Tests (F-series)**
- F1: Commit message generation
- F2: Code explanation (<100 words)
- F3: Simple code generation
- F4: Test generation
- F5: Type error fix

**Quality Tier Tests (Q-series)**
- Q1: Multi-file commit message
- Q2: Complex code explanation
- Q3: Code gen with edge cases
- Q4: Integration test generation
- Q5: Logic bug detection

**Advanced Tests (A-series)** - For model comparison
- A1: Multi-bug detection (find 2+ interacting bugs)
- A2: Race condition identification
- A3: Complex code generation with constraints
- A4: Security vulnerability detection

### Scoring

| Score | Meaning |
|-------|---------|
| PASS | Usable as-is or trivial edit |
| MARGINAL | Mostly right, needs cleanup |
| FAIL | Wrong, broken, unusable |

### Key Learnings

1. **Bigger isn't always better**: Qwen-14B outperformed Qwen-32B on reasoning tasks
2. **Code-specific training matters**: Qwen2.5-Coder beats general models on TypeScript
3. **Speed compounds**: 2x faster means more iterations, better results
4. **Security needs Claude**: Both local models missed critical vulnerabilities (path traversal)
5. **Warm-up effect**: First inference slower; subsequent calls 2-3x faster
