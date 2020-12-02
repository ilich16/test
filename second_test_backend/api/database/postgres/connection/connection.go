package connection

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func OpenConnection() error {
	ctx := context.Background()
	var err error
	DB, err = pgxpool.Connect(ctx, "postgresql://username:password@localhost:5432/secondtestv2")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CloseConnection() error {
	DB.Close()
	return nil
}
