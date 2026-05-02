# Maximus

> **Early WIP. No announcement yet.**

Maximus lets you keep a Qdrant vector store in sync with your Postgres database in near real time. You point it at a table, tell it which columns matter for embeddings, and it handles the rest: it watches the Postgres write-ahead log, skips rows whose relevant columns have not changed, calls the embedding API only when needed, and upserts or removes vectors in Qdrant automatically.

You configure everything in a single TOML file. You do not write any code.

## What you get

- Your vectors stay fresh within seconds of a row changing in Postgres.
- You stop paying to re-embed rows where only operational metadata changed (timestamps, counters, flags).
- You get a clean audit trail: every vector is keyed to the source row and the LSN that produced it.
- If Maximus crashes, it picks up from where it left off. No duplicates, no lost events.

## What Maximus is not

- Not a vector database.
- Not a replacement for Debezium if you are loading a data warehouse.
- Not a general-purpose ETL tool.
- Not enterprise-licensed. Maximus is MIT and will stay that way.

## Status

| Version | Status | Target |
|---------|--------|--------|
| v0.1 | In progress | End of May 2026 |
| v0.2 | Planned | End of June 2026 |

v0.1 gets the pipeline working end to end. v0.2 adds the column-aware skip logic that cuts embedding costs.

## Quickstart

Coming in v0.2. Until then, see `examples/config.toml` for the full annotated config reference.

```toml
[source]
dsn         = "postgres://user:pass@localhost:5432/mydb"
slot        = "maximus_slot"
publication = "maximus_pub"

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
maximus run --config config.toml
```

## Requirements

- Postgres 14+ with `wal_level = logical` and admin access to create a replication slot.
- A running Qdrant instance reachable on the network.
- An OpenAI API key (v0.1).

## Run it locally

```sh
# Start Postgres and Qdrant
docker compose up -d

# Create the test table and publication
psql "$DSN" -f dev/setup.sql

# Build the binary
make build

# Run the tests
make test
```

## License

MIT
