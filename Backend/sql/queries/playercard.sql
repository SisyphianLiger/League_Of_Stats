-- name: CreatePlayerCard :one
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
) RETURNING *;

-- name: FindAllPlayerCards :many
SELECT * FROM PlayerCards;
