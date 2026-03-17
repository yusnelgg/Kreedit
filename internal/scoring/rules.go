package scoring

import (
	"github.com/yusnelgg/kreedit/config"
	"github.com/yusnelgg/kreedit/internal/domain"
)

func calcDebtRatio(app domain.CreditApplication, cfg *config.Rules) (int, string) {
	if app.MonthlyIncome <= 0 {
		return cfg.DebtRatio.Points.Critical, "no_income_data"
	}
	ratio := app.MonthlyDebt / app.MonthlyIncome
	t := cfg.DebtRatio.Thresholds
	p := cfg.DebtRatio.Points
	switch {
	case ratio <= t.Excellent:
		return p.Excellent, "excellent_debt_ratio"
	case ratio <= t.Good:
		return p.Good, "good_debt_ratio"
	case ratio <= t.Moderate:
		return p.Moderate, "moderate_debt_ratio"
	case ratio <= t.High:
		return p.High, "high_debt_ratio"
	default:
		return p.Critical, "critical_debt_ratio"
	}
}

func calcPaymentHistory(app domain.CreditApplication, cfg *config.Rules) (int, string) {
	p := cfg.PaymentHistory.Points
	switch {
	case app.MissedPayments == 0:
		return p.Perfect, "perfect_payment_history"
	case app.MissedPayments == 1:
		return p.Good, "good_payment_history"
	case app.MissedPayments == 2:
		return p.Fair, "fair_payment_history"
	case app.MissedPayments <= 4:
		return p.Poor, "poor_payment_history"
	default:
		return p.VeryPoor, "very_poor_payment_history"
	}
}

func calcCreditAge(app domain.CreditApplication, cfg *config.Rules) (int, string) {
	t := cfg.CreditAge.Thresholds
	p := cfg.CreditAge.Points
	switch {
	case app.CreditHistory >= t.Long:
		return p.Long, "long_credit_history"
	case app.CreditHistory >= t.Good:
		return p.Good, "good_credit_history"
	case app.CreditHistory >= t.Moderate:
		return p.Moderate, "moderate_credit_history"
	case app.CreditHistory >= t.Short:
		return p.Short, "short_credit_history"
	default:
		return p.VeryShort, "very_short_credit_history"
	}
}

func calcAgeScore(app domain.CreditApplication, cfg *config.Rules) (int, string) {
	t := cfg.AgeScore.Thresholds
	p := cfg.AgeScore.Points
	switch {
	case app.Age >= t.OptimalMin && app.Age <= t.OptimalMax:
		return p.Optimal, "optimal_age_range"
	case app.Age >= t.GoodMin && app.Age <= t.GoodMax:
		return p.Good, "good_age_range"
	case app.Age >= 18:
		return p.Acceptable, "acceptable_age_range"
	default:
		return 0, "invalid_age"
	}
}
