## Why

OpenSpec must be the single requirements authority. The previous workflow leaves
two competing artifact trees and points repository agents at retired commands.

## What Changes

- Preserve numbered strategic specifications in dated OpenSpec archives, retaining
  task states and appending merged-PR provenance and unverified-task annotations.
- Move the constitution to `openspec/constitution.md` and update repository
  instructions, schema context, and references to use OpenSpec only.
- Remove the previous workflow's repository commands, skills, scripts and templates.
- Repair canonical specification structure without discarding requirements.

## Capabilities

### New Capabilities
- `repository-workflow`: One requirements authority and evidence-preserving archives.

### Modified Capabilities
None. Canonical specification repairs are structural only.

### Removed Capabilities
None. The retired repository helpers are not Dewey product capabilities.

## Impact

Repository documentation and agent assets only. `slash_commands.go` ships only
Dewey-specific commands and contains no previous-workflow assets. No production
code, downstream installation, API, or version changes are required. Existing
user-installed files are untouched because no installer behavior changes.

## Constitution Alignment

### I. Autonomous Collaboration
**Assessment**: PASS. OpenSpec artifacts remain the shared communication contract.

### II. Composability First
**Assessment**: PASS. Dewey's standalone runtime and installation are unchanged.

### III. Observable Quality
**Assessment**: PASS. Archives retain original task states and explicit provenance.

### IV. Testability
**Assessment**: PASS. Validate artifact structure, exact inventory and task-state
preservation, then run the repository build, vet, race tests and quality gates.
