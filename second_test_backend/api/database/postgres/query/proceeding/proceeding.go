package proceeding

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RowProceeding struct {
	Ciudad     string `json:"ciudad"`
	Actas      int    `json:"actas"`
	Porcentaje string `json:"porcentaje"`
}

func GetTotalProceedings(db *pgxpool.Pool) ([]RowProceeding, error) {
	ctx := context.Background()

	var totalActas int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM mesa").Scan(&totalActas); err != nil {
		fmt.Println(err)
	}

	ciudades := []string{"Santa Cruz", "La Paz", "Cochabamba", "Sucre", "Tarija", "Pando", "Beni", "Oruro", "Potosí"}

	actasPorCiudad := make(map[string]int)

	for _, ciudad := range ciudades {

		var cantidad int
		if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM mesa t1, recinto t2, ubicacion t3 WHERE t1.id_recinto = t2.id AND t2.id_ubicacion = t3.id AND t3.departamento = $1", ciudad).Scan(&cantidad); err != nil {
			fmt.Println(err)
		}

		actasPorCiudad[ciudad] = cantidad

	}

	porcentajePorCiudad := make(map[string]string)

	for _, ciudad := range ciudades {
		porcentaje := float64(actasPorCiudad[ciudad]) / float64(totalActas) * 100
		porcentajeS := fmt.Sprintf("%.2f", porcentaje)

		porcentajePorCiudad[ciudad] = strings.Replace(porcentajeS, ".", ",", -1)

	}

	result := make([]RowProceeding, 0)

	for _, ciudad := range ciudades {
		row := RowProceeding{ciudad, actasPorCiudad[ciudad], porcentajePorCiudad[ciudad]}
		result = append(result, row)
	}

	porcentaje := float64(totalActas) / float64(10) * 100
	porcentajeS := fmt.Sprintf("%.2f", porcentaje)
	rowActasValidas := RowProceeding{"Actas computadas", totalActas, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowActasValidas)

	rowActasTotal := RowProceeding{"Actas en total", 10, "100,00"}
	result = append(result, rowActasTotal)

	return result, nil

}

type Proceeding struct {
	Codigo          int    `json:"codigo"`
	Mesa            int    `json:"mesa"`
	Circunscripcion int    `json:"circunscripcion"`
	Departamento    string `json:"departamento"`
	Provincia       string `json:"provincia"`
	Municipio       string `json:"municipio"`
	Localidad       string `json:"localidad"`
	Recinto         string `json:"recinto"`
}

func GetUbicationProceeding(db *pgxpool.Pool, proceedingCode int) (Proceeding, error) {

	ctx := context.Background()

	var (
		codigo          int
		mesa            int
		circunscripcion int
		departamento    string
		provincia       string
		municipio       string
		localidad       string
		recinto         string
	)

	if err := db.QueryRow(ctx, "SELECT t1.id, t1.numero_de_mesa, t1.circunscripcion, t3.departamento, t3.provincia, t3.municipio, t3.localidad, t2.nombre FROM mesa t1, recinto t2, ubicacion t3 WHERE t1.id_recinto = t2.id AND t2.id_ubicacion = t3.id AND t1.id = $1", proceedingCode).Scan(&codigo, &mesa, &circunscripcion, &departamento, &provincia, &municipio, &localidad, &recinto); err != nil {
		fmt.Println(err)
	}

	result := Proceeding{codigo, mesa, circunscripcion, departamento, provincia, municipio, localidad, recinto}

	return result, nil

}

type ProceedingCode struct {
	Code int `json:"code"`
}

func ProceedingValidate(db *pgxpool.Pool) ([]ProceedingCode, error) {

	ctx := context.Background()

	var proceedingCode int
	result := make([]ProceedingCode, 0)

	rows, err := db.Query(ctx, "SELECT id FROM mesa")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&proceedingCode); err != nil {
			fmt.Println(err)
		} else {
			proceeding := ProceedingCode{proceedingCode}
			result = append(result, proceeding)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return result, nil

}

type RowProceedingVote struct {
	Codigo          int    `json:"codigo"`
	Recinto         string `json:"recinto"`
	Mesa            int    `json:"mesa"`
	Circunscripcion int    `json:"circunscripcion"`
	Dispositivo     string `json:"dispositivo"`
	Ubicacion       string `json:"ubicacion"`
	Departamento    string `json:"departamento"`
	Provincia       string `json:"provincia"`
	Municipio       string `json:"municipio"`
	Localidad       string `json:"localidad"`
	Partido         string `json:"partido"`
	VotosPresidente int    `json:"votosPresidente"`
	VotosDiputado   int    `json:"votosDiputado"`
}

func ProceedingVotes(db *pgxpool.Pool, proceedingCode int) ([]RowProceedingVote, error) {

	ctx := context.Background()

	var (
		recinto         string
		mesa            int
		circunscripcion int
		departamento    string
		provincia       string
		municipio       string
		localidad       string
		dispositivo     string
		ubicacion       string
	)

	if err := db.QueryRow(ctx, "SELECT t2.nombre, t1.numero_de_mesa, t1.circunscripcion, t3.departamento, t3.provincia, t3.municipio, t3.localidad, t1.dispositivo, t1.ubicacion FROM mesa t1, recinto t2, ubicacion t3 WHERE t1.id_recinto = t2.id AND t2.id_ubicacion = t3.id AND t1.id = $1", proceedingCode).Scan(&recinto, &mesa, &circunscripcion, &departamento, &provincia, &municipio, &localidad, &dispositivo, &ubicacion); err != nil {
		fmt.Println(err)
	}

	var votosValidosPresidente, votosBlancosPresidente, votosNulosPresidente, votosTotalPresidente int

	if err := db.QueryRow(ctx, "SELECT votos_validos, votos_blancos, votos_nulos, votos_total FROM votos_por_candidatura WHERE id_candidatura = 1 AND id_mesa = $1", proceedingCode).Scan(&votosValidosPresidente, &votosBlancosPresidente, &votosNulosPresidente, &votosTotalPresidente); err != nil {
		fmt.Println(err)
	}

	var votosValidosDiputado, votosBlancosDiputado, votosNulosDiputado, votosTotalDiputado int

	if err := db.QueryRow(ctx, "SELECT votos_validos, votos_blancos, votos_nulos, votos_total FROM votos_por_candidatura WHERE id_candidatura = 2 AND id_mesa = $1", proceedingCode).Scan(&votosValidosDiputado, &votosBlancosDiputado, &votosNulosDiputado, &votosTotalDiputado); err != nil {
		fmt.Println(err)
	}

	partidos := []string{"MAS-IPSP", "C.C.", "FPV", "MTS", "UCS", "21F", "PDC", "MNR", "PAN-BOL"}

	votosPorPartidoPresidente := make(map[string]int)

	for _, partido := range partidos {

		var cantidad int

		if err := db.QueryRow(ctx, "SELECT cantidad FROM votos_por_candidatura t1, votos_por_partido t2, partido t3 WHERE t1.id_mesa = t2.id_mesa AND t1.id_candidatura = t2.id_candidatura AND t2.id_partido = t3.id AND t1.id_mesa = $1 AND t2.id_mesa = $2 AND t1.id_candidatura = 1 AND t2.id_candidatura = 1 AND t3.id = (SELECT id FROM partido WHERE sigla = $3)", proceedingCode, proceedingCode, partido).Scan(&cantidad); err != nil {
			fmt.Println(err)
		} else {
			votosPorPartidoPresidente[partido] = cantidad
		}
	}

	votosPorPartidoDiputado := make(map[string]int)

	for _, partido := range partidos {

		var cantidad int

		if err := db.QueryRow(ctx, "SELECT cantidad FROM votos_por_candidatura t1, votos_por_partido t2, partido t3 WHERE t1.id_mesa = t2.id_mesa AND t1.id_candidatura = t2.id_candidatura AND t2.id_partido = t3.id AND t1.id_mesa = $1 AND t2.id_mesa = $2 AND t1.id_candidatura = 2 AND t2.id_candidatura = 2 AND t3.id = (SELECT id FROM partido WHERE sigla = $3)", proceedingCode, proceedingCode, partido).Scan(&cantidad); err != nil {
			fmt.Println(err)
		} else {
			votosPorPartidoDiputado[partido] = cantidad
		}

	}

	result := make([]RowProceedingVote, 0)

	for _, partido := range partidos {
		row := RowProceedingVote{proceedingCode, recinto, mesa, circunscripcion, dispositivo, ubicacion, departamento, provincia, municipio, localidad, partido, votosPorPartidoPresidente[partido], votosPorPartidoDiputado[partido]}
		result = append(result, row)
	}

	row := RowProceedingVote{proceedingCode, recinto, mesa, circunscripcion, dispositivo, ubicacion, departamento, provincia, municipio, localidad, "Votos válidos", votosValidosPresidente, votosValidosDiputado}
	result = append(result, row)

	row = RowProceedingVote{proceedingCode, recinto, mesa, circunscripcion, dispositivo, ubicacion, departamento, provincia, municipio, localidad, "Votos blancos", votosBlancosPresidente, votosBlancosDiputado}
	result = append(result, row)

	row = RowProceedingVote{proceedingCode, recinto, mesa, circunscripcion, dispositivo, ubicacion, departamento, provincia, municipio, localidad, "Votos nulos", votosNulosPresidente, votosNulosDiputado}
	result = append(result, row)

	row = RowProceedingVote{proceedingCode, recinto, mesa, circunscripcion, dispositivo, ubicacion, departamento, provincia, municipio, localidad, "Votos en total", votosTotalPresidente, votosTotalDiputado}
	result = append(result, row)

	return result, nil

}

func ExistsProceeding(db *pgxpool.Pool, proceeding int) bool {
	ctx := context.Background()

	var count int
	if err := db.QueryRow(ctx, "SELECT COUNT(*) FROM mesa WHERE id = $1", proceeding).Scan(&count); err != nil {
		fmt.Println(err)
	}

	return count != 0
}

func InsertProceeding(db *pgxpool.Pool, proceeding, numero, circunscripcion int, ubicacionRegistro, dispositivo string, ubicacion int) error {
	ctx := context.Background()

	commandTag, err := db.Exec(ctx, "INSERT INTO mesa(id, numero_de_mesa, circunscripcion, id_recinto, ubicacion, dispositivo) VALUES($1, $2, $3, $4, $5, $6)", proceeding, numero, circunscripcion, ubicacion, ubicacionRegistro, dispositivo)
	if err != nil {
		fmt.Println(err)
	}

	if commandTag.RowsAffected() != 1 {
		fmt.Println("Insert error")
	}

	return nil
}
