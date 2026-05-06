// Demonstrates an agent-style trace: a tool-call span nested inside an
// LLM span (the LLM decides to invoke a tool, which records its own
// span as a child).
package main

import (
	"context"
	"log"
	"os"

	"github.com/trulayer/client-go/trulayer"
)

func runTool(_ context.Context, _ string) string {
	return "the weather is sunny"
}

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	t, ctx := tl.NewTrace(ctx, "agent-example")
	t.SetInput("what is the weather?")

	llm, lctx := t.NewSpan(ctx, "llm-plan", trulayer.SpanTypeLLM,
		trulayer.WithSpanModel("gpt-4o-mini"),
	)

	// Tool call decided by the LLM — the tool span lives inside the
	// LLM span's context so the parent_span_id is set automatically.
	tool, tctx := t.NewSpan(lctx, "tool.weather_lookup", trulayer.SpanTypeTool)
	out := runTool(tctx, "san francisco")
	tool.SetOutput(out)
	tool.End(tctx)

	llm.SetOutput("called weather_lookup, got: " + out)
	llm.End(lctx)

	t.SetOutput("the weather is sunny")
	t.End(ctx)

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
}
