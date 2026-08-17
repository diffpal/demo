# Tiny Orders review policy

Review this repository for actionable correctness, security, authorization,
and data-integrity problems.

Repository-specific invariant:

- Order prices and totals must be derived from trusted server-side state.

Review guidance:

- Look for broken trust boundaries, access-control regressions, ignored errors,
  and false-success behavior.
- Report findings only when they are actionable and supported by changed code.
- Prefer correctness and security findings over style suggestions.
- Do not report formatting, naming, or subjective refactoring suggestions.
