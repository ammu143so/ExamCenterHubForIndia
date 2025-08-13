# ExamCenterHub (Indian Examination Center Assignment System)

A Go console application that assigns examination centers based on the nearest Indian cities (excluding the student's home city). Includes a basic flow and an advanced flow with preferences and seat capacities.

## Features
- Nearest city calculation using Haversine formula
- Multiple exam centers per city with seat capacity tracking
- Basic flow: quick nearest center suggestions
- Advanced flow: max distance, transport, accommodation preferences
- Predefined exams: JEE, NEET, UPSC, CAT, GATE, SSC, IBPS, IELTS

## Project Structure
```
ExamCenterHub/
├── go.mod
├── cmd/
│   └── examcenterhub/
│       └── main.go          # CLI entrypoint
└── internal/
    └── handler/
        ├── handler.go       # Business logic (basic + advanced)
        └── models.go        # Data models
```

## Requirements
- Go 1.21+

## Run
```bash
# From the project root
cd ExamCenterHub

# CLI (console)
go run ./cmd/examcenterhub

# Web UI (http://localhost:8080)
go run ./cmd/webui

# Or build binaries
#   CLI:    go build -o examcli ./cmd/examcenterhub
#   Web UI: go build -o examweb ./cmd/webui
```

## Usage
- Option 1: Basic Assignment – choose 1 and follow prompts
- Option 2: Advanced Assignment – choose 2 and set preferences
- Option 3: View Exam Types – choose 3 to list predefined exams
- Option 4: View Registration Summary – choose 4 to see session registrations

## Debugging
- CLI with Delve:
  - Install: `go install github.com/go-delve/delve/cmd/dlv@latest`
  - Run: `dlv debug ./cmd/examcenterhub`
- Web UI with Delve:
  - Run: `dlv debug ./cmd/webui --headless --listen=:2345 --api-version=2`
  - Attach from editor to `localhost:2345`
- VS Code (recommended):
  - Open folder `ExamCenterHub`
  - Add a launch config selecting program `cmd/examcenterhub` or `cmd/webui`
  - Press F5 to start debugging

## Notes
- Module path: `exam-center-assignment`
- Entry: `cmd/examcenterhub/main.go`
- Core logic: `internal/handler` (no external deps) 