// Demonstrates automatic Anthropic instrumentation using the TruLayer Go SDK.
//
// Wrapping the Anthropic client with InstrumentAnthropic causes every
// Messages.New call to emit a TruLayer span automatically — no manual
// span management required. The example runs without network access
// when ANTHROPIC_API_KEY is unset (the wrapped call short-circuits
// before hitting the network), but the auto-instrumentation path
// still executes so spans are produced.
package main

import (
	"context"
	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	instruments "github.com/trulayer/client-go/instruments/anthropic"
	"github.com/trulayer/client-go/trulayer"
)

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	acClient := anthropic.NewClient(option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")))
	client := instruments.InstrumentAnthropic(&acClient, tl)

	trace, ctx := tl.NewTrace(ctx, "anthropic-auto-example")
	trace.SetInput("Why is the sky blue?")
	defer trace.End(ctx)

	resp, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_5HaikuLatest,
		MaxTokens: 256,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("Why is the sky blue?")),
		},
	})
	if err != nil {
		log.Printf("completion error: %v", err)
		return
	}
	if len(resp.Content) > 0 {
		text := resp.Content[0].Text
		trace.SetOutput(text)
		log.Printf("response: %s", text)
	}

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
}
