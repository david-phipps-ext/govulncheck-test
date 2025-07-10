# Govulncheck PR Workflow Template

> **Note**: These templates currently use the feature branch `phipps/govulncheck-pr-tool` for testing. Once the action is merged to main, update the action reference to `@v1`.

This template provides ready-to-use workflow configurations for the `govulncheck-pr` action from the `csgda-platform-actions` repository.

## Quick Start

Copy one of the workflow examples below to `.github/workflows/govulncheck-pr.yml` in your repository.

### Option 1: Simple Weekly Scan

```yaml
name: 'govulncheck-pr'

on:
  schedule:
    - cron: '0 2 * * 1'  # Weekly on Mondays at 2 AM UTC
  workflow_dispatch:

jobs:
  govulncheck-pr:
    name: 'govulncheck-pr'
    runs-on: [external-k8s-v2]
    permissions:
      contents: write
      pull-requests: write
      issues: write
    steps:
      - uses: actions/checkout@v4
      - uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          base-branch: main
```

### Option 2: Triggered on Go Code Changes

```yaml
name: 'govulncheck-pr'

on:
  schedule:
    - cron: '0 2 * * 1'
  workflow_dispatch:
  push:
    branches: [main]
    paths:
      - 'go.mod'
      - 'go.sum'
      - '**/*.go'

jobs:
  govulncheck-pr:
    name: 'govulncheck-pr'
    runs-on: [external-k8s-v2]
    permissions:
      contents: write
      pull-requests: write
      issues: write
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          base-branch: main
```

### Option 3: With Manual Configuration

```yaml
name: 'govulncheck-pr'

on:
  schedule:
    - cron: '0 2 * * 1'
  workflow_dispatch:
    inputs:
      base-branch:
        description: 'Base branch for PR creation'
        required: false
        default: 'main'
  push:
    branches: [main]
    paths: ['go.mod', 'go.sum', '**/*.go']

jobs:
  govulncheck-pr:
    name: 'govulncheck-pr'
    runs-on: [external-k8s-v2]
    permissions:
      contents: write
      pull-requests: write
      issues: write
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          base-branch: ${{ inputs.base-branch || 'main' }}
```

## Required Permissions

All workflows require these permissions:

```yaml
permissions:
  contents: write      # Create branches and commit files
  pull-requests: write # Create and update PRs  
  issues: write        # Fallback issue creation
```

## How It Works

1. Scans your Go code for vulnerabilities using `govulncheck`
2. If vulnerabilities are found, creates or updates a PR on the `govulncheck-fixes` branch
3. Provides detailed vulnerability reports in the PR description
4. Updates the same PR with new scan results over time (no duplicate PRs)

## What You Get

- **Single PR**: Uses one consistent branch (`govulncheck-fixes`) for all vulnerability reports
- **Detailed Reports**: Comprehensive vulnerability information with remediation steps
- **Smart Updates**: Updates existing PRs instead of creating duplicates
- **Automatic Management**: Handles git operations, branch creation, and PR updates automatically

## Next Steps

1. Copy one of the workflow examples above
2. Save it as `.github/workflows/govulncheck-pr.yml` in your repository
3. Commit and push the workflow file
4. The workflow will run according to your schedule or triggers
5. Monitor for PRs titled "🚨 Security: govulncheck found X reachable vulnerabilities"

## Additional Configuration

### For Development Branches

To scan development branches:

```yaml
on:
  push:
    branches: [main, develop]  # Add your development branches
```

### For Different Runners

If not using `external-k8s-v2`:

```yaml
runs-on: ubuntu-latest  # or your preferred runner
```

### For Organizations with Different Secrets

If using a different GitHub token:

```yaml
with:
  github-token: ${{ secrets.YOUR_CUSTOM_TOKEN }}
```
