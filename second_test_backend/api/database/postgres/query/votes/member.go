package votes

import (
	"api/controller/functionality/visionai"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertVotesMember(db *pgxpool.Pool, proceeding visionai.Proceeding) error {
	ctx := context.Background()

	commandTag, err := db.Exec(ctx, "INSERT INTO votos_por_candidatura VALUES($1, 2, $2, $3, $4, $5)", proceeding.ID, proceeding.DiputadoValido, proceeding.DiputadoBlanco, proceeding.DiputadoNulo, proceeding.VotosTotal)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMAS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 1, $2)", proceeding.ID, proceeding.DiputadoMAS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMAS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagCC, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 2, $2)", proceeding.ID, proceeding.DiputadoCC)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagCC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagFPV, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 3, $2)", proceeding.ID, proceeding.DiputadoFPV)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagFPV.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMTS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 4, $2)", proceeding.ID, proceeding.DiputadoMTS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMTS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagUCS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 5, $2)", proceeding.ID, proceeding.DiputadoUCS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagUCS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagVeF, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 6, $2)", proceeding.ID, proceeding.DiputadoVeF)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagVeF.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPDC, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 7, $2)", proceeding.ID, proceeding.DiputadoPDC)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPDC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMNR, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 8, $2)", proceeding.ID, proceeding.DiputadoMNR)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMNR.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPAN, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 2, 9, $2)", proceeding.ID, proceeding.DiputadoPAN)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPAN.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil

}

func UpdateVotesMember(db *pgxpool.Pool, proceeding visionai.Proceeding) error {
	ctx := context.Background()

	commandTag, err := db.Exec(ctx, "UPDATE votos_por_candidatura SET votos_validos = $1, votos_blancos = $2, votos_nulos = $3, votos_total = $4 WHERE id_mesa = $5 AND id_candidatura = 2", proceeding.DiputadoValido, proceeding.DiputadoBlanco, proceeding.DiputadoNulo, proceeding.VotosTotal, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMAS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 1", proceeding.DiputadoMAS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMAS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagCC, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 2", proceeding.DiputadoCC, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagCC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagFPV, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 3", proceeding.DiputadoFPV, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagFPV.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMTS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 4", proceeding.DiputadoMTS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMTS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagUCS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 5", proceeding.DiputadoUCS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagUCS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagVeF, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 6", proceeding.DiputadoVeF, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagVeF.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPDC, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 7", proceeding.DiputadoPDC, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPDC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMNR, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 8", proceeding.DiputadoMNR, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMNR.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPAN, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 2 AND id_partido = 9", proceeding.DiputadoPAN, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPAN.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil
}
