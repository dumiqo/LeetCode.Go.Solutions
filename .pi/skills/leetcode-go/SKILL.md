---
name: leetcode-go
description: Solve a LeetCode problem in the LeetCode.Go.Solutions Go repository. Use when asked to solve, implement, add, or test a LeetCode problem here, especially from a leetcode.com URL or a daily-question link. Encodes the required folder layout, file naming, package, function-name, test-style and verification conventions so solutions fit the repo on the first try.
---

# Solving LeetCode Problems in LeetCode.Go.Solutions

This repository is a Go 1.22.5 module (`github.com/dumiqo/LeetCode.Go.Solutions`) that
collects LeetCode solutions, one folder per problem. `testify` is a declared dependency.

## Mandatory: go-feature workflow + go-development skill

Solve every problem with the **go-feature workflow** and the **go-development** house
style. Read these global assets in full and follow them:

- Workflow: `C:\Users\Костя\.pi\agent\prompts\go-feature.md`
- Style: `C:\Users\Костя\.pi\agent\skills\go-development\SKILL.md`
- Role agents: `C:\Users\Костя\.pi\agent\agents\` (`researcher`, `architect`,
  `go-developer`, `test-engineer`, `code-reviewer`)

Run the go-feature phases as **sequential subagents** (use the `pi-subagents` skill for
delegation): researcher (statement + repo reconnaissance), architect (implementation
plan), go-developer (implement), test-engineer (tests), code-reviewer (independent
review of the diff). If review finds High/Critical issues, loop back to go-developer,
then re-run test-engineer and code-reviewer. Never parallelize implementation and
review. For a tiny single-file problem you may inline a phase when spawning a subagent
is impractical, but all phases still run in order.

**Where the workflow conflicts with this repository's conventions below, the repo wins**
(`package leetcode`, plain `testing` over testify, per-folder verification — CRLF line
endings and known pre-existing failures make repo-wide `gofmt -l` / `go test ./...`
runs noisy — and re-declared local types instead of shared packages). The go-feature
phases and `go-development` govern everything else.

## Non-negotiable conventions

Every solution lives in its **own directory** and declares **`package leetcode`**
(all 59 `.go` files use it). Never use `package main`.

### Directory name

Use the LeetCode problem number and title, title-cased with hyphens:

```
<number>-<Title-With-Hyphens>/
```

Existing folders are inconsistent, so **match the dominant style**:
`840-Magic-Squares-In-Grid/`, `860-Lemonade-Change/`, `1140-Stone-Game-II/`.
Reuse an existing folder if the problem is already present — never duplicate it.
(Some older folders are all-lowercase, e.g. `3568-minimum-moves-to-clean-the-classroom/`;
that is legacy, do not imitate it for new problems.)

### Files inside the folder

Preferred (majority) naming — file names repeat the title with spaces:

```
<number>-<Title-With-Hyphens>/
├── solution.go          # solution
├── solution_test.go     # tests
└── README.md                       # problem statement
```

For the statement file use **`README.md`** (24 folders) over `readme.md` (5 folders).
On Windows the filesystem is case-insensitive, so `ls */readme.md` will *appear* to match
every folder — verify casing before assuming. The repository root contains BOTH
`README.md` and `readme.md`; the lower-case one is a personal Russian-language notes file
— do **not** edit either root file unless explicitly asked.

## Solution file rules

- `package leetcode`, one exported-free function named after the problem in camelCase,
  matching LeetCode's own function name. Examples already in the repo:
  `lemonadeChange`, `spiralMatrixIII`, `isRectangleOverlap`, `countSeniors`,
  `kthDistinct`, `nodesBetweenCriticalPoints`, `flipEquiv`, `minMoves`, `numberToWords`.
- Use the exact LeetCode signature and the exact return type (including `int64` when
  the problem returns one, e.g. `kthLargestLevelSum` returns `int64`).
- Helper functions are unexported and low-risk names (`dfs`, `parse`, `getTens`,
  `bruteForceCountCommas`). The repo's own helpers are intentionally named to avoid
  surprises: `1140-Stone-Game-II` and `877-Stone-Game` use `minOf`/`maxOf` while
  `650-2-Keys-Keyboard` defines a plain `func min`. Since Go 1.21 `min`/`max` are
  builtins, so a shadowing local helper is legal but confusing — prefer `minOf`/`maxOf`,
  and `grep` the folder for a name before redefining it in the same package.
- **Shared types are declared per-folder, not in a common package.** If the problem uses
  a linked list or tree, re-declare the type locally in that folder's solution file,
  exactly like `951-Flip-Equivalent-Binary-Trees` / `2583-Kth-Largest-Sum-in-a-Binary-Tree`:

  ```go
  type TreeNode struct {
      Val   int
      Left  *TreeNode
      Right *TreeNode
  }
  ```

  Do not import a shared `TreeNode` from another folder. Cross-folder imports are not used.

## Test file rules

- `package leetcode`, in the same folder as the solution.
- Table-driven with `t.Run` subtests is the house style:

  ```go
  package leetcode

  import (
      "testing"
  )

  func TestIsRectangleOverlap(t *testing.T) {
      var tests = []struct {
          name string
          rec1 []int
          rec2 []int
          want bool
      }{
          {"Overlapping rectangles", []int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
          {"Touch at an edge", []int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
      }
      // The execution loop
      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              ans := isRectangleOverlap(tt.rec1, tt.rec2)
              if ans != tt.want {
                  t.Errorf("got wrong result %v, want %v", ans, tt.want)
              }
          })
      }
  }
  ```

- Two styles are in use: plain `testing` (18 files) and `testify/assert` (11 files).
  Prefer **plain `testing`** — it is the default and has no import risk. Use testify only
  if the file already imports it.
- Test function names are `Test<PascalCaseProblemName>` (e.g. `TestIsRectangleOverlap`,
  `TestNumberToWords`). One legacy typo exists (`TestДemonadeChange`) — do not copy it.
- Always include the LeetCode examples as cases **plus** edge cases: empty/single input,
  boundary values, duplicates, negatives, and the largest allowed constraints.

## Workflow

1. **Check for an existing folder first** (`ls -d */`, or list a specific candidate).
   If the problem already exists, extend it instead of creating a new folder.
2. Read a nearby solution folder of the same shape (e.g. `860-Lemonade-Change/`) to
   confirm the current local style before writing.
3. Create the folder and the three files (`<Title>.go`, `<Title>_test.go`, `README.md`).
4. Transcribe the problem statement into `README.md` (examples + constraints). Use the
   text from the user's link; `web_search` may be unavailable depending on the configured
   model, in which case ask the user to paste the statement rather than guessing it.
5. Implement the solution.
6. Verify (below) and fix until green.

## Verification (required)

Run from the repository root and report the real output:

```bash
go test -count=1 ./<folder>/    # target package must pass
go vet ./<folder>/              # must be silent
gofmt -l ./<folder>/            # must print nothing
```

Then `go test -count=1 ./...` for a repo-wide sanity check.

### Known pre-existing failures — not your regression

`go test ./...` currently fails in folders unrelated to any new work:

- `145-Binary-Tree-Postorder-Traversal` — build error (`tt.input` used as `*TreeNode`)
- `624-Maximum-Distance-in-Arrays` — test logic failure

Do not "fix" these unless asked; just state that they predate the change and prove it by
showing that only your folders are modified (`git status --short`).

### CRLF caveat

Most committed files use CRLF line endings, so a repo-wide `gofmt -l .` flags many
untouched files. Only require `gofmt -l` to be clean for the folder you created; do not
run a repo-wide reformat.

## Report format

When done, summarize: the algorithm and its complexity, the files created (full relative
paths), the exact verification command output, and any residual risk or pre-existing
failures. Keep it concise.
