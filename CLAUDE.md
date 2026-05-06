# CLAUDE.md — Go Demos (demo-go)

## Project Purpose

End-to-end runnable Go examples demonstrating TruLayer AI SDK integration. Used for developer onboarding, documentation, and SDK integration testing in CI.

## Tech Stack

- Go 1.22+
- `github.com/trulayer/client-go` — the TruLayer Go SDK (from `client-go`)
- Standard Go AI provider clients (OpenAI Go SDK, Anthropic Go SDK)

## Merge Conflict Policy

**Merge conflicts are the engineer's responsibility.** Before opening a PR (and again before merging), rebase onto the latest `main` and resolve all conflicts:

```bash
git fetch origin && git rebase origin/main
```

Do not open a PR with a conflicting branch. If a conflict arises after the PR is open because `main` moved, the PR author owns the rebase — not the reviewer or TPM.

## Branch Management

**Delete the feature branch after the PR is squash-merged.** Run `git push origin --delete <branch-name>` or click "Delete branch" in the GitHub UI immediately after merge.

## Key Commands

```bash
go mod tidy                          # Sync dependencies
go build ./...                       # Verify all examples compile
go run examples/basic_trace/main.go  # Run a single example
go test ./...                        # Run CI smoke tests
```

## Project Layout

```text
examples/
  basic_trace/
    main.go           → manual trace + span creation
  openai_auto/
    main.go           → OpenAI auto-instrumentation
  rag_pipeline/
    main.go           → multi-span RAG pipeline
  agent/
    main.go           → tool-calling agent tracing
  async_example/
    main.go           → concurrent goroutine tracing
  feedback/
    main.go           → submitting feedback on traces
smoke/
  smoke_test.go       → CI smoke test: compile + run each example
go.mod
go.sum
```

## Example Standards

- Every example is a standalone `package main` that compiles and runs with `go run examples/<name>/main.go`
- Each `main.go` has a package-level comment: one sentence explaining what it demonstrates
- No business logic — pure demonstration of one SDK concept per file
- Use real API calls in integration CI; set `TRULAYER_DRY_RUN=true` for offline runs (SDK no-ops all HTTP)
- Examples must compile with `go build ./...` — a compilation failure is a CI failure

## CI Smoke Tests

`smoke/smoke_test.go` builds and runs each example binary with:
- `TRULAYER_DRY_RUN=true` (no network calls)
- Mocked AI provider responses via environment variables or test servers

All examples must exit 0 without panics.

## CI is gating

Every pull request must pass CI before it can be merged. If CI fails, the engineer who opened the PR owns the fix. Don't bypass with `--admin` or `--no-verify`. If a check is flaky, fix it or remove it — don't skip it.

## Shell Conventions

**Never use `cd <path> && git <command>`** — use `git -C <path> <command>` instead.

## Public Repository Policy

This repository ships to TruLayer customers. Do not introduce references to internal code, internal repositories (e.g. the TruLayer API service or dashboard), internal planning documents, internal Linear issue content, or internal architectural details. Refer to the platform as "TruLayer" or "the TruLayer API" — not as specific internal components. If in doubt, leave it out.
