package main


import (
	"net/http"
	"log"
)

// apiConfig Struct
type apiConfig struct {}

func main () {
	const apiPath = "/api"
	const port = "8081"

	server := http.NewServeMux()

	server.HandleFunc(apiPath+"/healthz", healthCheck)
	
	cfg := apiConfig{}

	handler := cfg.CorsMiddleware(server)
	
	localSever := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(localSever.ListenAndServe())
}
