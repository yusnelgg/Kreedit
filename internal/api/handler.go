package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yusnelgg/kreedit/internal/domain"
	"github.com/yusnelgg/kreedit/internal/scoring"
)

type Handler struct {
	engine *scoring.Engine
}

func NewHandler(engine *scoring.Engine) *Handler {
	return &Handler{engine: engine}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Post("/api/v1/score", h.Score)
	r.Get("/api/v1/health", h.Health)
}

func (h *Handler) Score(w http.ResponseWriter, r *http.Request) {

	var app domain.CreditApplication
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validate(app); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result := h.engine.Score(app)

	respondJSON(w, http.StatusOK, result)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "v1.0.0",
	})
}

func validate(app domain.CreditApplication) error {
	if app.ApplicantID == "" {
		return fmt.Errorf("applicant_id is required")
	}
	if app.MonthlyIncome <= 0 {
		return fmt.Errorf("monthly_income must be greater than 0")
	}
	if app.Age < 18 || app.Age > 100 {
		return fmt.Errorf("age must be between 18 and 100")
	}
	if app.RequestedAmount <= 0 {
		return fmt.Errorf("requested_amount must be greater than 0")
	}
	if app.MissedPayments < 0 {
		return fmt.Errorf("missed_payments cannot be negative")
	}
	return nil
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
