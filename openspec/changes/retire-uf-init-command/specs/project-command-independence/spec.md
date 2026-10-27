# Project Command Independence

## ADDED Requirements

### Requirement: Retire the UF customizer from project commands

Dewey's project-local OpenCode command catalog MUST NOT advertise
the `uf-init` customizer after the UF installation is removed.

#### Scenario: Enumerating project commands

- **WHEN** command Markdown files are discovered from either supported
  project command directory
- **THEN** no command named `uf-init` is available
- **AND** Dewey's `dewey-index` command remains available

### Requirement: Preserve the independent tool and local standards

The removal MUST preserve Dewey's source, data, standalone tools,
module and legal provenance, deployed conventions, CI thresholds
and meaningful review requirements.

#### Scenario: Reviewing the bounded removal

- **WHEN** the implementation diff is compared with the baseline
- **THEN** the only removed existing file is the UF customizer
- **AND** no existing production, convention, review, CI or data file
  is modified
