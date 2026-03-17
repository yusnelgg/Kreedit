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
	// 1. Crear el engine
	engine := scoring.NewEngine()

	// 2. Crear el handler
	handler := api.NewHandler(engine)

	// 3. Crear el router
	r := chi.NewRouter()

	// 4. Middlewares — logging y recovery automático ante panics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Registrar las rutas
	handler.RegisterRoutes(r)

	// 6. Arrancar el servidor
	port := ":8080"
	fmt.Println("Kreedit corriendo en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
