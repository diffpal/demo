# Tiny Orders review policy

Review this repository for actionable correctness, security, authorization,
and data-integrity problems.

Repository-specific invariant:

- Order prices and totals must be derived from trusted server-side state.

Review guidance:

- Look for broken trust boundaries, access-control regressions, ignored errors,
  and false-success behavior.
- Treat removed checks and other deleted behavior as first-class changed-code
  evidence, anchored to old-file lines on the left side of the diff.
- Give every actionable issue its own finding; do not leave an issue only in
  the review summary.
- Report findings only when they are actionable and supported by changed code.
- Prefer correctness and security findings over style suggestions.
- Do not report formatting, naming, or subjective refactoring suggestions.
