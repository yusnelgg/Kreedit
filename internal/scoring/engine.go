package scoring

import (
	"fmt"
	"time"

	"github.com/yusnelgg/kreedit/internal/domain"
)

const ModelVersion = "v1.0.0"

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

// Score es el método principal — recibe una solicitud y devuelve el resultado
func (e *Engine) Score(app domain.CreditApplication) domain.ScoringResult {

	// 1. Calcular cada regla
	debtPts, debtReason := calcDebtRatio(app)
	paymentPts, paymentReason := calcPaymentHistory(app)
	creditAgePts, creditReason := calcCreditAge(app)
	agePts, ageReason := calcAgeScore(app)

	// 2. Armar el breakdown
	breakdown := domain.ScoreBreakdown{
		DebtRatioScore:      debtPts,
		PaymentHistoryScore: paymentPts,
		CreditAgeScore:      creditAgePts,
		AgeScore:            agePts,
		Total:               debtPts + paymentPts + creditAgePts + agePts,
	}

	// 3. Juntar las razones
	reasons := []string{debtReason, paymentReason, creditReason, ageReason}

	// 4. Tomar la decisión
	decision, riskTier := decide(breakdown.Total)

	// 5. Calcular límite de crédito
	creditLimit := calcCreditLimit(app.MonthlyIncome, decision)

	// 6. Armar y devolver el resultado
	return domain.ScoringResult{
		ApplicationID: generateID(),
		ApplicantID:   app.ApplicantID,
		Score:         breakdown.Total,
		Decision:      decision,
		RiskTier:      riskTier,
		CreditLimit:   creditLimit,
		Reasons:       reasons,
		Breakdown:     breakdown,
		ModelVersion:  ModelVersion,
		ProcessedAt:   time.Now(),
	}
}

// decide convierte el score en una decisión y tier
func decide(score int) (domain.Decision, domain.RiskTier) {
	switch {
	case score >= 800:
		return domain.DecisionApproved, domain.RiskTierA
	case score >= 700:
		return domain.DecisionApproved, domain.RiskTierB
	case score >= 500:
		return domain.DecisionReview, domain.RiskTierC
	default:
		return domain.DecisionRejected, domain.RiskTierD
	}
}

// calcCreditLimit define cuánto dinero puede pedir según la decisión
func calcCreditLimit(monthlyIncome float64, decision domain.Decision) float64 {
	switch decision {
	case domain.DecisionApproved:
		return monthlyIncome * 4
	case domain.DecisionReview:
		return monthlyIncome * 1.5
	default:
		return 0
	}
}

// generateID crea un ID único para cada solicitud
func generateID() string {
	return fmt.Sprintf("app_%d", time.Now().UnixNano())
}
