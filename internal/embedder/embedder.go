// Package embedder defines the Embedder interface and provides implementations.
package embedder

import "context"

// Embedder converts a text string into a dense vector embedding.
type Embedder interface {
	// Embed returns a float32 vector for the given text.
	Embed(ctx context.Context, text string) ([]float32, error)
}
