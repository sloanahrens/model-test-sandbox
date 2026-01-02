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

### Recommendations

**Fast Tier: Llama-3.2-3B-Instruct-4bit**
- Reason: Passes error fixes (1B fails), still fast at 168 tok/s
- Use for: Simple explanations, type fixes, quick code gen
- Memory: 2.1 GB (acceptable)

**Quality Tier: DeepSeek-Coder-V2-Lite-Instruct-4bit**
- Reason: Best on commit messages, faster than Mistral (143 vs 82 tok/s)
- Use for: Commit messages, complex code gen, detailed explanations
- Memory: 9.1 GB (higher but worth it)

**Delete (not needed):**
- Llama-3.2-1B: Too unreliable on basic tasks
- Mistral-7B: Slower than DeepSeek with worse commit format

### Cleanup Commands

```bash
# Delete unused models
rm -rf ~/.cache/huggingface/hub/models--mlx-community--Llama-3.2-1B-Instruct-4bit
rm -rf ~/.cache/huggingface/hub/models--mlx-community--Mistral-7B-Instruct-v0.3-4bit

# Delete incomplete downloads (8K stubs)
rm -rf ~/.cache/huggingface/hub/models--mlx-community--DeepSeek-R1-Distill-Qwen-32B-4bit
rm -rf ~/.cache/huggingface/hub/models--mlx-community--DeepSeek-R1-Distill-Qwen-7B-4bit
rm -rf ~/.cache/huggingface/hub/models--mlx-community--Llama-3.3-70B-Instruct-4bit
rm -rf ~/.cache/huggingface/hub/models--mlx-community--Meta-Llama-3.1-8B-Instruct-4bit
rm -rf ~/.cache/huggingface/hub/models--mlx-community--Meta-Llama-3.1-8B-Instruct-8bit
```

### Final Model Configuration

After cleanup, keep these models:

| Tier | Model | Speed | Use Cases |
|------|-------|-------|-----------|
| Fast | Llama-3.2-3B-Instruct-4bit | 168 tok/s | Explanations, simple fixes |
| Quality | DeepSeek-Coder-V2-Lite-4bit | 143 tok/s | Commits, code gen |

Total disk usage: ~10 GB (down from ~14 GB + stubs)
