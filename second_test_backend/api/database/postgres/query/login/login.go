package login

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func VerifyUser(db *pgxpool.Pool, username, password string) bool {
	ctx := context.Background()

	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM usuario WHERE nombre_de_usuario = $1", username).Scan(&count); err != nil {
		fmt.Println(err)
	}

	if count == 0 {
		return false
	}

	var result int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM usuario WHERE nombre_de_usuario = $1 AND contrasena = $2;", username, password).Scan(&result); err != nil {
		fmt.Println(err)
	}

	return result != 0

}
