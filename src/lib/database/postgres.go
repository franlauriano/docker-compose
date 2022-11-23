package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func OpenPostgres(user, password, dbname, host string, port int) {
	strConn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable timezone=UTC", user, password, dbname, host, port)
	dbInstance, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatalf("Could not open database connection. err = %v", err)
	}

	log.Println("Database connection opened")
	db = dbInstance
}

type contextKey string

const dbContextKey = contextKey("postgres.db")

func Tx(ctx context.Context) *sql.Tx {
	return ctx.Value(dbContextKey).(*sql.Tx)
}

func DB() *sql.DB {
	return db
}

func Begin() (context.Context, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, dbContextKey, tx), nil
}

func Commit(ctx context.Context) error {
	return Tx(ctx).Commit()
}

func Rollback(ctx context.Context) error {
	return Tx(ctx).Rollback()
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatalf("Error on close database. error = %v", err)
	}
	log.Println("Database connection closed")
}