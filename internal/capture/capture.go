// Package capture subscribes to a Postgres logical replication slot and
// emits ChangeEvents onto a bounded channel.
package capture

import (
	"context"

	"github.com/maximus/internal/pipeline"
)

// Config holds the parameters the capture stage needs to connect to Postgres.
type Config struct {
	DSN         string
	Slot        string
	Publication string
	// StartLSN is the LSN to resume from. 0 means start from the current tip.
	StartLSN uint64
}

// Capture connects to Postgres logical replication and decodes pgoutput
// messages into ChangeEvents.
type Capture struct {
	cfg Config
}

// New creates a new Capture. It does not connect until Run is called.
func New(cfg Config) *Capture {
	panic("not implemented")
}

// Run starts consuming the replication stream and sends decoded ChangeEvents
// to out. It blocks until ctx is cancelled or a fatal error occurs.
func (c *Capture) Run(ctx context.Context, out chan<- pipeline.ChangeEvent) error {
	panic("not implemented")
}
