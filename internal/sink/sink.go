// Package sink defines the Sink interface and provides implementations.
package sink

import "context"

// Sink writes embedding results to a vector store.
type Sink interface {
	// Upsert inserts or replaces the vector for the given key.
	Upsert(ctx context.Context, key any, vector []float32) error

	// Delete removes the vector for the given key.
	Delete(ctx context.Context, key any) error
}
