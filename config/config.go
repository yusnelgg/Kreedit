package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Rules struct {
	ModelVersion string `yaml:"model_version"`

	DebtRatio struct {
		Thresholds struct {
			Excellent float64 `yaml:"excellent"`
			Good      float64 `yaml:"good"`
			Moderate  float64 `yaml:"moderate"`
			High      float64 `yaml:"high"`
		} `yaml:"thresholds"`
		Points struct {
			Excellent int `yaml:"excellent"`
			Good      int `yaml:"good"`
			Moderate  int `yaml:"moderate"`
			High      int `yaml:"high"`
			Critical  int `yaml:"critical"`
		} `yaml:"points"`
	} `yaml:"debt_ratio"`

	PaymentHistory struct {
		Points struct {
			Perfect  int `yaml:"perfect"`
			Good     int `yaml:"good"`
			Fair     int `yaml:"fair"`
			Poor     int `yaml:"poor"`
			VeryPoor int `yaml:"very_poor"`
		} `yaml:"points"`
	} `yaml:"payment_history"`

	CreditAge struct {
		Thresholds struct {
			Long     int `yaml:"long"`
			Good     int `yaml:"good"`
			Moderate int `yaml:"moderate"`
			Short    int `yaml:"short"`
		} `yaml:"thresholds"`
		Points struct {
			Long      int `yaml:"long"`
			Good      int `yaml:"good"`
			Moderate  int `yaml:"moderate"`
			Short     int `yaml:"short"`
			VeryShort int `yaml:"very_short"`
		} `yaml:"points"`
	} `yaml:"credit_age"`

	AgeScore struct {
		Thresholds struct {
			OptimalMin int `yaml:"optimal_min"`
			OptimalMax int `yaml:"optimal_max"`
			GoodMin    int `yaml:"good_min"`
			GoodMax    int `yaml:"good_max"`
		} `yaml:"thresholds"`
		Points struct {
			Optimal    int `yaml:"optimal"`
			Good       int `yaml:"good"`
			Acceptable int `yaml:"acceptable"`
		} `yaml:"points"`
	} `yaml:"age_score"`

	Decisions struct {
		ApprovedA int `yaml:"approved_a"`
		ApprovedB int `yaml:"approved_b"`
		Review    int `yaml:"review"`
	} `yaml:"decisions"`

	CreditLimitMultipliers struct {
		Approved float64 `yaml:"approved"`
		Review   float64 `yaml:"review"`
		Rejected float64 `yaml:"rejected"`
	} `yaml:"credit_limit_multipliers"`
}

func Load(path string) (*Rules, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rules Rules
	if err := yaml.Unmarshal(data, &rules); err != nil {
		return nil, err
	}
	return &rules, nil
}
