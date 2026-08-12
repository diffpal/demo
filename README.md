# DiffPal Demo

A tiny Go service with a deliberately flawed pull request showing how DiffPal
reviews code inside GitHub Actions.

[View the live demo review](https://github.com/diffpal/demo/pulls) · [DiffPal documentation](https://diffpal.github.io) · [DiffPal source repository](https://github.com/diffpal/diffpal)

## What this demonstrates

The open demo pull request shows normal tests passing alongside a DiffPal review
summary, actionable inline findings, repository-owned review instructions, an
optional merge gate, and machine-readable artifacts.

## Try the service locally

```bash
go test ./...
go run ./cmd/server
```

Create an order:

```bash
curl -X POST http://localhost:8080/orders \
  -H 'Content-Type: application/json' -H 'X-User-ID: user-123' \
  -d '{"product_id":"keyboard","quantity":2}'
```

Retrieve it:

```bash
curl http://localhost:8080/orders/ord-000001 -H 'X-User-ID: user-123'
```

## Run DiffPal in your own repository

1. Create a repository from this template.
2. Add an `OPENAI_API_KEY` repository secret.
3. Open a same-repository pull request.
4. Inspect the DiffPal review.

The configured provider receives review context. Keep secrets out of workflows
that execute untrusted fork pull request code.

## Why the DiffPal check is red

The red DiffPal check is intentional. The permanent demo PR contains
high-severity regressions so the repository can demonstrate merge gating.

## Demo repository policy

`main` is the clean baseline. The canonical demo PR remains open and should not
be merged.
