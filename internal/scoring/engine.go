package scoring

import (
	"fmt"
	"time"

	"github.com/yusnelgg/kreedit/config"
	"github.com/yusnelgg/kreedit/internal/domain"
)

type Engine struct {
	cfg *config.Rules
}

func NewEngine(cfg *config.Rules) *Engine {
	return &Engine{cfg: cfg}
}

func (e *Engine) Score(app domain.CreditApplication) domain.ScoringResult {
	debtPts, debtReason := calcDebtRatio(app, e.cfg)
	paymentPts, paymentReason := calcPaymentHistory(app, e.cfg)
	creditAgePts, creditReason := calcCreditAge(app, e.cfg)
	agePts, ageReason := calcAgeScore(app, e.cfg)

	breakdown := domain.ScoreBreakdown{
		DebtRatioScore:      debtPts,
		PaymentHistoryScore: paymentPts,
		CreditAgeScore:      creditAgePts,
		AgeScore:            agePts,
		Total:               debtPts + paymentPts + creditAgePts + agePts,
	}

	reasons := []string{debtReason, paymentReason, creditReason, ageReason}
	decision, riskTier := decide(breakdown.Total, e.cfg)
	creditLimit := calcCreditLimit(app.MonthlyIncome, decision, e.cfg)

	return domain.ScoringResult{
		ApplicationID: generateID(),
		ApplicantID:   app.ApplicantID,
		Score:         breakdown.Total,
		Decision:      decision,
		RiskTier:      riskTier,
		CreditLimit:   creditLimit,
		Reasons:       reasons,
		Breakdown:     breakdown,
		ModelVersion:  e.cfg.ModelVersion,
		ProcessedAt:   time.Now(),
	}
}

func decide(score int, cfg *config.Rules) (domain.Decision, domain.RiskTier) {
	switch {
	case score >= cfg.Decisions.ApprovedA:
		return domain.DecisionApproved, domain.RiskTierA
	case score >= cfg.Decisions.ApprovedB:
		return domain.DecisionApproved, domain.RiskTierB
	case score >= cfg.Decisions.Review:
		return domain.DecisionReview, domain.RiskTierC
	default:
		return domain.DecisionRejected, domain.RiskTierD
	}
}

func calcCreditLimit(income float64, decision domain.Decision, cfg *config.Rules) float64 {
	m := cfg.CreditLimitMultipliers
	switch decision {
	case domain.DecisionApproved:
		return income * m.Approved
	case domain.DecisionReview:
		return income * m.Review
	default:
		return 0
	}
}

func generateID() string {
	return fmt.Sprintf("app_%d", time.Now().UnixNano())
}
