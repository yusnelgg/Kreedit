package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/yusnelgg/kreedit/config"
	"github.com/yusnelgg/kreedit/internal/api"
	"github.com/yusnelgg/kreedit/internal/scoring"
	"github.com/yusnelgg/kreedit/internal/storage"
)

func main() {
	godotenv.Load()

	cfg, err := config.Load("config/rules.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := storage.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	repo := storage.NewRepository(db)
	engine := scoring.NewEngine(cfg)
	handler := api.NewHandler(engine, repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	handler.RegisterRoutes(r)

	port := ":8080"
	fmt.Println("Kreedit corriendo en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
