# Retire the UF Customizer

## Why

The operator is removing the Unbound Force installation. Dewey still
advertises a project-local `/uf-init` customizer that requires that
installation and directs users to run `uf init` and `uf setup`.
This command is not part of Dewey's CLI or MCP interface.

## What Changes

- Retire `.opencode/commands/uf-init.md` from Dewey's command catalog.
- Add a regression check using the two supported project command
  directories, with a retained Dewey command as a positive control.
- Preserve the already-deployed conventions, review instructions,
  OpenSpec/Speckit commands, Dewey implementation, data and provenance.

## Capabilities

### New Capabilities

- `project-command-independence`: the project command catalog does
  not offer the retired UF customizer.

## Impact

Only a development-harness command is retired. No Dewey executable,
MCP tool, module path, dependency, data format or container image
changes. Other UF harness dependencies remain outside this gap.

## Constitution Alignment

- Composability First: PASS; removes a command requiring another tool.
- Autonomous Collaboration: PASS; no Dewey MCP contract changes.
- Observable Quality: PASS; record the removed catalog entry and checks.
- Testability: PASS; test real project command files without a service.
