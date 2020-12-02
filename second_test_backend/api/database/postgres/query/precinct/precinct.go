package precinct

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ExistsPrecinct(db *pgxpool.Pool, recinto string, ubicacion int) bool {
	ctx := context.Background()

	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM recinto WHERE nombre = $1 AND id_ubicacion = $2", recinto, ubicacion).Scan(&count); err != nil {
		fmt.Println(err)
	}

	return count != 0
}

func InsertPrecinct(db *pgxpool.Pool, recinto string, ubicacion int) error {
	ctx := context.Background()

	var id int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM recinto").Scan(&id); err != nil {
		fmt.Println(err)
	}

	id++
	commandTag, err := db.Exec(ctx, "INSERT INTO recinto VALUES($1, $2, $3)", id, recinto, ubicacion)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil

}

func GetPrecinctID(db *pgxpool.Pool, recinto string, ubicacion int) int {
	ctx := context.Background()

	var id int
	if err := db.QueryRow(ctx, "SELECT id FROM recinto WHERE nombre = $1 AND id_ubicacion = $2", recinto, ubicacion).Scan(&id); err != nil {
		fmt.Println(err)
	}

	return id
}
