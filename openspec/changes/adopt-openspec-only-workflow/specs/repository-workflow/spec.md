## ADDED Requirements

### Requirement: Single requirements authority

The repository MUST use OpenSpec as its only requirements pipeline and MUST keep
its constitution at `openspec/constitution.md`. Repository agents MUST refer to
OpenSpec artifacts and available commands, not retired workflow helpers.

#### Scenario: Contributor starts a change
- **GIVEN** a contributor reads the repository instructions
- **WHEN** they start a requirements-bearing change
- **THEN** they use `openspec/changes/<change>/` and an `opsx/<change>` branch

### Requirement: Preserve historical evidence

Numbered strategic specifications MUST move to dated OpenSpec archives with their
original requirements and task states preserved. Open tasks MUST remain open
unless independently proven complete, with unverified work identified explicitly.

#### Scenario: Feature is merged but an individual task is not proven
- **GIVEN** a feature has a merged PR and an unchecked historical task
- **WHEN** its artifacts are archived
- **THEN** the task stays unchecked with an unverified annotation and PR provenance
