package scoring

import (
	"testing"

	"github.com/yusnelgg/kreedit/internal/domain"
)

func TestScore(t *testing.T) {
	engine := NewEngine()

	tests := []struct {
		name         string
		input        domain.CreditApplication
		wantDecision domain.Decision
		wantRiskTier domain.RiskTier
		wantScoreMin int
		wantScoreMax int
		wantLimit    float64
	}{
		{
			name: "solicitante perfecto",
			input: domain.CreditApplication{
				ApplicantID: "usr_001", Age: 35,
				MonthlyIncome: 5000, MonthlyDebt: 500,
				CreditHistory: 96, MissedPayments: 0,
				RequestedAmount: 10000,
			},
			wantDecision: domain.DecisionApproved,
			wantRiskTier: domain.RiskTierA,
			wantScoreMin: 900, wantScoreMax: 1000,
			wantLimit: 20000,
		},
		{
			name: "deuda alta",
			input: domain.CreditApplication{
				ApplicantID: "usr_002", Age: 28,
				MonthlyIncome: 3000, MonthlyDebt: 2400,
				CreditHistory: 24, MissedPayments: 2,
				RequestedAmount: 5000,
			},
			wantDecision: domain.DecisionRejected,
			wantRiskTier: domain.RiskTierD,
			wantScoreMin: 0, wantScoreMax: 499,
			wantLimit: 0,
		},
		{
			name: "sin historial crediticio",
			input: domain.CreditApplication{
				ApplicantID: "usr_004", Age: 22,
				MonthlyIncome: 2000, MonthlyDebt: 0,
				CreditHistory: 0, MissedPayments: 0,
				RequestedAmount: 3000,
			},
			wantDecision: domain.DecisionApproved,
			wantRiskTier: domain.RiskTierB,
			wantScoreMin: 700, wantScoreMax: 799,
			wantLimit: 8000,
		},
		{
			name: "score tier B",
			input: domain.CreditApplication{
				ApplicantID: "usr_005", Age: 40,
				MonthlyIncome: 4000, MonthlyDebt: 1200,
				CreditHistory: 48, MissedPayments: 1,
				RequestedAmount: 8000,
			},
			wantDecision: domain.DecisionApproved,
			wantRiskTier: domain.RiskTierA,
			wantScoreMin: 800, wantScoreMax: 1000,
			wantLimit: 16000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.Score(tt.input)

			if result.Decision != tt.wantDecision {
				t.Errorf("decision: got %s, want %s", result.Decision, tt.wantDecision)
			}
			if result.RiskTier != tt.wantRiskTier {
				t.Errorf("risk tier: got %s, want %s", result.RiskTier, tt.wantRiskTier)
			}
			if result.Score < tt.wantScoreMin || result.Score > tt.wantScoreMax {
				t.Errorf("score %d fuera del rango [%d-%d]", result.Score, tt.wantScoreMin, tt.wantScoreMax)
			}
			if result.CreditLimit != tt.wantLimit {
				t.Errorf("credit limit: got %.0f, want %.0f", result.CreditLimit, tt.wantLimit)
			}
			if result.ApplicationID == "" {
				t.Error("application_id no puede estar vacío")
			}
			if len(result.Reasons) == 0 {
				t.Error("reasons no puede estar vacío")
			}
		})
	}
}
