// Demonstrates manual trace and span creation with the TruLayer Go SDK.
package main

import (
	"context"
	"log"
	"os"

	"github.com/trulayer/client-go/trulayer"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("TRULAYER_API_KEY")

	tl := trulayer.NewClient(apiKey)
	defer func() { _ = tl.Shutdown(ctx) }()

	t, ctx := tl.NewTrace(ctx, "basic-example")
	t.SetInput("hello")

	s1, ctx := t.NewSpan(ctx, "step-1", trulayer.SpanTypeOther)
	s1.SetOutput("step-1-done")
	s1.End(ctx)

	s2, ctx := t.NewSpan(ctx, "step-2", trulayer.SpanTypeOther)
	s2.SetOutput("step-2-done")
	s2.End(ctx)

	t.SetOutput("world")
	t.End(ctx)

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
	log.Printf("trace %s submitted", t.ID())
}
