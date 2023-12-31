package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/veryordinary11/RSS-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Get the value of PORT
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	db := database.New(conn)
	apiCfg := &apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	// Create a new chi router
	router := chi.NewRouter()

	// Add a CORS middleware to the router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := createV1Router(apiCfg)

	router.Mount("/v1", v1Router)

	// Start the server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on port %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
