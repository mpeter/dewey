## Context

The repository contains 15 numbered strategic-spec directories, repository-only
workflow helpers, three canonical specs, and existing OpenSpec changes. Migration
follows gaze-py PRs #93 and #94, without modifying unrelated product behavior.

## Goals / Non-Goals

### Goals
- One authoritative requirements tree, with usable OpenSpec-only instructions.
- Preserve requirements, scenarios and historical task completion states.

### Non-Goals
- Change Dewey commands, rewrite Git history, delete remote branches, or release.
- Repair unrelated pre-existing active-change validation defects.

## Decisions

Archive numbered specs using the merged feature PR's UTC merge date and original
feature name. Preserve every original artifact. Rewrite retired terminology
neutrally, update moved paths, append migration evidence, and annotate every open
task as unverified rather than inferring completion from a feature-level merge.

The constitution's workflow wording is updated under the explicit migration
request; its core principles and release rule remain unchanged. The release rule
requires MAJOR bumps for breaking MCP contracts or backend interfaces; neither
changes, and no product version is bumped.

No legacy assets are embedded or shipped by Dewey. Therefore a separate product
removal change, stale-file deletion behavior, and website publication are not
applicable. Existing installed files remain untouched.

Constitution alignment is assessed in the proposal. No dependencies or runtime
coupling are introduced; provenance and full regression verification preserve
Observable Quality and Testability.

## Risks / Trade-offs

Historical plans may describe behavior that changed later. They remain evidence,
not a second active pipeline. Existing OpenSpec changes have pre-existing layout
errors; report them without expanding this migration into unrelated repairs.
