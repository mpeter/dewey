## ✅⚠️[MegaLinter](https://megalinter.io/9.3.0) analysis: Success with warnings



| Descriptor  |                                   Linter                                    |Files|Fixed|Errors|Warnings|Elapsed time|
|-------------|-----------------------------------------------------------------------------|----:|----:|-----:|-------:|-----------:|
|✅ GO        |[golangci-lint](https://megalinter.io/9.3.0/descriptors/go_golangci_lint)    |  yes|  yes|    no|      no|      52.65s|
|⚠️ MARKDOWN  |[markdownlint](https://megalinter.io/9.3.0/descriptors/markdown_markdownlint)|    7|    1|    58|       0|       1.75s|
|✅ REPOSITORY|[gitleaks](https://megalinter.io/9.3.0/descriptors/repository_gitleaks)      |  yes|     |    no|      no|       2.83s|
|✅ REPOSITORY|[git_diff](https://megalinter.io/9.3.0/descriptors/repository_git_diff)      |  yes|     |    no|      no|       0.31s|
|✅ YAML      |[yamllint](https://megalinter.io/9.3.0/descriptors/yaml_yamllint)            |   19|     |     0|       0|       1.14s|

## Detailed Issues

<details>
<summary>⚠️ MARKDOWN / markdownlint - 58 errors</summary>

```
AGENTS.md:5:401 error MD013/line-length Line length [Expected: 400; Actual: 701]
AGENTS.md:79:401 error MD013/line-length Line length [Expected: 400; Actual: 461]
AGENTS.md:139:11 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:139:21 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:139:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:139:11 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:236:8 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:236:16 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:236:30 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:236:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:236:8 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:236:16 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:245:11 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:245:25 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:245:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:245:11 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:265:8 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:265:22 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:265:40 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:265:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:265:8 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:265:22 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:306:401 error MD013/line-length Line length [Expected: 400; Actual: 410]
AGENTS.md:368:8 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:368:17 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:368:31 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:368:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:368:8 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:368:17 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:423:401 error MD013/line-length Line length [Expected: 400; Actual: 569]
AGENTS.md:449:15 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:449:27 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:449:41 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:449:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:449:15 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:449:27 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:516:401 error MD013/line-length Line length [Expected: 400; Actual: 539]
AGENTS.md:521:8 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:521:18 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:521:33 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:521:42 error MD060/table-column-style Table column style [Table pipe is missing space to the left for style "compact"]
AGENTS.md:521:1 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:521:8 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:521:18 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:521:33 error MD060/table-column-style Table column style [Table pipe is missing space to the right for style "compact"]
AGENTS.md:564:401 error MD013/line-length Line length [Expected: 400; Actual: 424]
AGENTS.md:573:401 error MD013/line-length Line length [Expected: 400; Actual: 412]
openspec/changes/adopt-openspec-only-workflow/design.md:1 error MD041/first-line-heading/first-line-h1 First line in a file should be a top-level heading [Context: "## Context"]
openspec/changes/adopt-openspec-only-workflow/proposal.md:1 error MD041/first-line-heading/first-line-h1 First line in a file should be a top-level heading [Context: "## Why"]
openspec/changes/adopt-openspec-only-workflow/tasks.md:1 error MD041/first-line-heading/first-line-h1 First line in a file should be a top-level heading [Context: "Implementation status: migrati..."]
openspec/changes/uf-dewey-workspace/tasks.md:1 error MD041/first-line-heading/first-line-h1 First line in a file should be a top-level heading [Context: "## 1. Introduce Workspace Cons..."]
openspec/changes/upgrade-embedding-model-r2/tasks.md:13 error MD041/first-line-heading/first-line-h1 First line in a file should be a top-level heading [Context: "## 1. Add configurable chunk l..."]
openspec/changes/upgrade-embedding-model-r2/tasks.md:15:401 error MD013/line-length Line length [Expected: 400; Actual: 555]
openspec/changes/upgrade-embedding-model-r2/tasks.md:18:401 error MD013/line-length Line length [Expected: 400; Actual: 603]
openspec/changes/upgrade-embedding-model-r2/tasks.md:26:401 error MD013/line-length Line length [Expected: 400; Actual: 905]
openspec/changes/upgrade-embedding-model-r2/tasks.md:29:401 error MD013/line-length Line length [Expected: 400; Actual: 442]
openspec/changes/upgrade-embedding-model-r2/tasks.md:54:401 error MD013/line-length Line length [Expected: 400; Actual: 411]
openspec/changes/upgrade-embedding-model-r2/tasks.md:55:401 error MD013/line-length Line length [Expected: 400; Actual: 452]
```

</details>

See detailed reports in MegaLinter artifacts
_Set `VALIDATE_ALL_CODEBASE: true` in mega-linter.yml to validate all sources, not only the diff_

[![MegaLinter is graciously provided by OX Security](https://raw.githubusercontent.com/oxsecurity/megalinter/main/docs/assets/images/ox-banner.png)](https://www.ox.security/?ref=megalinter)