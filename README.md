# GovulnCheck Test Repository

This repository is designed to test the `govulncheck-pr` GitHub Actions workflow.

## Purpose

- Tests the automatic vulnerability scanning using govulncheck
- Verifies PR creation when vulnerabilities are found
- Uses dependencies that may have known vulnerabilities for testing

## Dependencies

- `github.com/gin-gonic/gin` - Web framework
- `github.com/gorilla/mux` - HTTP router
- `gopkg.in/yaml.v2` - YAML parsing (potentially vulnerable)

## Usage

1. Push to main branch or create a PR
2. The workflow will automatically run govulncheck
3. If vulnerabilities are found, a PR will be created with the report

## Manual Testing

```bash
go run main.go
```

## Setup

To initialize dependencies:
```bash
go mod tidy
```

To test govulncheck locally:
```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```
