package sink

import "context"

// QdrantSink writes vectors to a Qdrant collection via the Qdrant Go client.
type QdrantSink struct {
	url        string
	collection string
}

// NewQdrant creates a QdrantSink.
func NewQdrant(url, collection string) *QdrantSink {
	panic("not implemented")
}

// Upsert implements Sink.
func (q *QdrantSink) Upsert(ctx context.Context, key any, vector []float32) error {
	panic("not implemented")
}

// Delete implements Sink.
func (q *QdrantSink) Delete(ctx context.Context, key any) error {
	panic("not implemented")
}
