# Govulncheck-PR Workflow Collection

> **⚠️ Development Notice**: These workflows currently reference the feature branch `phipps/govulncheck-pr-tool` for testing purposes. Once the action is validated and merged to the main branch, all references should be updated to `@v1`.

This directory contains various workflow configurations for the `govulncheck-pr` action from the `bayer-int/csgda-platform-actions` repository.

## 📁 Available Workflows

### 1. `govulncheck-pr-simple.yml` - Minimal Configuration
**Best for**: Quick setup, basic vulnerability scanning
```yaml
# Minimal weekly scan with manual trigger
# Perfect for getting started quickly
```

**Features**:
- ✅ Weekly scheduled scans (Mondays 2 AM UTC)
- ✅ Manual execution via workflow_dispatch
- ✅ Triggers on Go file changes
- ✅ Essential permissions only

---

### 2. `govulncheck-pr.yml` - Feature-Rich Configuration  
**Best for**: Comprehensive monitoring with detailed reporting
```yaml
# Enhanced workflow with Go detection and detailed summaries
# Includes comprehensive error handling and user feedback
```

**Features**:
- ✅ All features from simple configuration
- ✅ Automatic Go project detection
- ✅ Detailed workflow summaries
- ✅ Smart skipping for non-Go repositories
- ✅ Enhanced error handling and user feedback

---

### 3. `govulncheck-pr-production.yml` - Enterprise-Ready
**Best for**: Production environments, enterprise repositories
```yaml
# Production-grade workflow with advanced features
# Includes comprehensive monitoring and security considerations
```

**Features**:
- ✅ All features from feature-rich configuration
- ✅ Advanced conditional execution
- ✅ Security-events permission for future integration
- ✅ Comprehensive step summaries
- ✅ Enhanced project detection logic
- ✅ Detailed troubleshooting guidance

---

## 🚀 Quick Start

### Step 1: Choose Your Workflow
Pick the workflow that best fits your needs:
- **New to govulncheck?** → Use `govulncheck-pr-simple.yml`
- **Want detailed reporting?** → Use `govulncheck-pr.yml`  
- **Production environment?** → Use `govulncheck-pr-production.yml`

### Step 2: Copy to Your Repository
```bash
# Copy your chosen workflow
cp govulncheck-pr-simple.yml /path/to/your/repo/.github/workflows/govulncheck-pr.yml
```

### Step 3: Customize (Optional)
Adjust these common settings:
```yaml
# Change schedule
schedule:
  - cron: '0 6 * * 1'  # Mondays at 6 AM instead of 2 AM

# Change base branch
base-branch: develop  # Instead of main

# Change runner
runs-on: ubuntu-latest  # Instead of external-k8s-v2
```

### Step 4: Commit and Push
```bash
git add .github/workflows/govulncheck-pr.yml
git commit -m "Add govulncheck-pr vulnerability scanning workflow"
git push
```

## 🔧 Configuration Options

### Schedule Customization
```yaml
schedule:
  - cron: '0 2 * * 1'    # Weekly on Monday 2 AM
  - cron: '0 2 * * *'    # Daily at 2 AM  
  - cron: '0 2 1 * *'    # Monthly on 1st at 2 AM
  - cron: '0 2 * * 1-5'  # Weekdays at 2 AM
```

### Trigger Customization
```yaml
# Scan specific branches
push:
  branches: [main, develop, release/*]

# Scan on specific file changes
push:
  paths:
    - 'go.mod'
    - 'go.sum'
    - '**/*.go'
    - 'Dockerfile'  # Also scan when Docker files change
```

### Runner Customization
```yaml
# For GitHub-hosted runners
runs-on: ubuntu-latest

# For self-hosted runners
runs-on: [self-hosted, linux]

# For Bayer external runners (default)
runs-on: [external-k8s-v2]
```

## 📋 Required Permissions

All workflows require these permissions:
```yaml
permissions:
  contents: write      # Create branches and commit files
  pull-requests: write # Create and update PRs
  issues: write        # Fallback issue creation (if PR fails)
```

## 🎯 Expected Behavior

### When Vulnerabilities Are Found
1. **Branch Creation**: Creates/updates branch `govulncheck-fixes`
2. **PR Creation**: Creates PR with title "🚨 Security: govulncheck found X reachable vulnerabilities"
3. **Detailed Report**: Adds comprehensive vulnerability report to PR description
4. **File Storage**: Saves detailed JSON report to `.govulncheck/latest.json`

### When No Vulnerabilities Found
1. **Clean Exit**: Workflow completes successfully
2. **No PR Created**: No unnecessary PRs or branches
3. **Summary**: Workflow summary indicates clean scan

### Subsequent Scans
1. **Smart Updates**: Updates existing PR instead of creating new ones
2. **Progress Comments**: Adds comments to PR about scan updates
3. **Report Updates**: Refreshes vulnerability reports with latest data

## 🔍 Monitoring & Alerts

### GitHub Notifications
- Watch for PRs with title containing "govulncheck found"
- Monitor the `govulncheck-fixes` branch for updates
- Check workflow run summaries for scan results

### Integration with Other Tools
```yaml
# Example: Notify teams on vulnerability detection
- name: Notify on vulnerabilities
  if: contains(github.event.pull_request.title, 'govulncheck found')
  uses: bayer-int/csgda-platform-actions/actions/teams-notifier@v1
  with:
    url: ${{ secrets.TEAMS_WEBHOOK }}
    subject: "Security Alert: Vulnerabilities Found"
    text: "Govulncheck detected vulnerabilities in ${{ github.repository }}"
```

## 🛠️ Troubleshooting

### Common Issues

**Issue**: Workflow runs but no PR is created
**Solution**: Check that vulnerabilities were actually found. The action only creates PRs when vulnerabilities exist.

**Issue**: Permission denied errors
**Solution**: Ensure the workflow has `contents: write` and `pull-requests: write` permissions.

**Issue**: "No Go files found" 
**Solution**: Verify your repository contains `.go` files or a `go.mod` file.

**Issue**: Action not found
**Solution**: Ensure you're using the correct action reference: `bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool`

### Debug Mode
Add this step for debugging:
```yaml
- name: Debug information
  run: |
    echo "Repository: ${{ github.repository }}"
    echo "Event: ${{ github.event_name }}"
    echo "Branch: ${{ github.ref_name }}"
    find . -name "*.go" -type f | head -5
    ls -la go.*
```

## 📚 Additional Resources

- **Action Documentation**: [govulncheck-pr README](../actions/govulncheck-pr/README.md)
- **Govulncheck Tool**: https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck
- **GitHub Actions**: https://docs.github.com/en/actions
- **Bayer Platform Actions**: https://github.com/bayer-int/csgda-platform-actions

## 📝 Templates

See `TEMPLATE-govulncheck-pr.md` for copy-paste ready workflow examples.

---

**Need Help?** Check the troubleshooting section above or review the action logs for detailed error information.
