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
