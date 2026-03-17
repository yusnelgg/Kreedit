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

// RegisterRoutes registra todas las rutas de Kreedit
func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Post("/api/v1/score", h.Score)
	r.Get("/api/v1/health", h.Health)
}

// Score recibe una solicitud y devuelve el resultado
func (h *Handler) Score(w http.ResponseWriter, r *http.Request) {

	// 1. Leer el body
	var app domain.CreditApplication
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// 2. Validar los datos
	if err := validate(app); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// 3. Llamar al engine
	result := h.engine.Score(app)

	// 4. Devolver el resultado
	respondJSON(w, http.StatusOK, result)
}

// Health es un endpoint para verificar que el servidor está vivo
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "v1.0.0",
	})
}

// validate verifica que los datos del request tengan sentido
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

// respondJSON escribe una respuesta JSON
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError escribe un error en formato JSON
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
