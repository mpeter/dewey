# Tasks: Doctor Emoji Markers

**Input**: Design documents from `openspec/changes/archive/2026-03-30-doctor-emoji-markers/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, quickstart.md

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Phase 1: Implementation — Consistent Diagnostic Output (Priority: P1) 🎯 MVP

**Goal**: Replace `[PASS]`/`[WARN]`/`[FAIL]` text markers with emoji markers (`✅`/`⚠️`/`❌`) in `printCheck()` and update all test assertions to match.

**Independent Test**: Run `dewey doctor --vault .` and confirm all check lines use emoji markers, not text markers.

- [x] T001 [P] [US1] Update `printCheck()` format string to use emoji markers via switch mapping in `cli.go` (~line 1124)
- [x] T002 [P] [US1] Update `printCheck` GoDoc comment to reflect emoji format in `cli.go` (~lines 1109-1114)
- [x] T003 [P] [US1] Update test assertions in `TestDoctorCounter_PrintCheck`, `TestDoctorCmd_WithInitializedVault`, and `TestDoctorCmd_MissingDeweyDir` to expect emoji markers instead of text markers in `cli_test.go` (6 assertions)
- [x] T005 [P] [US2] Remove misleading `--no-embeddings` flag from Fix hint for 0-page sources in `cli.go` (~line 1267, FR-010)
- [x] T006 [P] [US1] Apply ANSI dim escape codes to all 6 section headings in `cli.go` (FR-011)

**Checkpoint**: `go test -race -count=1 ./...` passes. All check lines display emoji markers.

---

## Phase 2: Verification

**Purpose**: Confirm build, lint, and tests pass — CI parity gate.

- [x] T004 [US1] Run `go build ./...`, `go vet ./...`, and `go test -race -count=1 ./...` to verify no regressions

---

## Dependencies & Execution Order

### Phase Dependencies

- **Phase 1**: T001, T002, and T003 are all `[P]` — they touch different files (`cli.go` vs `cli_test.go`) or different sections of the same file (format string vs GoDoc comment), so they can run in parallel.
- **Phase 2**: Depends on Phase 1 completion. T004 validates the combined result.

### Parallel Opportunities

T001 and T002 both modify `cli.go` but at different line ranges (~30 lines apart). T003 modifies `cli_test.go` only. All three can be executed in parallel by a single implementer working through them sequentially, or T003 can be assigned to a separate worker.

---

## Notes

- Total scope: 2 files modified (`cli.go`, `cli_test.go`), ~15 lines production code, ~10 lines test assertions
- No new files, no new dependencies, no data model changes
- The `Fix:` hint indentation (US2/P2) already uses 5-space indent — no code change needed (confirmed in research.md)
- The summary box already uses emoji — no change needed (FR-007)
- The `printCheck()` method signature is unchanged (FR-009)
<!-- spec-review: passed -->

## Migration Evidence

- Migrated to the OpenSpec archive on 2026-11-13.
- Archive date and metadata creation date use the original feature merge date (2026-03-30).
- Merged PR: https://github.com/unbound-force/dewey/pull/21 (merged 2026-03-30T13:01:19Z; merge SHA `ee8b12d4ddd8db2f4ce2d590c3f08c5b49a6efbe`).
- A feature-level merge is not proof of completion of individual tasks. Original task checkbox states are preserved; unchecked tasks remain unverified.
- These documents are historical artifacts, not active workflow instructions or current-schema validation targets.
