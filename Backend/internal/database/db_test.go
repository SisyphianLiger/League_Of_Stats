package database

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/SisyphianLiger/League_Of_Stats/internal/utils"
)

type Context struct {
	ctx context.Context
}


func TestPreLoadEnvFile(t *testing.T) {

	dirErr := os.Chdir("../../../")
	if dirErr != nil {
		t.Logf("If dirErr does not load here we are in trouble %v", dirErr)
	}
	utils.LoadEnvFile(".env")
}


func TestFindAllPlayerCardsSuccess(t *testing.T) {

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


func TestPutPlayerCardToInMemStruct(t *testing.T) {
	type PlayerCards struct {
		Name         string
		Age          uint16
		Birthday     time.Time
		Country      string
		Team         string
		Position     string 
		WorldsTitles uint16
	}

	var players []PlayerCards
	
	dbURL := utils.EnvironmentVarExists("DB_URL")
	db := utils.OpenDB("postgres", dbURL)

	q := Queries{
		db: db,
	}
	// Make a note about Background for testing
	ctx := context.Background()
	p, err := q.FindAllPlayerCards(ctx)

	if err != nil {
		t.Fatalf("The Database could not be queried: %v", err)
	}

	for _, ps := range p {
		add_player := PlayerCards {
			Name: ps.Name,
			Age: uint16(ps.Age.Int16),
			Birthday: ps.Birthday.Time,
			Country: ps.Country.String,
			Team: ps.Team.String,
			Position: ps.Position.String,
			WorldsTitles: uint16(ps.WorldsTitles.Int16),
		}

		players = append(players, add_player)
	}

	if len(players) != 2 {
		t.Fatalf("The Query should have 2 in DB, found %d", len(players))
	}
}
