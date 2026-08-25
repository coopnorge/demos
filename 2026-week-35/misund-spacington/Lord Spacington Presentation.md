# 🎩 Lord Spacington the Third
### Devourer of Tabs

---

## Focus on what you can control,
## never be bothered by the rest.

---

## Source Control Management

- Make our lives easier
- Prevent errors
- Help us recover from errors

---
## Source Control Management

- Before: Folders with some state
- After: ✨ Patches ✨
- All the versions

---

```dataviewjs
await dv.view("contact-view", {
  contacts: ["Just Thomas Hiorth Misund"],
  style: "card",
});
```

---
## Wikipedia introduction

> **Version control** (also known as **revision control**, **source control**, and **source code management**) is the [software engineering](https://en.wikipedia.org/wiki/Software_engineering_management "Software engineering management") practice of controlling, organizing, and tracking different versions in history of [computer files](https://en.wikipedia.org/wiki/Computer_file "Computer file") – primarily [source code](https://en.wikipedia.org/wiki/Source_code "Source code") [text files](https://en.wikipedia.org/wiki/Text_files "Text files"), but generally any type of file.

---
## Git is for Source Code

- Source code in git
- Artifacts in a registry

---
## The Complication

- Generated code
	- It's Source Code (tm)
- Existing pipelines that bypass artifact registries
- Failures with deterministic fixes
	- `lint --auto-fix`
	- `format`

---
## Errare Humanum Est

- I forgot to regenerate when the generator was upgraded
- I forgot to build when I had changed the source code
- I sometimes edited stuff in an editor without the linter
- We accept contributions from outside the team

---
## First Level of Maturity

- There is drift
- We recognize it after a while
- We fix it (build and push)

---
## Second Level of Maturity

- There is drift
- We catch it in a test
- We fix it before we merge

---
## The Grunt

- A always leads to B
- Outdated generated code: regenerate
- Forgot to build: Build
- Tabs in place of spaces: Push the button

---
## Repetetive Manual Labor

- Grunt work
- Human middleware
- Automation debt

---
## We Know What's Wrong

- The test went red
- To get there we had to calculate the difference between
	- Status quo
	- The passing version

---
## ELI5

- It's a ✨ patch ✨

---

## 🎩 Enter Lord Spacington the Third
### Devourer of Tabs


He commits the patch.

---
## 🎩 Lord Spacington

A CI automation pattern that performs **deterministic, autofixable** repository maintenance and commits the result back to a PR branch.

- Runs fix/generate commands in CI
- Commits changes automatically  
- Rebases before pushing to handle concurrency
- Stops cleanly if it can't rebase

---
## 🎩 Third Level of Matutiry: Deterministic Automation

Stop having humans stop in their tracks to fix formatting:

1. Run the build/lint/formatter command
2. Detect changes
3. Commit the result to the PR branch
4. Rebase before push (concurrency safety)
5. Push the commits back

The PR will pass tests and everyone's happy.

---

## Maturity Model for Generated Artifacts

**Level 1: Ignore**
- Observe the drift ("I am a leaf on the wind, watch how I soar")

**Level 2: Detect**
- Fail in CI

**🎩 Level 3: Fix**
- Rebuild and commit in CI

**Level 4: Omit**
- Ignore in git, build in CI, publish to registry

---

## The Concurrency Problem

What if two PRs are being auto-fixed at the same time?

```bash
Commit A: Push commit "auto-format"
      ↓ (race condition)
Commit B: Push commit "auto-build"
      ↓ (race condition)
Commit C: Push commit "auto-lint"
      ↓ (race condition)
Commit D: Push commit "go generate"
```

---

## Fix Concurrency

Before pushing, rebase against the target branch:

```bash
git pull --rebase origin $branch
git push origin HEAD:$branch
```

- Rebasing will fix most issues
- If you still have a conflict, you just won't push
- Lord Spacington will try again on the commit you have a conflict with

---

## Real-World Examples

### IDP Auth0 Config: Prettier

`.github/workflows/commit-prettier-format.yaml`

- Runs `npm run fmt` (Prettier)
- Commits to the PR branch
- Fail on changes

---

## Real-World Examples

### IDP Auth0 Config: Terraform

`.github/workflows/commit-tofu-fmt.yaml`

- Runs `tofu fmt -recursive`
- Auto-formats HCL
- Commits to the PR branch
- Fail on changes
 
---

## Real-World Examples

### CMDM: Go Style Fixes

`.github/workflows/commit-go-style.yaml`

Runs on every PR:
- `go fix ./...`
- `mage go:fix`
- Commits to the PR branch
- Fail on changes

---
## Real-World Examples

### CMDM: Go Style Fixes

Without his Lordship, some grunt is required:
https://github.com/coopnorge/idp-db-sync/pull/184
https://github.com/coopnorge/idp-auth0-db-backup/pull/1091

With his Lordship, the PR gets itself to a passing state:
https://github.com/coopnorge/customer-master-data-management/pull/1433/commits (concurrency resolution)
https://github.com/coopnorge/idp-auth0-config/pull/4900/commits (with concurrency resolution)

---

![[Pasted image 20260825121019.png]]

---
![[Pasted image 20260825121042.png]]
---
## The Custom Action

```yaml
name: Commit and Push with Lord Spacington
```

Reusable GitHub Action that:
- Check if the repo is dirty
- Configure git identity ("Lord Spacington the Third")
- Commit changes
- Rebase (best-effort concurrency fix)
- Push
- Output whether we produced changes

[Sauce (reusable action)](https://github.com/coopnorge/customer-master-data-management/blob/b22350e/.github/actions/commit-and-push-with-spacington/action.yaml)
[Example usage (tofo fmt workflow)](https://github.com/coopnorge/customer-master-data-management/blob/b22350e/.github/workflows/commit-tofu-fmt.yaml)

---
## Implementation Checklist

To add Lord Spacington to your repo:

1. **Create a fix workflow** (e.g., `commit-tofu-fmt.yaml`)
   - Make automated changes here

2. **Use the custom action**
   ```yaml
   - uses: ./.github/actions/commit-and-push-with-spacington
     with:
       commit-message: "chore: auto-format with Prettier 🎩"
       branch: ${{ github.event.pull_request.head.ref }}
   ```

3. **Fail the job if changes were detected**
   ```yaml
   - name: Fail on changes detected
     if: steps.commit.outputs.changed == 'true'
     run: exit 1
   ```
   This keeps people from merging the faulty PR without the fix.

[Sauce (reusable action)](https://github.com/coopnorge/customer-master-data-management/blob/b22350e/.github/actions/commit-and-push-with-spacington/action.yaml)
[Example usage (tofu fmt workflow)](https://github.com/coopnorge/customer-master-data-management/blob/b22350e/.github/workflows/commit-tofu-fmt.yaml)

---
## Principles

1. **Automation over enforcement**
   - Fix, don't fail

1. **When to use it**
   - Deterministic tools

2. **Transparent to developers**
   - Clear git identity ("Lord Spacington")
   - Clear commit messages
 
---

## The Philosophy

> "Never bother with things you can't control."


---

## Be Bothered by

- Source code
- Design decisions
- Logic

---

## Do Not Be Bothered by

- Linting
- Formatting
- Tabs and semicolons
- Artifacts
- Trusted, deterministic modernization

---

## Repositories using Lord Spacington

  - `customer-master-data-management`
  - `idp-auth0-config`
  - `idp-user-service`


---

## Errare Humanum Est

... perseverare autem diabolicum

---

## Thank You

**🎩 Lord Spacington the Third, Devourer of Tabs**

*WHEN YOU FORGET, IT JUST WORKS*