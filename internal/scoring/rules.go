package scoring

import "github.com/yusnelgg/kreedit/internal/domain"

func calcDebtRatio(app domain.CreditApplication) (int, string) {
	if app.MonthlyIncome <= 0 {
		return 20, "no_income_data"
	}

	ratio := app.MonthlyDebt / app.MonthlyIncome
	switch {
	case ratio < 0.20:
		return 300, "excellent_debt_ratio"
	case ratio < 0.35:
		return 240, "good_debt_ratio"
	case ratio < 0.50:
		return 160, "moderate_debt_ratio"
	case ratio < 0.65:
		return 80, "high_debt_ratio"
	default:
		return 20, "critical_debt_ratio"
	}
}

func calcPaymentHistory(app domain.CreditApplication) (int, string) {
	switch {
	case app.MissedPayments == 0:
		return 350, "perfect_payment_history"
	case app.MissedPayments == 1:
		return 260, "good_payment_history"
	case app.MissedPayments == 2:
		return 180, "fair_payment_history"
	case app.MissedPayments == 4:
		return 90, "poor_payment_history"
	default:
		return 20, "very_poor_payment_history"
	}
}

func calcCreditAge(app domain.CreditApplication) (int, string) {
	switch {
	case app.CreditHistory >= 84:
		return 200, "long_credit_history"
	case app.CreditHistory >= 48:
		return 160, "good_credit_history"
	case app.CreditHistory >= 24:
		return 100, "moderate_credit_history"
	case app.CreditHistory >= 12:
		return 50, "short_credit_history"
	default:
		return 10, "very_short_credit_history"
	}
}

func calcAgeScore(app domain.CreditApplication) (int, string) {
	switch {
	case app.Age >= 30 && app.Age <= 55:
		return 150, "optimal_age_range"
	case app.Age >= 25 && app.Age <= 60:
		return 110, "good_age_range"
	case app.Age >= 18:
		return 70, "acceptable_age_range"
	default:
		return 0, "invalid_age"
	}
}
