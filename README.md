# TruLayer AI — Go Demos

Runnable, end-to-end Go examples that show how to trace AI applications with the
[`github.com/trulayer/client-go`](https://pkg.go.dev/github.com/trulayer/client-go) SDK.
Every example emits real traces and spans; the `feedback` demo also posts user feedback against a trace.

## Quick start

```bash
# From this directory:
cp .env.example .env      # then fill in your keys
go mod tidy
go run ./examples/basic_trace/
```

Set in `.env` at minimum:

```
TRULAYER_API_KEY=tl_...
TRULAYER_ENDPOINT=https://api.trulayer.ai
OPENAI_API_KEY=sk-...
ANTHROPIC_API_KEY=sk-ant-...
```

## Offline / CI mode

Set `TRULAYER_DRY_RUN=true` and all HTTP calls to TruLayer are no-oped — no API key required,
no network touched. Used by the smoke test suite.

```bash
TRULAYER_DRY_RUN=true go run ./examples/basic_trace/
TRULAYER_DRY_RUN=true go test ./smoke/
```

## Examples

| Directory | Shows |
|-----------|-------|
| `examples/basic_trace/` | Manual `NewTrace` + `NewSpan` with a real OpenAI call |
| `examples/openai_auto/` | OpenAI auto-instrumentation — wraps `github.com/openai/openai-go` |
| `examples/anthropic_auto/` | Anthropic auto-instrumentation — wraps `github.com/anthropics/anthropic-sdk-go` |
| `examples/rag_pipeline/` | Embed → retrieve → generate, three span types in one trace |
| `examples/agent/` | Tool-calling agent loop, one span per tool + one per LLM turn |
| `examples/feedback/` | Trace an answer and attach a thumbs-up feedback record |

Each example is a standalone `package main` and runs with:

```bash
go run ./examples/basic_trace/
go run ./examples/openai_auto/
go run ./examples/anthropic_auto/
go run ./examples/rag_pipeline/
go run ./examples/agent/
```

## Tests

```bash
go test ./...
```

The smoke test suite compiles and runs every example with `TRULAYER_DRY_RUN=true`,
asserting each exits cleanly with no panics.

## Project layout

```
demo-go/
├── .env.example
├── go.mod
├── go.sum
├── examples/
│   ├── basic_trace/
│   │   └── main.go     # manual trace + span creation
│   ├── openai_auto/
│   │   └── main.go     # OpenAI auto-instrumentation (instruments/openai)
│   ├── anthropic_auto/
│   │   └── main.go     # Anthropic auto-instrumentation (instruments/anthropic)
│   ├── rag_pipeline/
│   │   └── main.go     # multi-span RAG pipeline
│   ├── agent/
│   │   └── main.go     # tool-calling agent tracing
│   └── feedback/
│       └── main.go     # submitting feedback on a trace
└── smoke/
    └── smoke_test.go   # CI smoke tests (dry-run, no network)
```

## Links

- [Go SDK](https://github.com/trulayer/client-go)
- [Documentation](https://docs.trulayer.ai)
- [Go SDK reference](https://docs.trulayer.ai/sdks/go/reference)
