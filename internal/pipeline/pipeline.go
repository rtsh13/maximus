// Package pipeline defines the core types and interfaces that flow through
// the Maximus capture → embedder → sink pipeline.
package pipeline

import (
	"time"
)

// Op is the type of database change operation.
type Op string

const (
	OpInsert Op = "INSERT"
	OpUpdate Op = "UPDATE"
	OpDelete Op = "DELETE"
)

// ChangeEvent is the internal representation of a single row-level change.
// It is produced by the capture stage and consumed by the embedder stage.
type ChangeEvent struct {
	// LSN is the log sequence number of the commit that produced this event.
	// Used for checkpointing.
	LSN uint64

	Op  Op
	Table string

	// OldRow is the previous state of the row. Nil for INSERT.
	OldRow map[string]any

	// NewRow is the new state of the row. Nil for DELETE.
	NewRow map[string]any

	// PrimaryKey holds the value(s) of the primary key column(s).
	PrimaryKey []any

	// CommittedAt is the wall-clock time of the commit on the source.
	CommittedAt time.Time
}

// EmbedResult is the output of the embedder stage.
type EmbedResult struct {
	LSN        uint64
	PrimaryKey []any
	// Vector is nil for DELETE events (no embedding needed).
	Vector []float32
	Op     Op
}
