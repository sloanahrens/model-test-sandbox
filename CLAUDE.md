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

| Tier | Model | Size | Speed | Accuracy |
|------|-------|------|-------|----------|
| **Fast** | Qwen2.5-Coder-14B-Instruct-4bit | 7.7 GB | 15-50 tok/s | 100% |
| **Quality** | Qwen2.5-Coder-32B-Instruct-4bit | 17.2 GB | 6-23 tok/s | 100% |

**Total disk usage**: ~25 GB

### Benchmark Results

| Test | Task | DeepSeek | Qwen-14B | Qwen-32B |
|------|------|----------|----------|----------|
| F1 | Commit message | **FAIL** | PASS | PASS |
| F5 | Fix type error | **FAIL** | PASS | PASS |
| Q5 | Logic bug detection | MARGINAL | PASS | PASS |

### Test Details

**F1 - Commit Message Generation**
- Prompt: Generate conventional commit for `validatePhone` addition
- DeepSeek: `chore(phone-validation):...` - wrong type
- Qwen-14B: `feat(validator): add US phone number validation function` - correct
- Qwen-32B: `feat(validation): add US phone number validation function` - correct

**F5 - Type Error Fix**
- Prompt: Fix `Promise<string>` return type on function returning void
- DeepSeek: Returned code unchanged - FAIL
- Qwen-14B: Changed to `Promise<void>` - correct
- Qwen-32B: Changed to `Promise<void>` - correct

**Q5 - Logic Bug Detection**
- Prompt: Find bug in backoff calculation (`attempt` vs `attempt-1`)
- DeepSeek: Confused explanation but correct fix - MARGINAL
- Qwen-14B: Clear explanation + correct fix - PASS
- Qwen-32B: Concise explanation + correct fix - PASS

### Speed Measurements

| Model | Cold Start | Warm |
|-------|------------|------|
| Qwen-14B | 14.9 tok/s | 49.7 tok/s |
| Qwen-32B | 5.7 tok/s | 22.6 tok/s |

### Model Routing

```yaml
fast_tier:
  model: mlx-community/Qwen2.5-Coder-14B-Instruct-4bit
  use_for:
    - commit_messages
    - type_fixes
    - code_explanation
    - bug_detection
    - documentation

quality_tier:
  model: mlx-community/Qwen2.5-Coder-32B-Instruct-4bit
  use_for:
    - complex_refactoring
    - architectural_review
    - when_14B_uncertain

cloud_tier:
  model: claude-sonnet
  use_for:
    - security_review
    - architecture_decisions
    - multi-repo_analysis
```

### Deprecated Models

These models were tested and removed:

| Model | Size | Why Removed |
|-------|------|-------------|
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

### Scoring

| Score | Meaning |
|-------|---------|
| PASS | Usable as-is or trivial edit |
| MARGINAL | Mostly right, needs cleanup |
| FAIL | Wrong, broken, unusable |

### Key Learnings

1. **Code-specific training matters**: Qwen2.5-Coder outperforms general models on commit format and TypeScript semantics
2. **Size isn't everything**: Qwen-14B (8GB) beats DeepSeek (9GB) despite being smaller
3. **Warm-up effect**: First inference slower as weights load; subsequent calls 2-3x faster
4. **Speed vs accuracy**: Qwen-14B offers best balance for most tasks
