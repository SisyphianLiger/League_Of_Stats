-- +goose Up
CREATE TABLE PlayerCards(
	Name TEXT NOT NULL UNIQUE,
	Age smallint,
	Birthday date,
	Country TEXT, 
	Team TEXT, 
	Position TEXT,
	Worlds_Titles smallint
);

INSERT INTO PlayerCards (Name, Age, Birthday, Country, Team, Position, Worlds_Titles)
VALUES 
    ('Faker', 27, '1996-05-07', 'South Korea', 'T1', 'Mid', 4),
    ('ShowMaker', 23, '2000-07-22', 'South Korea', 'Dplus KIA', 'Mid', 1);

-- +goose Down
DROP TABLE PlayerCards;
