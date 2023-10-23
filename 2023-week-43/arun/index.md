---
title: Release Drafter
description: Your helper to almost creating a release
author: Arun Poudel
keywords: release, github, automation, go, golang
marp: true
theme: default
class:
  - invert
style: |
  #why-image {
    left: 10em;
  }
---

# Release Drafter

---
<!-- backgroundColor: #232529 -->

![Let the boss speak](./problem-statement.png)

---

## Full discussion on the issue

https://github.com/coopnorge/engineering-issues/issues/325

---

## Available Option

- Release Drafter (https://github.com/release-drafter/release-drafter, we didn't want to use commit message and/or branch names)
- Gorelease (Looked abandoned, but still works) (https://github.com/golang/go/issues/26420)
- Apidiff (https://pkg.go.dev/golang.org/x/exp/apidiff)

---
<!-- backgroundColor: #232529 -->

![bg](./precise-endeavor.jpg)

---
<!-- backgroundColor: #232529 -->

![Let the boss speak](./problem-statement-2.png)

---

### So, the solution was to use PR title and body 😉

https://github.com/release-drafter/release-drafter

---

# Kinda

---

#
<!-- backgroundColor: white -->

![Create PR](./create-pr-seq.png)

---

#
<!-- backgroundColor: white -->
![Merge](./merge-seq.png)

---
<!-- backgroundColor: #232529 -->

# Demo

---
<!-- backgroundColor: #232529 -->

# Further Improvements?

- Read the PR body and extract additional labels
- What does the team need?

---
<!-- backgroundColor: #232529 -->

# Thank you