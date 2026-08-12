# Tiny Orders review policy

Review this repository for actionable correctness, security, authorization,
and data-integrity problems.

Repository invariants:

- Product prices and totals must always be derived from the server-side catalog.
- Never trust client-supplied monetary values.
- A user may only read orders owned by that user.
- Repository and persistence errors must be propagated.
- The API must not report success when an order was not saved.
- Report findings only when they are actionable and supported by changed code.
- Prefer correctness and security findings over style suggestions.
- Do not report formatting, naming, or subjective refactoring suggestions.
