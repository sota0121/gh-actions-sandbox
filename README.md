# GitHub Actions Sandbox

## References

- [GitHub Actions 実践入門](https://amzn.asia/d/hjMxwqo)
  - Sample code is [here](https://github.com/github-actions-up-and-running/)


## About this repository

- This repository is for GitHub Actions Sandbox.
- I will try to use GitHub Actions in this repository with sample code.

### Sample code

- `cmd/main.go`
- `lib/*.go`

### GitHub Actions Workflow

- `.github/workflows/*.yml`

---

## Workflows description

### **`Simple Continuous Integration`**

#### Related files

- workflow file
  - `.github/workflows/ci.yml`
- used by workflow file
  - `Makefile`
- target files
  - `cmd/main.go`
  - `lib/*.go`

#### What is this workflow?

1. Triggered by `push` or `pull_request` event
2. Set up Go environment
3. Run `make test` command
4. Run `golangci-lint run` command
5. If `make test` and `golangci-lint run` are successful, then run `make build` command


### **`GolangCI-Lint`**

#### Related files

- workflow file
  - `.github/workflows/golangci-lint.yml`
- used by workflow file
  - `.golangci.yml`
- target files
  - `cmd/main.go`
  - `lib/*.go`


#### What is this workflow?

1. Triggered by below events
   - `push` event with tags (e.g. `v1.0.0`) on `master` or `main` branch
   - `pull_request` event
2. Set up Go environment
3. Use [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action) action


### **`reviewdog + golangci-lint`**

#### Related files

- workflow file
  - `.github/workflows/reviewdog-golangci-lint.yml`

