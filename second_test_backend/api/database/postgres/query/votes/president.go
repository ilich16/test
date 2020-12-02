package votes

import (
	"api/controller/functionality/visionai"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertVotesPresident(db *pgxpool.Pool, proceeding visionai.Proceeding) error {
	ctx := context.Background()

	commandTag, err := db.Exec(ctx, "INSERT INTO votos_por_candidatura VALUES($1, 1, $2, $3, $4, $5)", proceeding.ID, proceeding.PresidenteValido, proceeding.PresidenteBlanco, proceeding.PresidenteNulo, proceeding.VotosTotal)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMAS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 1, $2)", proceeding.ID, proceeding.PresidenteMAS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMAS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagCC, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 2, $2)", proceeding.ID, proceeding.PresidenteCC)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagCC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagFPV, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 3, $2)", proceeding.ID, proceeding.PresidenteFPV)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagFPV.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMTS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 4, $2)", proceeding.ID, proceeding.PresidenteMTS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMTS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagUCS, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 5, $2)", proceeding.ID, proceeding.PresidenteUCS)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagUCS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagVeF, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 6, $2)", proceeding.ID, proceeding.PresidenteVeF)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagVeF.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPDC, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 7, $2)", proceeding.ID, proceeding.PresidentePDC)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPDC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMNR, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 8, $2)", proceeding.ID, proceeding.PresidenteMNR)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMNR.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPAN, err := db.Exec(ctx, "INSERT INTO votos_por_partido VALUES($1, 1, 9, $2)", proceeding.ID, proceeding.PresidentePAN)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPAN.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil

}

func UpdateVotesPresident(db *pgxpool.Pool, proceeding visionai.Proceeding) error {
	ctx := context.Background()

	commandTag, err := db.Exec(ctx, "UPDATE votos_por_candidatura SET votos_validos = $1, votos_blancos = $2, votos_nulos = $3, votos_total = $4 WHERE id_mesa = $5 AND id_candidatura = 1", proceeding.PresidenteValido, proceeding.PresidenteBlanco, proceeding.PresidenteNulo, proceeding.VotosTotal, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMAS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 1", proceeding.PresidenteMAS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMAS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagCC, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 2", proceeding.PresidenteCC, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagCC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagFPV, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 3", proceeding.PresidenteFPV, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagFPV.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMTS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 4", proceeding.PresidenteMTS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMTS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagUCS, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 5", proceeding.PresidenteUCS, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagUCS.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagVeF, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 6", proceeding.PresidenteVeF, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagVeF.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPDC, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 7", proceeding.PresidentePDC, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPDC.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagMNR, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 8", proceeding.PresidenteMNR, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagMNR.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	commandTagPAN, err := db.Exec(ctx, "UPDATE votos_por_partido SET cantidad = $1 WHERE id_mesa = $2 AND id_candidatura = 1 AND id_partido = 9", proceeding.PresidentePAN, proceeding.ID)
	if err != nil {
		fmt.Println(err)
	}

	if commandTagPAN.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil
}
