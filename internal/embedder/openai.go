package embedder

import "context"

// OpenAIEmbedder calls the OpenAI embeddings API.
// It uses net/http directly — no SDK — to keep the dependency surface thin.
type OpenAIEmbedder struct {
	apiKey string
	model  string
}

// NewOpenAI creates an OpenAIEmbedder.
func NewOpenAI(apiKey, model string) *OpenAIEmbedder {
	panic("not implemented")
}

// Embed implements Embedder.
func (e *OpenAIEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	panic("not implemented")
}
