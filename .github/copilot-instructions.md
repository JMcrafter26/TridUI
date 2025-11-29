## TridUI — Copilot Instructions

This file gives concise, actionable guidance for AI coding agents working in this repository.

**Big Picture**:
- **Backend**: Go (Wails) — `main.go` embeds the frontend (`//go:embed all:frontend/dist`) and binds the `App` struct to the frontend via `Bind: []interface{}{app}`. Add exported methods on `App` to expose new IPC functions.
- **Frontend**: SvelteKit + TypeScript in `frontend/` — routes live under `frontend/src/routes` (`+page.svelte`, `+page.ts`). UI uses `wailsjs` generated bridge code in `frontend/wailsjs/` to call Go methods.
- **TrID engine**: Native Go implementation in `trid/` (`trid/trid.go`) — key types and funcs: `Analyzer`, `LoadDefinitions`, `LoadDefinitionsFromBytes`, `AnalyzeFile`, `AnalyzeBytes`, `GetBestMatch`.

**Where to look for concrete examples**:
- Wails bind and startup: `main.go` (see `Bind: []interface{}{app}` and `startup` method in `app.go`).
- TRD parsing and analyzer API: `trid/trid.go` (pattern matching, `Analyzer.AnalyzeFile`, `Analyzer.LoadDefinitions`).
- Frontend models used by Svelte: `frontend/wailsjs/go/models.ts` (maps Go results to TypeScript classes such as `TridScanResult`).
- Frontend store patterns: `frontend/src/lib/stores/*` (e.g. `updateStore.ts` uses `writable<main.UpdateInfo | null>`).

**Developer workflows / useful commands**:
- Run both backend and frontend (dev): from repo root run `wails dev` (Wails will watch frontend via `wails.json` `frontend:dev:watcher`).
- Frontend only (fast iteration):
  - `cd frontend`
  - `pnpm install`
  - `pnpm dev`
- Quick full build (platform):
  - Windows: `.\\build\\build-windows.bat`
  - macOS: `./build/build-darwin.sh`
  - Linux: `./build/build-linux.sh`
- Manual Wails build examples (from `build/README.md`):
  - `wails build -platform windows/amd64`
  - `wails build -upx` (with compression)
- Prerequisites: Go >= 1.22, Node 20+, pnpm 10+, Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`).

**Patterns & conventions specific to this repo**:
- IPC exposure: any exported method on the `App` struct gets bound to the frontend (see `main.go` `Bind`). To add new functionality, implement an exported receiver method on `App` (keep methods concise and return JSON-able types or the types present in `frontend/wailsjs/go/models.ts`).
- Frontend/Backend types: the TypeScript classes in `frontend/wailsjs/go/models.ts` are generated/expected shapes — keep Go return types compatible (field names and casing matter for JSON mapping).
- Asset embedding: production frontend is embedded using Go `embed` from `frontend/dist`. During dev use `wails dev` or `pnpm dev` for the frontend.
- TrID definitions: the app can download `triddefs.trd` automatically; the manual location is platform-specific (Windows `%APPDATA%\\TridUI\\triddefs.trd`). When updating the analyzer behavior, inspect `Analyzer.LoadDefinitions*` and consumers.

**Integration points**:
- `wails.json` contains `frontend:install`, `frontend:build`, and `frontend:dev:watcher` hooks — editing these affects dev/build flow.
- `build/` scripts orchestrate cross-platform packaging and call `wails build`. Use them to reproduce CI artifacts.
- Generated bridge code in `frontend/wailsjs` is the contract between Go and Svelte. Regenerate or inspect when changing Go exports or models.

**Small, high-value examples**:
- Expose a new scan function: add `func (a *App) ScanFile(path string) (main.TridScanResult, error)` in Go, then call `wails dev` — frontend can call `window.backend.Main.ScanFile(...)` via generated `wailsjs` helpers.
- Load local definitions in tests: call `Analyzer.LoadDefinitionsFromBytes(data)` to test parsing without filesystem dependencies.

**What to avoid / notes**:
- Do not change field names on Go models without updating/confirming `frontend/wailsjs/go/models.ts` mapping.
- UI expects certain shapes (see `TridScanResult` and `TridFileTypeResult` in `frontend/wailsjs/go/models.ts`) — maintain backward-compatible changes.

If anything here is unclear or you'd like me to expand specific areas (Wails IPC patterns, frontend routing conventions, or the TRD parser internals), tell me which part to iterate on.
