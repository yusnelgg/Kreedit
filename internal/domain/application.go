package domain

import "time"

type Decision string

const (
	DecisionApproved Decision = "approved"
	DecisionReview   Decision = "review"
	DecisionRejected Decision = "rejected"
)

type RiskTier string

const (
	RiskTierA RiskTier = "A"
	RiskTierB RiskTier = "B"
	RiskTierC RiskTier = "C"
	RiskTierD RiskTier = "D"
)

type CreditApplication struct {
	ApplicantID     string  `json:"applicant_id"`
	Age             int     `json:"age"`
	MonthlyIncome   float64 `json:"monthly_income"`
	MonthlyDebt     float64 `json:"monthly_debt"`
	CreditHistory   int     `json:"credit_history_months"`
	MissedPayments  int     `json:"missed_payments"`
	RequestedAmount float64 `json:"requested_amount"`
}

type ScoreBreakdown struct {
	DebtRatioScore      int `json:"debt_ratio_score"`
	PaymentHistoryScore int `json:"payment_history_score"`
	CreditAgeScore      int `json:"credit_age_score"`
	AgeScore            int `json:"age_score"`
	Total               int `json:"total"`
}

type ScoringResult struct {
	ApplicationID string         `json:"application_id"`
	ApplicantID   string         `json:"applicant_id"`
	Score         int            `json:"score"`
	Decision      Decision       `json:"decision"`
	RiskTier      RiskTier       `json:"risk_tier"`
	CreditLimit   float64        `json:"credit_limit"`
	Reasons       []string       `json:"reasons"`
	Breakdown     ScoreBreakdown `json:"breakdown"`
	ModelVersion  string         `json:"model_version"`
	ProcessedAt   time.Time      `json:"processed_at"`
}
