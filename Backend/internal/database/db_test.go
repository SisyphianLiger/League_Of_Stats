package database

import (
	"context"
	"testing"
	"os"
	"github.com/SisyphianLiger/League_Of_Stats/internal/utils"
)

type Context struct {
	ctx context.Context
}

func TestFindAllPlayerCardsSuccess(t *testing.T) {

	dirErr := os.Chdir("../../../")
	if dirErr != nil {
		t.Logf("If dirErr does not load here we are in trouble %v", dirErr)
	}
	
	utils.LoadEnvFile(".env")
	dbURL := utils.EnvironmentVarExists("DB_URL")
	db := utils.OpenDB("postgres", dbURL)

	q := Queries{
		db: db,
	}
	// Make a note about Background for testing
	ctx := context.Background()
	_, err := q.FindAllPlayerCards(ctx)

	if err != nil {
		t.Fatalf("The Database could not be queried: %v", err)
	}

}
