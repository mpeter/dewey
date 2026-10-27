# UF Customizer Retirement Design

## Scope and Decision

Delete only the UF customizer command. Its instructions regenerate
other harness files; they are not the sole copy of Dewey's current
conventions or review requirements. Those deployed files remain
byte-for-byte unchanged. No convention relocation is needed here.

Keep the existing OpenSpec schema for this change's planning record;
changing the project's schema or constitution is a separate gap.
Do not rename source/module identities or remove legal provenance.

## Verification

The regression enumerates Markdown commands in `.opencode/command/`
and `.opencode/commands/`, the documented discovery directories.
It rejects `uf-init` and requires `dewey-index` to remain available.
Run it before deletion to establish the failure, then after deletion.

Replicate the current CI build, vet, race/coverage tests and Gaze
threshold gate (48 CRAPload, 45 GazeCRAPload), and run the configured
Go, Markdown, YAML, secret and whitespace checks without auto-fixes.
Use an isolated GOBIN for CI's Gaze dependency; do not overwrite the
operator's working tools. Independent review follows the final diff.

Native OpenCode command discovery must not start this project's MCP
servers or local inference. If that cannot be established safely,
report the filesystem-contract test rather than claim a live TUI run.

## Documentation and Delivery

The planning record documents the harness-only removal. Dewey's public
CLI/MCP documentation and website behavior are unchanged. Leave the
primary checkout and its existing source/history untouched; publish
the feature branch to the existing operator fork when permitted.
No new fork, default-branch merge or release is part of this gap.

## Verification Results

**Readiness: BLOCKED by the existing Go formatting gate.** The
customizer removal passes its regression, but this branch is not
CI-clean or ready for a pull request. No gate was relaxed.

Checks ran with Go 1.25.10 against the unchanged production source
and new `project_commands_test.go`:

| Check | Result |
| --- | --- |
| Focused regression before deletion | Expected failure: retired customizer still advertised |
| Focused regression after deletion, `-race -count=1` | PASS |
| `go mod download`, `go build -v ./...`, `go vet ./...` | PASS |
| `go test -race -v -count=1 -coverprofile=coverage.out ./...` | PASS; 79.2% aggregate statement coverage |
| Gaze v1.9.0 threshold gate, limits 48/45 | PASS; CRAPload 40, GazeCRAPload 35 |
| Pinned golangci-lint 2.7.2, `run ./...` | FAIL; three existing gofmt findings |
| Pinned formatter diff on the new regression test | PASS; no diff |
| markdownlint-cli 0.47.0 on this change's Markdown | PASS; supplemental because the configured gate excludes `specs/` |
| yamllint 1.37.1, strict, on this change's metadata | PASS |
| Gitleaks 8.30.0, pinned default config, current-file scan | PASS; no leaks |
| OpenSpec strict validation and `git diff --check` | PASS |

Gaze was installed using CI's `@latest` command into a task-local
GOBIN; it resolved to v1.9.0. No installed working tool was replaced.
The gate ran without AI calls:

```bash
gaze crap --format=json --coverprofile=coverage.out \
  --max-crapload=48 --max-gaze-crapload=45 ./...
```

Tests used an isolated Git identity rather than changing global Git
configuration. Markdown and YAML checks initially identified missing
top-level headings and a YAML document marker in the new planning
records; those formatting issues were corrected and checks passed.

### Blocking Baseline Findings

The pinned Go linter reports gofmt failures at `cli.go:240`,
`curate/curate_test.go:55` and `main_test.go:1424`. These three files
and `.golangci.yml` are byte-identical to the baseline. SHA-256
comparisons establish that the findings do not come from this change:

| File | Baseline and current SHA-256 |
| --- | --- |
| `cli.go` | `c9f2b36522a638563cd25159aefaf11dadbaab604b1ebeb120ebdd909453a534` |
| `curate/curate_test.go` | `f5b6a5be7ba772c2a046e926b73cc79e31abdf77a34d4f571bfacdadb4a460d1` |
| `main_test.go` | `c5d9c27bb12299dcd287b225276ac95826136226cb8dc4ac79d846b498352287` |
| `.golangci.yml` | `14c67e0b80241525638a4e5f56c7ac1883ed5234c10e541622475d3aed49f5eb` |

A non-mutating formatter diff confirms indentation/alignment fixes
are required in those existing files. They are outside this removal's
source scope and remain unchanged. Resolve them as a separate
formatting-only change, then rerun the required lint gate before
declaring this branch ready. Do not disable the formatter or narrow
the protected project-wide gate to make this change pass.

### Independent Review and Verification Limits

Independent architecture, security, intent and test-quality reviews
approved the bounded source change without findings. Their approvals
do not override the lint failure. A separate cold-reader review passed
the receipt's scope and preservation checks.

The documented command-directory contract was tested directly; a live
OpenCode TUI/API discovery run was not performed because isolation
from project MCP startup and local inference was not established.
The individual pinned lint tools were run without auto-fixes, not the
full MegaLinter container or hosted action. No hosted CI result is
claimed for the implementation snapshot.

## Preservation and Remaining Scope

The baseline is `f1dbda22c39041b18b319f4c3959b64cd3a8dfdf`.
Only `.opencode/commands/uf-init.md` is removed among existing files.
The regression and this change's planning records are additions.
The installed Dewey, Gaze and Replicator binaries and shared Replicator
database match their pre-change SHA-256 checksums. No data migration,
module rename, legal attribution edit or image substitution occurs.

This is a branch-scoped retirement, not complete UF removal. The
primary checkout retains its baseline until separate integration.
Other deployed UF commands, personas, convention-pack wiring, schema
selection and organization inheritance remain. The next acceptance
gap is to resolve the three baseline formatting findings separately
and rerun lint. The next UF source gap after that is to remove the
inheritance from Dewey's instructions while retaining the substantive
local conventions and independent review/CI checks.
Other projects and the provider-blocked account partition are not
addressed by this change.

## Sources

- [OpenCode V2 command discovery](https://opencode.ai/v2/docs/commands)
- [OpenCode V2 migration guide](https://opencode.ai/v2/docs/migrate-v1)
- [OpenCode V2 command API](https://opencode.ai/v2/docs/api)
- [MegaLinter's pinned linter descriptors](https://github.com/oxsecurity/megalinter/tree/42bb470545e359597e7f12156947c436e4e3fb9a/docs/descriptors)
- [Dewey CI workflow](../../../.github/workflows/ci.yml)
- [Dewey MegaLinter configuration](../../../.mega-linter.yml)
