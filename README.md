# ghx


AI-powered GitHub workflow assistant built as a Go TUI.

## Current Slice

- Bubble Tea app shell with a menu-driven home screen.
- Git status screen backed by `go-git`.
- AI provider/config scaffolding with Ollama as the local default.
- Prompt templates for commit messages, PR descriptions, and code explanations.

## Run

```bash
go mod tidy
go run .
```

Use arrow keys or `j`/`k` to navigate, `enter` to open a workflow, `b` to go back, and `q` to quit.
