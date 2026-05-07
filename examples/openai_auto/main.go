// Demonstrates automatic OpenAI instrumentation using the TruLayer Go SDK.
//
// Wrapping the OpenAI client with InstrumentOpenAI causes every chat
// completion call to emit a TruLayer span automatically — no manual
// span management required. The example runs without network access
// when OPENAI_API_KEY is unset (the wrapped call short-circuits before
// hitting the network), but the auto-instrumentation path still
// executes so spans are produced.
package main

import (
	"context"
	"log"
	"os"

	"github.com/openai/openai-go"
	instruments "github.com/trulayer/client-go/instruments/openai"
	"github.com/trulayer/client-go/trulayer"
)

func main() {
	ctx := context.Background()
	tl := trulayer.NewClient(os.Getenv("TRULAYER_API_KEY"))
	defer func() { _ = tl.Shutdown(ctx) }()

	oaiClient := openai.NewClient()
	client := instruments.InstrumentOpenAI(&oaiClient, tl)

	trace, ctx := tl.NewTrace(ctx, "openai-auto-example")
	trace.SetInput("Why is the sky blue?")
	defer trace.End(ctx)

	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4oMini,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Why is the sky blue?"),
		},
	})
	if err != nil {
		log.Printf("completion error: %v", err)
		return
	}
	if len(resp.Choices) > 0 {
		trace.SetOutput(resp.Choices[0].Message.Content)
		log.Printf("response: %s", resp.Choices[0].Message.Content)
	}

	if err := tl.Flush(ctx); err != nil {
		log.Printf("flush: %v", err)
	}
}
