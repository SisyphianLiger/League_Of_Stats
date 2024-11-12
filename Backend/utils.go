package main

import (
	"database/sql"
	"log"
	"os"
)

func environmentVarExists(key string) string {
	payload, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("PLATFORM environment variable not set for %s", key)
	}
	return payload
}

func openDB(dbType string, dbURL string) *sql.DB {
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		log.Printf("Unable to Open DB, check ports and connecion string")
		return nil
	}
	return db
}
