# Go Vulnerability Scan with PR Creation

> **Note**: This workflow currently uses the feature branch `phipps/govulncheck-pr-tool` for testing. Once validated, the action will be available on the main branch as `@v1`.

This workflow runs [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) to scan for Go vulnerabilities and automatically creates or updates a Pull Request if any reachable vulnerabilities are found. It uses the `govulncheck-pr` action from the `csgda-platform-actions` repository to provide intelligent PR management with non-duplicating branch and PR handling.

## Overview

This workflow is designed to:
- Automatically scan Go projects for known vulnerabilities
- Create a single, persistent Pull Request when vulnerabilities are detected
- Update the same PR with new scan results over time
- Provide detailed vulnerability reports and remediation guidance
- Integrate seamlessly with existing CI/CD workflows

## Key Features

- **Smart PR Management**: Uses a consistent branch name (`govulncheck-fixes`) and updates the same PR instead of creating duplicates
- **Detailed Reporting**: Generates comprehensive vulnerability reports with affected modules and remediation steps
- **Automatic Updates**: Updates existing PRs with new scan results and adds comments about changes
- **Fallback to Issues**: Creates an issue if PR creation fails (e.g., due to permissions)
- **Flexible Scheduling**: Can be run on schedule, manual trigger, or code changes
- **Go Project Detection**: Automatically detects if the repository contains Go files

## Workflow Triggers

### Scheduled Scans
```yaml
schedule:
  - cron: '0 2 * * 1'  # Weekly scan on Mondays at 2 AM UTC
```

### Manual Execution
```yaml
workflow_dispatch:
  inputs:
    force-scan:
      description: 'Force scan even if no code changes detected'
      required: false
      type: boolean
      default: false
    base-branch:
      description: 'Base branch for PR creation'
      required: false
      type: string
      default: 'main'
```

### Automatic on Code Changes
```yaml
push:
  branches: [main]
  paths:
    - 'go.mod'
    - 'go.sum'
    - '**/*.go'
pull_request:
  branches: [main]
  paths:
    - 'go.mod'
    - 'go.sum'
    - '**/*.go'
```

## Inputs

| Name        | Description                                    | Required | Default | Type    |
|-------------|------------------------------------------------|----------|---------|---------|
| force-scan  | Force scan even if no code changes detected   | `false`  | `false` | boolean |
| base-branch | Base branch for PR creation                    | `false`  | `main`  | string  |

## Required Permissions

The workflow requires the following permissions:

```yaml
permissions:
  contents: write      # Required to create branches and commit files
  pull-requests: write # Required to create and update PRs
  issues: write        # Required for fallback issue creation
```

## How It Works

1. **Repository Check**: Automatically detects if the repository contains Go files
2. **Vulnerability Scan**: Runs `govulncheck` to scan for known vulnerabilities
3. **PR Management**: 
   - Checks for existing PRs with govulncheck results
   - Updates existing PR/branch if found
   - Creates new PR/branch if none exists
4. **Detailed Reporting**: Generates comprehensive vulnerability reports
5. **Continuous Updates**: Each scan updates the same PR with new results

## Example Usage

### Basic Setup

Copy this workflow to `.github/workflows/govulncheck-pr.yml` in your repository:

```yaml
name: 'Go Vulnerability Scan with PR Creation'

on:
  schedule:
    - cron: '0 2 * * 1'  # Weekly scan on Mondays at 2 AM UTC
  workflow_dispatch:
  push:
    branches: [main]
    paths:
      - 'go.mod'
      - 'go.sum'
      - '**/*.go'

jobs:
  govulncheck-pr:
    name: 'Go Vulnerability Scan and PR Creation'
    runs-on: [external-k8s-v2]
    permissions:
      contents: write
      pull-requests: write
      issues: write
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Check if Go files exist
        id: check-go
        run: |
          if find . -name "*.go" -type f | grep -q .; then
            echo "go-files=true" >> $GITHUB_OUTPUT
          else
            echo "go-files=false" >> $GITHUB_OUTPUT
          fi

      - name: Run govulncheck and create PR if vulnerabilities found
        if: steps.check-go.outputs.go-files == 'true'
        uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          base-branch: main
```

### Advanced Configuration

For custom base branch and forced scanning:

```yaml
name: 'Go Vulnerability Scan - Development'

on:
  workflow_dispatch:
    inputs:
      force-scan:
        description: 'Force scan even if no code changes'
        type: boolean
        default: false
      base-branch:
        description: 'Base branch for PR'
        type: string
        default: 'develop'
  push:
    branches: [develop]

jobs:
  govulncheck-pr:
    name: 'Go Vulnerability Scan and PR Creation'
    runs-on: [external-k8s-v2]
    permissions:
      contents: write
      pull-requests: write
      issues: write
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Run govulncheck and create PR if vulnerabilities found
        uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          base-branch: ${{ inputs.base-branch || 'develop' }}
```

## Branch and PR Strategy

The workflow uses a **single, persistent branch and PR** approach:

- **Branch Name**: `govulncheck-fixes` (consistent across all scans)
- **PR Title**: "🚨 Security: govulncheck found X reachable vulnerabilities"
- **Updates**: Each scan updates the same PR with new results and adds comments
- **No Duplicates**: Prevents creation of multiple PRs for the same purpose

## Vulnerability Report Structure

When vulnerabilities are found, the action creates a comprehensive report including:

- **Vulnerability Summary**: List of CVEs and affected modules
- **Scan Details**: Date, scanner version, and database information
- **Next Steps**: Clear remediation instructions
- **Historical Context**: Comments showing scan progression over time

## Integration with Existing Workflows

This workflow can be easily integrated with existing CI/CD pipelines:

### With Code Quality Checks

```yaml
name: 'CI Pipeline'

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    name: 'Tests'
    runs-on: [external-k8s-v2]
    steps:
      - uses: actions/checkout@v4
      - name: Run tests
        run: go test ./...

  vulnerability-scan:
    name: 'Vulnerability Scan'
    needs: test
    uses: ./.github/workflows/govulncheck-pr.yml
    secrets: inherit
```

### With Multiple Go Modules

For repositories with multiple Go modules:

```yaml
strategy:
  matrix:
    module: [api, worker, cli]
    
steps:
  - name: Run govulncheck for module
    uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
    with:
      github-token: ${{ secrets.GITHUB_TOKEN }}
      base-branch: main
    env:
      WORKING_DIRECTORY: ./${{ matrix.module }}
```

## Troubleshooting

### Common Issues

1. **Permission Denied**: Ensure the workflow has required permissions (`contents: write`, `pull-requests: write`)
2. **No Vulnerabilities Found**: The action only creates PRs when vulnerabilities are detected
3. **Branch Conflicts**: The action handles git conflicts by removing conflicting files before checkout
4. **No Go Files**: The workflow automatically skips scanning for non-Go repositories

### Debug Information

The workflow provides detailed logging including:
- Go file detection results
- Vulnerability scan output
- PR creation and update operations
- Comprehensive workflow summary

## Security Considerations

- Uses GitHub's actions bot account for git operations
- Handles authentication via GitHub tokens
- Validates inputs to prevent injection attacks
- Provides fallback mechanisms for permission issues

## Dependencies

The underlying action automatically installs:
- Go 1.21
- govulncheck (latest version)
- jq (JSON processor)
- gh (GitHub CLI)

## Best Practices

1. **Schedule Regular Scans**: Use cron schedule for regular vulnerability monitoring
2. **Monitor PR Updates**: Set up notifications for the `govulncheck-fixes` branch
3. **Act on Vulnerabilities**: Prioritize updating dependencies when vulnerabilities are found
4. **Review Reports**: Examine the detailed reports in `.govulncheck/latest.json`
5. **Test Fixes**: Run `govulncheck ./...` locally to verify fixes before closing PRs

## Reference

- **Action Repository**: [bayer-int/csgda-platform-actions](https://github.com/bayer-int/csgda-platform-actions)
- **govulncheck Documentation**: https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck
- **Action Source**: [govulncheck-pr](https://github.com/bayer-int/csgda-platform-actions/actions/govulncheck-pr)

## Changelog

### v1.0.0
- Initial workflow implementation
- Integration with `govulncheck-pr` action
- Automatic Go project detection
- Comprehensive vulnerability reporting
- Smart PR management with single branch strategy
