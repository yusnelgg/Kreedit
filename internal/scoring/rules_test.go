package scoring

import (
	"testing"

	"github.com/yusnelgg/kreedit/internal/domain"
)

func TestCalcDebtRatio(t *testing.T) {
	cfg := loadTestConfig()

	tests := []struct {
		name       string
		income     float64
		debt       float64
		wantPoints int
		wantReason string
	}{
		{"ratio excelente", 5000, 500, 300, "excellent_debt_ratio"},
		{"ratio bueno", 5000, 1500, 240, "good_debt_ratio"},
		{"ratio moderado", 5000, 2200, 160, "moderate_debt_ratio"},
		{"ratio alto", 5000, 3000, 80, "high_debt_ratio"},
		{"ratio critico", 5000, 4500, 20, "critical_debt_ratio"},
		{"sin ingreso", 0, 1000, 20, "no_income_data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := domain.CreditApplication{MonthlyIncome: tt.income, MonthlyDebt: tt.debt}
			pts, reason := calcDebtRatio(app, cfg)
			if pts != tt.wantPoints {
				t.Errorf("puntos: got %d, want %d", pts, tt.wantPoints)
			}
			if reason != tt.wantReason {
				t.Errorf("reason: got %s, want %s", reason, tt.wantReason)
			}
		})
	}
}

func TestCalcPaymentHistory(t *testing.T) {
	cfg := loadTestConfig()

	tests := []struct {
		name       string
		missed     int
		wantPoints int
		wantReason string
	}{
		{"sin atrasos", 0, 350, "perfect_payment_history"},
		{"un atraso", 1, 260, "good_payment_history"},
		{"dos atrasos", 2, 180, "fair_payment_history"},
		{"cuatro atrasos", 4, 90, "poor_payment_history"},
		{"muchos atrasos", 8, 20, "very_poor_payment_history"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := domain.CreditApplication{MissedPayments: tt.missed}
			pts, reason := calcPaymentHistory(app, cfg)
			if pts != tt.wantPoints {
				t.Errorf("puntos: got %d, want %d", pts, tt.wantPoints)
			}
			if reason != tt.wantReason {
				t.Errorf("reason: got %s, want %s", reason, tt.wantReason)
			}
		})
	}
}

func TestCalcCreditAge(t *testing.T) {
	cfg := loadTestConfig()

	tests := []struct {
		name       string
		months     int
		wantPoints int
		wantReason string
	}{
		{"historial largo", 96, 200, "long_credit_history"},
		{"historial bueno", 60, 160, "good_credit_history"},
		{"historial moderado", 30, 100, "moderate_credit_history"},
		{"historial corto", 18, 50, "short_credit_history"},
		{"sin historial", 3, 10, "very_short_credit_history"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := domain.CreditApplication{CreditHistory: tt.months}
			pts, reason := calcCreditAge(app, cfg)
			if pts != tt.wantPoints {
				t.Errorf("puntos: got %d, want %d", pts, tt.wantPoints)
			}
			if reason != tt.wantReason {
				t.Errorf("reason: got %s, want %s", reason, tt.wantReason)
			}
		})
	}
}

func TestCalcAgeScore(t *testing.T) {
	cfg := loadTestConfig()

	tests := []struct {
		name       string
		age        int
		wantPoints int
		wantReason string
	}{
		{"edad optima", 35, 150, "optimal_age_range"},
		{"edad buena", 27, 110, "good_age_range"},
		{"edad aceptable", 19, 70, "acceptable_age_range"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := domain.CreditApplication{Age: tt.age}
			pts, reason := calcAgeScore(app, cfg)
			if pts != tt.wantPoints {
				t.Errorf("puntos: got %d, want %d", pts, tt.wantPoints)
			}
			if reason != tt.wantReason {
				t.Errorf("reason: got %s, want %s", reason, tt.wantReason)
			}
		})
	}
}
