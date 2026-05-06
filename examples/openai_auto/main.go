// Demonstrates wrapping an OpenAI-style call inside a TruLayer span.
//
// The example does not call OpenAI directly; the OpenAI SDK surface is
// approximated with a stub so the example is runnable without network
// access or an API key. Replace the stub with the real OpenAI client in
// your own code.
package main

import (
	"context"
	"log"
	"os"

	"github.com/trulayer/client-go/trulayer"
)

// callOpenAI is a stand-in for `client.Chat.Completions.New(ctx, params)`.
// In real code, replace this with a call to the OpenAI Go SDK.
func callOpenAI(ctx context.Context, prompt string) (string, error) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		return "stubbed response (set OPENAI_API_KEY to use the real SDK)", nil
	}
	// TODO: replace with the real OpenAI SDK call.
	return "stubbed response", nil
}

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	t, ctx := tl.NewTrace(ctx, "openai-auto-example")
	t.SetInput("Why is the sky blue?")

	span, spanCtx := t.NewSpan(ctx, "openai.chat.completions",
		trulayer.SpanTypeLLM,
		trulayer.WithSpanModel("gpt-4o-mini"),
		trulayer.WithSpanInput("Why is the sky blue?"),
	)
	out, err := callOpenAI(spanCtx, "Why is the sky blue?")
	if err != nil {
		span.SetError(err.Error())
	} else {
		span.SetOutput(out)
	}
	span.End(spanCtx)

	t.SetOutput(out)
	t.End(ctx)

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
}
