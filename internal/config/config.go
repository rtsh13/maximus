// Package config loads and validates the Maximus TOML configuration file.
package config

// Source holds the Postgres connection parameters.
type Source struct {
	DSN         string `toml:"dsn"`
	Slot        string `toml:"slot"`
	Publication string `toml:"publication"`
}

// Sink holds the Qdrant connection parameters.
type Sink struct {
	URL        string `toml:"url"`
	Collection string `toml:"collection"`
}

// Embedder holds the embedding API parameters.
type Embedder struct {
	Provider string `toml:"provider"`
	APIKey   string `toml:"api_key"`
	Model    string `toml:"model"`
}

// Mapping describes how one Postgres table maps to a vector collection.
type Mapping struct {
	Table            string   `toml:"table"`
	PrimaryKey       string   `toml:"primary_key"`
	Template         string   `toml:"template"`
	Dimension        int      `toml:"dimension"`
	EmbeddingColumns []string `toml:"embedding_columns"` // v0.2
}

// Config is the top-level configuration struct.
type Config struct {
	Source   Source    `toml:"source"`
	Sink     Sink      `toml:"sink"`
	Embedder Embedder  `toml:"embedder"`
	Mappings []Mapping `toml:"mapping"`
}

// Load reads and parses a TOML configuration file from the given path.
func Load(path string) (*Config, error) {
	panic("not implemented")
}

// Validate checks that required fields are present and values are sensible.
func (c *Config) Validate() error {
	panic("not implemented")
}
