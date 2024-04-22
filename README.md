# golangci-lint

## Installation

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2
```

```bash
golangci-lint --version
```

## Run

Check all

```bash
golangci-lint run ./...
```

Check specific file

```bash
golangci-lint run govet.go
```
