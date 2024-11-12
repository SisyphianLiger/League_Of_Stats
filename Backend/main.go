package main

import (
	"log"
	"net/http"

	"github.com/SisyphianLiger/League_Of_Stats/internal/database"
	"github.com/SisyphianLiger/League_Of_Stats/internal/utils"
)

// apiConfig Struct
type apiConfig struct {
	dbq *database.Queries
}

func main() {
	const apiPath = "/api"
	const port = "8081"

	utils.LoadEnvFile(".env")
	server := http.NewServeMux()

	server.HandleFunc(apiPath+"/healthz", healthCheck)

	dbURL := utils.EnvironmentVarExists("DB_URL")

	// Make DB Connection extracted
	db := utils.OpenDB("postgres", dbURL)
	dbQueries := database.New(db)

	cfg := apiConfig{
		dbq: dbQueries,
	}

	handler := cfg.CorsMiddleware(server)

	localSever := &http.Server{
		Handler:     handler,
		Addr:        ":" + port,
		ReadTimeout: 20,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(localSever.ListenAndServe())
}
