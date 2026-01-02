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

## Purpose

This repo exists to test local MLX models (Llama, Mistral, DeepSeek) on realistic development tasks:
- Commit message generation
- Code explanation
- Test generation
- Error fixing
- Simple refactoring
