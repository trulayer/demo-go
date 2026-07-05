# Codex instructions

This is the public TruLayer Go demos repo. `CLAUDE.md` is the detailed source of truth; read it before making any non-trivial change.

## Scope

- Runnable Go examples for `github.com/trulayer/client-go`.
- Public customer-facing repo. Do not expose private service names, repo paths, planning issues, or private architecture.

## Working rules

- Make changes on a feature/fix branch and open a PR to `main`. Never commit directly to `main`.
- Keep examples focused: one SDK concept per example directory.
- Every example should support `TRULAYER_DRY_RUN=true` for offline CI.
- Keep examples aligned with the SDK and docs.

## Verification

Run before opening a PR:

```bash
go build ./...
go test ./...
```

For changed examples, also run the example directly, for example:

```bash
go run examples/basic_trace/main.go
```
