// Demonstrates submitting feedback against a previously ingested trace.
package main

import (
	"context"
	"log"
	"os"

	"github.com/trulayer/client-go/trulayer"
)

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	// First, produce a trace we can attach feedback to.
	t, _ := tl.NewTrace(ctx, "feedback-example")
	t.SetInput("ping")
	t.SetOutput("pong")
	t.End(ctx)
	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}

	// Then submit feedback using its ID.
	if err := tl.SubmitFeedback(ctx, t.ID(), trulayer.FeedbackData{
		Label:   "good",
		Comment: "looked great",
	}); err != nil {
		log.Printf("feedback: %v", err)
	}
}
