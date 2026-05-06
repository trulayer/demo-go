// Demonstrates a multi-span trace covering a retrieval step followed by
// an LLM call — the canonical RAG pipeline shape.
package main

import (
	"context"
	"log"
	"os"

	"github.com/trulayer/client-go/trulayer"
)

func retrieveDocs(_ context.Context, _ string) []string {
	return []string{"doc-1", "doc-2", "doc-3"}
}

func callLLM(_ context.Context, _ string, _ []string) string {
	return "Paris, France."
}

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	question := "What is the capital of France?"

	t, ctx := tl.NewTrace(ctx, "rag-pipeline")
	t.SetInput(question)

	rs, rctx := t.NewSpan(ctx, "vector-retrieval", trulayer.SpanTypeRetrieval)
	docs := retrieveDocs(rctx, question)
	rs.SetOutput("retrieved 3 documents")
	rs.End(rctx)

	ls, lctx := t.NewSpan(ctx, "llm-answer", trulayer.SpanTypeLLM,
		trulayer.WithSpanModel("gpt-4o-mini"),
		trulayer.WithSpanInput(question),
	)
	answer := callLLM(lctx, question, docs)
	ls.SetOutput(answer)
	ls.SetTokens(120, 8)
	ls.End(lctx)

	t.SetOutput(answer)
	t.End(ctx)

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
}
