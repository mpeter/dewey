# Tasks: Pluggable Providers

All tasks are implementation-complete. This document records the work done.

## 1. Provider Config and Factories
- [x] ProviderConfig struct in `embed/provider.go` and `llm/provider.go`
- [x] `embed.NewEmbedderFromConfig()` factory with ollama/vertex switch
- [x] `llm.NewSynthesizerFromConfig()` factory with ollama/vertex switch
- [x] `embed.ReadEmbeddingConfig()` with global config fallback and env var overrides
- [x] `llm.ReadSynthesisConfig()` with global config fallback and compile_model backward compat
- [x] Tests for both factories and config readers

## 2. Vertex AI Embedder
- [x] `embed/vertex.go` — VertexEmbedder with Embed, EmbedBatch, Available, ModelID
- [x] OAuth via golang.org/x/oauth2/google (pure Go, no CGO)
- [x] Token function injection for testability
- [x] `embed/vertex_test.go` — httptest tests (success, batch, auth error, org policy)

## 3. Vertex AI Synthesizer
- [x] `llm/vertex.go` — VertexSynthesizer with Synthesize, Available, ModelID
- [x] Anthropic Messages format via rawPredict endpoint
- [x] Clear error messages for common failures
- [x] `llm/vertex_test.go` — httptest tests (success, auth error, model not found)

## 4. store_compiled MCP Tool
- [x] StoreCompiledInput type in `types/tools.go`
- [x] StoreCompiled handler in `tools/compile.go`
- [x] compiled_by frontmatter, tag validation, overwrite support
- [x] Registered in `server.go`
- [x] Tests in `tools/compile_test.go`

## 5. Consolidate Construction Sites
- [x] cli.go compile/curate use factory functions
- [x] cli.go createIndexEmbedder uses ReadEmbeddingConfig + factory
- [x] main.go backgroundCuration uses ReadSynthesisConfig + factory
- [x] Removed readCompileModel()
- [x] Preserved DEWEY_GENERATION_MODEL and DEWEY_EMBEDDING_* env var fallbacks

## 7. Rate Limiting (429 Retry)
- [x] Add exponential backoff retry on HTTP 429 to `embed/vertex.go` EmbedBatch
- [x] Add exponential backoff retry on HTTP 429 to `llm/vertex.go` Synthesize
- [x] Add Retry-After header parsing to both providers
- [x] Add tests: 429 then success, 429 exhaustion, Retry-After header respected, context cancellation
- [x] Update AGENTS.md with rate limiting documentation

## 6. Documentation
- [x] AGENTS.md updated with provider configuration section
- [x] GoDoc on all new exported types and functions
- [x] Website issue filed: unbound-force/website#113

## Migration Evidence

- Migrated to the OpenSpec archive on 2026-11-13.
- Archive date and metadata creation date use the original feature merge date (2026-05-08).
- Merged PR: https://github.com/unbound-force/dewey/pull/56 (merged 2026-05-08T16:23:47Z; merge SHA `6c94aa7e5e9d65c1e7650d96f01a2d6d5c9efe66`).
- A feature-level merge is not proof of completion of individual tasks. Original task checkbox states are preserved; unchecked tasks remain unverified.
- These documents are historical artifacts, not active workflow instructions or current-schema validation targets.
