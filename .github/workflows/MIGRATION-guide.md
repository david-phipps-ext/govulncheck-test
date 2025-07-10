# Migration Guide: Feature Branch to Main

This guide explains how to update the workflows once the `govulncheck-pr` action is merged from the feature branch `phipps/govulncheck-pr-tool` to the main branch.

## Current State (Testing)

All workflows currently use:
```yaml
uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@phipps/govulncheck-pr-tool
```

## After Merge (Production)

Once the feature branch is merged to main and tagged, update to:
```yaml
uses: bayer-int/csgda-platform-actions/actions/govulncheck-pr@v1
```

## Files to Update

### Workflow Files
1. `govulncheck-pr.yml` (line 48)
2. `govulncheck-pr-simple.yml` (line 22)  
3. `govulncheck-pr-production.yml` (line 62)

### Documentation Files
1. `README-govulncheck-pr.md` (3 references)
2. `TEMPLATE-govulncheck-pr.md` (3 references)
3. `README-workflows.md` (1 reference)

## Quick Update Commands

Use these sed commands to update all references at once:

### For macOS:
```bash
find . -name "*.yml" -o -name "*.md" | xargs sed -i '' 's/@phipps\/govulncheck-pr-tool/@v1/g'
```

### For Linux:
```bash
find . -name "*.yml" -o -name "*.md" | xargs sed -i 's/@phipps\/govulncheck-pr-tool/@v1/g'
```

## Manual Update Process

If you prefer to update manually:

1. **Search and Replace**: Use your editor's find/replace function
   - Find: `@phipps/govulncheck-pr-tool`
   - Replace: `@v1`

2. **Remove Development Notes**: Remove the warning notes about feature branch usage from:
   - `README-govulncheck-pr.md` (line 3-4)
   - `TEMPLATE-govulncheck-pr.md` (line 3-4)  
   - `README-workflows.md` (line 3-4)

## Verification

After updating, verify the changes:

```bash
# Check that no feature branch references remain
grep -r "phipps/govulncheck-pr-tool" .

# Check that v1 references are present
grep -r "@v1" . | grep govulncheck-pr

# Validate workflow syntax
yamllint .github/workflows/*.yml
```

## Testing After Migration

1. **Commit Changes**: Commit the updated workflows
2. **Test Manually**: Run `workflow_dispatch` to test the action
3. **Monitor Scheduled Runs**: Check that scheduled scans work correctly
4. **Verify PR Creation**: Ensure PRs are created when vulnerabilities are found

## Rollback Plan

If issues occur after migration, temporarily revert to the feature branch:

```bash
find . -name "*.yml" -o -name "*.md" | xargs sed -i 's/@v1/@phipps\/govulncheck-pr-tool/g'
```

Then investigate and fix the issue before re-attempting the migration.
