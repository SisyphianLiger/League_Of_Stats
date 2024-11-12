// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: playercard.sql

package database

import (
	"context"
	"database/sql"
)

const createPlayerCard = `-- name: CreatePlayerCard :one
INSERT INTO PlayerCards (
    Name,
    Age,
    Birthday,
    Country,
    Team,
    Position,
    Worlds_Titles
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING name, age, birthday, country, team, position, worlds_titles
`

type CreatePlayerCardParams struct {
	Name         string
	Age          sql.NullInt16
	Birthday     sql.NullTime
	Country      sql.NullString
	Team         sql.NullString
	Position     sql.NullString
	WorldsTitles sql.NullInt16
}

func (q *Queries) CreatePlayerCard(ctx context.Context, arg CreatePlayerCardParams) (Playercard, error) {
	row := q.db.QueryRowContext(ctx, createPlayerCard,
		arg.Name,
		arg.Age,
		arg.Birthday,
		arg.Country,
		arg.Team,
		arg.Position,
		arg.WorldsTitles,
	)
	var i Playercard
	err := row.Scan(
		&i.Name,
		&i.Age,
		&i.Birthday,
		&i.Country,
		&i.Team,
		&i.Position,
		&i.WorldsTitles,
	)
	return i, err
}

const findAllPlayerCards = `-- name: FindAllPlayerCards :many
SELECT name, age, birthday, country, team, position, worlds_titles FROM PlayerCards
`

func (q *Queries) FindAllPlayerCards(ctx context.Context) ([]Playercard, error) {
	rows, err := q.db.QueryContext(ctx, findAllPlayerCards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playercard
	for rows.Next() {
		var i Playercard
		if err := rows.Scan(
			&i.Name,
			&i.Age,
			&i.Birthday,
			&i.Country,
			&i.Team,
			&i.Position,
			&i.WorldsTitles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
