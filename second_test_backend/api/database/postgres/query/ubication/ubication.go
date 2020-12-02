package ubication

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ExistsUbication(db *pgxpool.Pool, departamento, provincia, municipio, localidad string) bool {
	ctx := context.Background()

	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM ubicacion WHERE departamento = $1 AND provincia = $2 AND municipio = $3 AND localidad = $4", departamento, provincia, municipio, localidad).Scan(&count); err != nil {
		fmt.Println(err)
	}

	return count != 0
}

func InsertUbication(db *pgxpool.Pool, departamento, provincia, municipio, localidad string) error {
	ctx := context.Background()

	var id int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM ubicacion").Scan(&id); err != nil {
		fmt.Println(err)
	}

	id++
	commandTag, err := db.Exec(ctx, "INSERT INTO ubicacion VALUES($1, $2, $3, $4, $5)", id, departamento, provincia, municipio, localidad)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil

}

func GetUbicationID(db *pgxpool.Pool, departamento, provincia, municipio, localidad string) int {
	ctx := context.Background()

	var id int
	if err := db.QueryRow(ctx, "SELECT id FROM ubicacion WHERE departamento = $1 AND provincia = $2 AND municipio = $3 AND localidad = $4", departamento, provincia, municipio, localidad).Scan(&id); err != nil {
		fmt.Println(err)
	}

	return id
}
