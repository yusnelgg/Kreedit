package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yusnelgg/kreedit/internal/api"
	"github.com/yusnelgg/kreedit/internal/scoring"
)

func main() {
	engine := scoring.NewEngine()

	handler := api.NewHandler(engine)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler.RegisterRoutes(r)

	port := ":8080"
	fmt.Println("Kreedit corriendo en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
