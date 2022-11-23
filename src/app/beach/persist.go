package beach

import (
	"log"
	"time"

	"github.com/franlauriano/docker-compose/lib/database"
)

var persist persistent = &postgres{}

type persistent interface {
	list() ([]Beach, error)
	create(*Beach) error
}

type postgres struct {
}

func (psql *postgres) list() ([]Beach, error) {
	command := `
	 SELECT id, created_at, updated_at, name, state, ranking
	 FROM beaches
	 WHERE deleted_at IS NULL
	 ORDER BY ranking ASC`

	ctx, err := database.Begin()
	if err != nil {
		log.Printf("Could not create tx: %s", err)
	}

	db := database.Tx(ctx)
	stmt, err := db.PrepareContext(ctx, command)
	if err != nil {
		log.Printf("Error on prepare statment to find. Error = %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("Error on query to search. Error = %v", err)
		return nil, err
	}

	beaches := make([]Beach, 0)
	for rows.Next() {
		var beach Beach
		if err := rows.Scan(
			&beach.ID,
			&beach.CreatedAt,
			&beach.UpdatedAt,
			&beach.Name,
			&beach.State,
			&beach.Ranking,
		); err != nil {
			log.Printf("Error on read rows to find. Error = %v", err)
			return nil, err
		}
		beaches = append(beaches, beach)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration to find. Error = %v", err)
		return nil, err
	}

	return beaches, nil
}

func (psql *postgres) create(beach *Beach) error {
	command := `
	 INSERT INTO beaches (created_at, updated_at, name, state, ranking)
	 VALUES ($1, $2, $3, $4, $5)
	 RETURNING id`

	ctx, err := database.Begin()
	if err != nil {
		log.Printf("Could not create tx: %s", err)
	}

	db := database.Tx(ctx)
	stmt, err := db.PrepareContext(ctx, command)
	if err != nil {
		log.Printf("Error on prepare statment to create. Error = %v", err)
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, time.Now(), time.Now(), beach.Name, beach.State, beach.Ranking).Scan(&beach.ID)
	if err != nil {
		log.Printf("Error on execute command to create. Error = %v", err)
		return err
	}

	return database.Commit(ctx)
}
