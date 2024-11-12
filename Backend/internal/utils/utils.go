package utils

import (
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func LoadEnvFile(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("File not Found: %v", err)
	}
}

func EnvironmentVarExists(key string) string {
	payload, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("PLATFORM environment variable not set for %s", key)
	}
	return payload
}

func OpenDB(dbType string, dbURL string) *sql.DB {
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		log.Printf("Unable to Open DB, check ports and connecion string: %v", err)
		return nil
	}
	return db
}
