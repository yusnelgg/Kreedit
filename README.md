# Kreedit

Real-time credit scoring engine with explainable decisions, configurable risk rules, and full audit trail — built in Go.

## Overview

Kreedit evaluates loan applications in under 500ms. Submit an application, get back a score, a risk tier, a credit limit, and exactly why the decision was made.
```json
{
  "score": 960,
  "decision": "approved",
  "risk_tier": "A",
  "credit_limit": 20000,
  "reasons": [
    "excellent_debt_ratio",
    "perfect_payment_history",
    "good_credit_history",
    "optimal_age_range"
  ]
}
```

## Features

- Real-time scoring in under 500ms
- Explainable decisions with human-readable reasons
- Risk tiers A / B / C / D
- Full score breakdown per rule
- 93.5% test coverage
- Model versioning on every decision

## Scoring Model

| Rule | Weight | Max points |
|---|---|---|
| Debt / income ratio | 30% | 300 |
| Payment history | 35% | 350 |
| Credit history age | 20% | 200 |
| Applicant age | 15% | 150 |
| **Total** | | **1000** |

## Decision Thresholds

| Score | Decision | Risk tier |
|---|---|---|
| 800 – 1000 | Approved | A |
| 700 – 799 | Approved | B |
| 500 – 699 | Manual review | C |
| 0 – 499 | Rejected | D |

## Getting Started

### Requirements

- Go 1.21+

### Run locally
```bash
git clone https://github.com/tuusuario/kreedit
cd kreedit
go mod tidy
go run cmd/api/main.go
```

## API

### Score an application
```
POST /api/v1/score
```

Request:
```json
{
  "applicant_id": "usr_001",
  "age": 32,
  "monthly_income": 5000,
  "monthly_debt": 800,
  "credit_history_months": 48,
  "missed_payments": 0,
  "requested_amount": 20000
}
```

Response:
```json
{
  "application_id": "app_1234567890",
  "applicant_id": "usr_001",
  "score": 870,
  "decision": "approved",
  "risk_tier": "A",
  "credit_limit": 20000,
  "reasons": ["excellent_debt_ratio", "perfect_payment_history"],
  "breakdown": {
    "debt_ratio_score": 300,
    "payment_history_score": 350,
    "credit_age_score": 160,
    "age_score": 150,
    "total": 960
  },
  "model_version": "v1.0.0",
  "processed_at": "2026-03-17T14:32:00Z"
}
```

### Health check
```
GET /api/v1/health
```

## Project Structure
```
kreedit/
├── cmd/api/          entry point
├── internal/
│   ├── domain/       types and constants
│   ├── scoring/      rules engine and tests
│   ├── storage/      audit trail (coming soon)
│   └── api/          HTTP handler
└── config/           configurable rules (coming soon)
```

## Tests
```bash
go test ./... -cover
```

coverage: 93.5% of statements

## License

MIT