# Maximus

> **Early WIP — no announcement yet.**

Maximus is a single-binary, open source CDC tool that keeps vector stores synchronized with Postgres in near real time — without Kafka, without triggers, and without re-embedding rows whose semantically relevant columns have not changed.

## What it does

- Captures `INSERT`, `UPDATE`, and `DELETE` events from Postgres via logical replication (`pgoutput`).
- Builds an embedding text string from a user-supplied template.
- Calls an embedding API (OpenAI `text-embedding-3-small` in v0.1) only when needed.
- Upserts or deletes vectors in Qdrant, keyed by the row's primary key.
- Checkpoints the last acknowledged LSN to SQLite so it resumes cleanly after a crash.
- Configures entirely from a single TOML file — no code required.

## What it is not

- Not a vector database.
- Not a Debezium replacement for warehouse loading.
- Not a general-purpose ETL tool.
- Not enterprise-licensed. Maximus is and will remain open source (MIT).

## Status

| Version | Status | Target |
|---------|--------|--------|
| v0.1    | In progress | End of May 2026 |
| v0.2    | Planned | End of June 2026 |

## Quickstart

> Not available yet. Coming in v0.2 once the binary is stable.

```toml
# examples/config.toml — see the file for annotated fields
[source]
dsn         = "postgres://user:pass@localhost:5432/mydb"
slot        = "driftwatch_slot"
publication = "driftwatch_pub"

[sink]
url        = "http://localhost:6333"
collection = "my_collection"

[embedder]
provider = "openai"
api_key  = "sk-..."
model    = "text-embedding-3-small"

[[mapping]]
table       = "documents"
primary_key = "id"
template    = "{{.title}} {{.body}}"
dimension   = 1536
```

```sh
driftwatch run --config config.toml
```

## Requirements

- Postgres 14+ with `wal_level = logical`.
- A running Qdrant instance.
- An OpenAI API key (v0.1).

## Development

```sh
# Start local Postgres + Qdrant
docker compose up -d

# Apply test schema
psql "$DSN" -f dev/setup.sql

# Build
make build

# Run tests
make test
```

## License

MIT
