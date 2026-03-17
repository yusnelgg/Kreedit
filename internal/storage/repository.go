package storage

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/yusnelgg/kreedit/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (r *Repository) Save(result domain.ScoringResult) error {
	reasons, err := json.Marshal(result.Reasons)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		INSERT INTO scoring_results (
			application_id,
			applicant_id,
			score,
			decision,
			risk_tier,
			credit_limit,
			reasons,
			model_version,
			processed_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		result.ApplicationID,
		result.ApplicantID,
		result.Score,
		string(result.Decision),
		string(result.RiskTier),
		result.CreditLimit,
		string(reasons),
		result.ModelVersion,
		result.ProcessedAt,
	)
	return err
}

func (r *Repository) FindByApplicant(applicantID string) ([]domain.ScoringResult, error) {
	rows, err := r.db.Query(`
		SELECT application_id, applicant_id, score, decision,
		       risk_tier, credit_limit, reasons, model_version, processed_at
		FROM scoring_results
		WHERE applicant_id = $1
		ORDER BY processed_at DESC`,
		applicantID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.ScoringResult
	for rows.Next() {
		var r domain.ScoringResult
		var decision, riskTier string
		var reasons []byte

		err := rows.Scan(
			&r.ApplicationID, &r.ApplicantID, &r.Score,
			&decision, &riskTier, &r.CreditLimit,
			&reasons, &r.ModelVersion, &r.ProcessedAt,
		)
		if err != nil {
			return nil, err
		}

		r.Decision = domain.Decision(decision)
		r.RiskTier = domain.RiskTier(riskTier)
		json.Unmarshal(reasons, &r.Reasons)
		results = append(results, r)
	}
	return results, nil
}
