// Package checkpoint persists the last acknowledged replication LSN (and, in
// v0.2, content hashes per primary key) to a local SQLite database.
package checkpoint

// Store reads and writes checkpoint state to SQLite.
type Store struct {
	path string
}

// New opens (or creates) the SQLite database at path and runs any pending
// schema migrations.
func New(path string) (*Store, error) {
	panic("not implemented")
}

// LoadLSN returns the last saved LSN and whether one exists.
func (s *Store) LoadLSN() (lsn uint64, ok bool, err error) {
	panic("not implemented")
}

// SaveLSN persists the given LSN atomically.
func (s *Store) SaveLSN(lsn uint64) error {
	panic("not implemented")
}

// Close closes the underlying database connection.
func (s *Store) Close() error {
	panic("not implemented")
}
