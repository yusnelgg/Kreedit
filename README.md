![Kreedit logo](https://raw.githubusercontent.com/yusnelgg/Kreedit/refs/heads/main/imgs/kreedit.png)

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
  ],
  "breakdown": {
    "debt_ratio_score": 300,
    "payment_history_score": 350,
    "credit_age_score": 160,
    "age_score": 150,
    "total": 960
  }
}
```

## Features

- Real-time scoring in under 500ms
- Explainable decisions with human-readable reasons
- Configurable risk rules via YAML — no redeployment needed
- Risk tiers A / B / C / D
- Full score breakdown per rule
- Audit trail persisted in PostgreSQL
- Model versioning on every decision
- 94.4% test coverage
- Swagger documentation

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

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Router | Chi |
| Database | PostgreSQL (Supabase) |
| Config | YAML |
| Docs | Swagger |
| Tests | Table-driven tests |

## Getting Started

### Requirements

- Go 1.21+
- PostgreSQL database (or Supabase free tier)

### Run locally
```bash
git clone https://github.com/yusnelgg/kreedit
cd kreedit
go mod tidy
```

Create a `.env` file in the root:
```
DATABASE_URL=postgresql://user:password@host:5432/postgres?sslmode=require
```
```bash
go run cmd/api/main.go
```

## API

### Score an application
```
POST /api/v1/score
```
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

### Get scoring history
```
GET /api/v1/history/{applicantID}
```

### Health check
```
GET /api/v1/health
```

### Swagger UI
```
http://localhost:8080/swagger/index.html
```

## Project Structure
```
kreedit/
├── cmd/api/          entry point
├── config/           configurable rules via YAML
├── internal/
│   ├── domain/       types and constants
│   ├── scoring/      rules engine and tests
│   ├── storage/      PostgreSQL repository
│   └── api/          HTTP handler
└── docs/             swagger generated docs
```

## Tests
```bash
go test ./... -cover
```
```
coverage: 94.4% of statements
```

## License

MIT

