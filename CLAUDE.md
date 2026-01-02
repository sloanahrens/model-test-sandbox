# Sandbox

Test repository for evaluating local MLX models on development tasks.

## Structure

```
sandbox/
├── packages/utils/     # TypeScript utilities (pnpm workspace)
├── cmd/api/           # Go API server
├── internal/handler/  # Go HTTP handlers
└── docs/plans/        # Design documents
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

| Branch | Purpose |
|--------|---------|
| `test/f1-simple-commit` | Single file addition (validatePhone) |
| `test/f5-type-error` | Type error in sleep() return type |
| `test/q1-multi-file-refactor` | Logger module + retry integration |
| `test/q5-logic-bug` | Subtle bug in backoff calculation |

---

## MLX Model Evaluation Results

**Date**: 2026-01-01
**Goal**: Find optimal models for fast and quality tiers to offload development tasks

### Models Tested

| Model | Size | Memory | Speed (tok/s) |
|-------|------|--------|---------------|
| Llama-3.2-1B-Instruct-4bit | 680M | 0.9 GB | 392 |
| Llama-3.2-3B-Instruct-4bit | 1.7G | 2.1 GB | 168 |
| Mistral-7B-Instruct-v0.3-4bit | 3.8G | 4.4 GB | 82 |
| DeepSeek-Coder-V2-Lite-Instruct-4bit | 8.2G | 9.1 GB | 143 |

### Benchmark Results

| Test | Task | Llama-1B | Llama-3B | Mistral-7B | DeepSeek |
|------|------|----------|----------|------------|----------|
| F1 | Commit message | MARGINAL | MARGINAL | MARGINAL | **PASS** |
| F2 | Code explanation | PASS | PASS | PASS | PASS |
| F5 | Fix type error | FAIL | **PASS** | - | MARGINAL |

#### Detailed Findings

**F1 - Commit Message Generation**
- Prompt: Generate conventional commit for `validatePhone` addition
- Llama-1B: Wrong format ("Added US phone number validation")
- Llama-3B: Correct format, wrong type (`fix` instead of `feat`)
- Mistral-7B: Inverted format (`utils(validators):` instead of `feat(validators):`)
- DeepSeek: Correct (`feat(utils): add phone number validation function`)

**F2 - Code Explanation**
- Prompt: Explain the `retry` function in under 50 words
- All models passed with accurate explanations
- DeepSeek and Mistral provided more detailed responses

**F5 - Type Error Fix**
- Prompt: Fix `Promise<string>` return type on function returning void
- Llama-1B: FAIL - Changed behavior to return string instead of fixing type
- Llama-3B: PASS - Correctly changed to `Promise<void>`
- DeepSeek: MARGINAL - Tried to make function return string

### Final Decision (2026-01-01)

**Single Model: DeepSeek-Coder-V2-Lite-Instruct-4bit**
- Speed: 143 tok/s
- Memory: 9.1 GB
- Use for: Fast local tasks (explanations, commit messages, simple code gen)

**Quality tasks: Claude (native)**
- All code review, complex generation, and critical decisions go through Claude
- No local quality model needed - simpler setup, better results

**Deleted:**
- Llama-3.2-1B, Llama-3.2-3B, Mistral-7B (all removed)
- All incomplete stubs (removed)

### Dogfooding Plan

Testing DeepSeek-Coder in real development over coming days:
- Track success/failure on actual tasks
- Note where it helps vs where Claude is needed
- Adjust usage patterns based on experience
