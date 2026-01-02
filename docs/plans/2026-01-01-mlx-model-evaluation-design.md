# MLX Model Evaluation Design

**Goal**: Evaluate local MLX models to offload simple development tasks, reducing latency while preserving quality for complex work.

## Model Candidates

| Tier | Candidates | Size |
|------|------------|------|
| **Fast** | Llama-3.2-1B-Instruct-4bit | 680M |
| **Fast** | Llama-3.2-3B-Instruct-4bit | 1.7G |
| **Quality** | Mistral-7B-Instruct-v0.3-4bit | 3.8G |
| **Quality** | DeepSeek-Coder-V2-Lite-Instruct-4bit | 8.2G |

## Evaluation Phases

### Phase 1: Synthetic Benchmarks

Run each model through 10 standardized tasks. Score outputs as Pass (2), Marginal (1), or Fail (0).

**Minimum thresholds:**
- Fast tier: ≥6/10 on fast tests
- Quality tier: ≥7/10 on quality tests

**Tiebreaker**: If both candidates pass, pick the smaller/faster one.

### Phase 2: Latency Measurement

Measure tokens/second on standard prompts. Record time-to-first-token and total generation time.

### Phase 3: Live Dogfooding (3-5 days)

Use surviving models for real work. Track success/failure rate. Target >80% success.

---

## Phase 1 Test Cases

### Fast Tier Tests

| # | Task | Prompt | Pass Criteria |
|---|------|--------|---------------|
| F1 | Commit msg | "Write commit message for staged changes" | Conventional format, accurate |
| F2 | Explain | "What does the `retry` function do?" | Correct, <100 words |
| F3 | Code gen | "Add `validatePhone(phone: string): boolean`" | Works, matches style |
| F4 | Test gen | "Add tests for `formatDate` and `parseDate`" | Tests run, cover edges |
| F5 | Error fix | "Fix the TypeScript error" | Compiles after fix |

### Quality Tier Tests

| # | Task | Prompt | Pass Criteria |
|---|------|--------|---------------|
| Q1 | Commit msg | "Write commit message" (multi-file) | Captures scope |
| Q2 | Explain | "Explain request flow from main.go to CreateItem" | Accurate, includes mutex |
| Q3 | Code gen | "Add pagination to ListItems (limit/offset)" | Works, handles edges |
| Q4 | Test gen | "Add integration test for POST /api/items" | Runs, tests HTTP flow |
| Q5 | Error fix | "Fix the bug" (retry backoff logic) | Correct fix |

---

## Scoring Rubric

| Score | Meaning | Weight |
|-------|---------|--------|
| Pass | Usable as-is or trivial edit | 2 |
| Marginal | Mostly right, needs cleanup | 1 |
| Fail | Wrong, broken, unusable | 0 |

---

## Post-Evaluation Actions

1. Delete losing models from `~/.cache/huggingface/hub/`
2. Delete incomplete downloads (8K stubs)
3. Document winners in workspace CLAUDE.md
4. Update slash commands to reference model tiers

---

## Test Branches

Create these branches in sandbox for test scenarios:

- `test/f1-simple-commit` - Single file addition for F1
- `test/q1-multi-file-refactor` - Multi-file change for Q1
- `test/f5-type-error` - Intentional TypeScript error for F5
- `test/q5-logic-bug` - Retry backoff bug for Q5
