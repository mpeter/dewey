# UF Customizer Retirement Tasks

## 1. Regression and Removal

- [ ] 1.1 Add the command-catalog regression and observe it fail
      against the existing `uf-init` entry.
- [ ] 1.2 Remove only `.opencode/commands/uf-init.md`; verify the
      regression passes and retained command assets are unchanged.

## 2. Verification and Delivery

- [ ] 2.1 Run the current CI-equivalent build, vet, race/coverage
      tests, Gaze thresholds and configured lint/secret checks.
- [ ] 2.2 Complete independent review and document limitations,
      remaining scope and preservation checks.
- [ ] 2.3 Commit the scoped change and publish to the existing
      operator fork if access permits, without a new fork or merge.

## Blocking Gate

All completion checkboxes remain open under the CI parity gate.
The red/green regression, scoped removal and independent reviews were
performed, but are not recorded as CI-complete tasks. Task 2.1 blocks
completion: pinned golangci-lint 2.7.2 reports three
pre-existing gofmt findings in unchanged `cli.go`,
`curate/curate_test.go` and `main_test.go`. Build, vet, race/coverage,
Gaze thresholds and the other scoped checks pass; independent source
reviews do not waive the failed lint gate. See `design.md` for results,
baseline hashes and verification limits. This change is not PR-ready.

A WIP snapshot may be committed and pushed to the existing feature
branch to preserve the work; that is not verified delivery under 2.3.
