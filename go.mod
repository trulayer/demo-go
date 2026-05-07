module github.com/trulayer/demo-go

go 1.22

require (
	github.com/anthropics/anthropic-sdk-go v0.2.0-beta.3
	github.com/sashabaranov/go-openai v1.32.0
	github.com/trulayer/client-go v0.0.0
	github.com/trulayer/client-go/instruments/anthropic v0.0.0
	github.com/trulayer/client-go/instruments/openai v0.0.0
)

replace (
	github.com/trulayer/client-go => ../client-go
	github.com/trulayer/client-go/instruments/anthropic => ../client-go/instruments/anthropic
	github.com/trulayer/client-go/instruments/openai => ../client-go/instruments/openai
)
