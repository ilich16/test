package president

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Precinct struct {
	Nombre string `json:"nombre"`
}

func GetPrecincts(db *pgxpool.Pool, city string) ([]Precinct, error) {
	ctx := context.Background()

	rows, err := db.Query(ctx, "SELECT nombre FROM recinto t1, ubicacion t2 WHERE t1.id_ubicacion = t2.id AND t2.departamento = $1", city)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := make([]Precinct, 0)
	var precinctName string

	for rows.Next() {
		if err := rows.Scan(&precinctName); err != nil {
			fmt.Println(err)
		} else {
			precinct := Precinct{precinctName}
			result = append(result, precinct)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return result, nil
}

type ReportResultsFromRecinct struct {
	id         int    `json:"id"`
	Candidato  string `json:"candidato"`
	Partido    string `json:"partido"`
	Votos      int    `json:"votos"`
	Porcentaje string `json:"porcentaje"`
}

func GetResultsFromRecinct(db *pgxpool.Pool, recinct string) ([]ReportResultsFromRecinct, error) {
	ctx := context.Background()

	rowsFirstQuery, err := db.Query(ctx, "SELECT t1.id, votos_validos, votos_blancos, votos_nulos, votos_total FROM recinto t1, mesa t2, candidatura t3, votos_por_candidatura t4 WHERE t1.id = t2.id_recinto AND t2.id = t4.id_mesa AND t3.id = t4.id_candidatura AND t3.id = 1 AND t4.id_candidatura = 1 AND t1.nombre = $1", recinct)
	if err != nil {
		fmt.Println(err)
	}

	defer rowsFirstQuery.Close()

	var id, totalVotosValidos, totalVotosBlancos, totalVotosNulos, totalVotosTotal int

	for rowsFirstQuery.Next() {
		var votosValidos, votosBlancos, votosNulos, votosTotal int
		if err := rowsFirstQuery.Scan(&id, &votosValidos, &votosBlancos, &votosNulos, &votosTotal); err != nil {
			fmt.Println(err)
		} else {
			totalVotosValidos += votosValidos
			totalVotosBlancos += votosBlancos
			totalVotosNulos += votosNulos
			totalVotosTotal += votosTotal
		}
	}

	if err := rowsFirstQuery.Err(); err != nil {
		fmt.Println(err)
	}

	partidos := []string{"MAS-IPSP", "C.C.", "FPV", "MTS", "UCS", "21F", "PDC", "MNR", "PAN-BOL"}

	votosPorPartido := make(map[string]int)

	for _, partido := range partidos {

		var cantidadTotal, cantidad int

		rows, err := db.Query(ctx, "SELECT cantidad FROM recinto t1, mesa t2, candidatura t3, votos_por_candidatura t4, votos_por_partido t5 WHERE t1.id = t2.id_recinto AND t2.id = t4.id_mesa AND t3.id = t4.id_candidatura AND t4.id_mesa = t5.id_mesa AND t4.id_candidatura = t5.id_candidatura AND t3.id = 1 AND t4.id_candidatura = 1 AND t5.id_candidatura = 1 AND t1.nombre = $1 AND t5.id_partido = (SELECT id FROM partido WHERE sigla = $2)", recinct, partido)
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			if err := rows.Scan(&cantidad); err != nil {
				fmt.Println(err)
			} else {
				cantidadTotal += cantidad
			}
		}

		rows.Close()

		votosPorPartido[partido] = cantidadTotal

	}

	porcentajePorPartido := make(map[string]string)

	for _, partido := range partidos {
		porcentaje := float64(votosPorPartido[partido]) / float64(totalVotosValidos) * 100
		porcentajeS := fmt.Sprintf("%.2f", porcentaje)

		porcentajePorPartido[partido] = strings.Replace(porcentajeS, ".", ",", -1)

	}

	candidatosPorPartido := make(map[string]string)

	for _, partido := range partidos {

		candidatos := make(map[string]string)

		var nombres, apellidos, descripcion string

		rows, err := db.Query(ctx, "SELECT nombres, apellidos, descripcion FROM persona t1, candidato t2, partido t3 WHERE t1.id = t2.id_persona AND t1.id_partido = t3.id AND t3.sigla = $1 AND (t2.descripcion = 'Presidente' OR t2.descripcion = 'Vicepresidente')", partido)
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			if err := rows.Scan(&nombres, &apellidos, &descripcion); err != nil {
				fmt.Println(err)
			} else {
				candidatos[descripcion] = nombres + " " + apellidos
			}
		}

		rows.Close()

		candidatosPorPartido[partido] = candidatos["Presidente"] + " - " + candidatos["Vicepresidente"]

	}

	result := make([]ReportResultsFromRecinct, 0)

	for _, partido := range partidos {
		row := ReportResultsFromRecinct{id, candidatosPorPartido[partido], nombrePartidos[partido], votosPorPartido[partido], porcentajePorPartido[partido]}
		result = append(result, row)
	}

	porcentaje := float64(totalVotosValidos) / float64(totalVotosTotal) * 100
	porcentajeS := fmt.Sprintf("%.2f", porcentaje)
	rowVotosValidos := ReportResultsFromRecinct{id, "Votos válidos", "", totalVotosValidos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosValidos)

	porcentaje = float64(totalVotosBlancos) / float64(totalVotosTotal) * 100
	porcentajeS = fmt.Sprintf("%.2f", porcentaje)
	rowVotosBlancos := ReportResultsFromRecinct{id, "Votos blancos", "", totalVotosBlancos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosBlancos)

	porcentaje = float64(totalVotosNulos) / float64(totalVotosTotal) * 100
	porcentajeS = fmt.Sprintf("%.2f", porcentaje)
	rowVotosNulos := ReportResultsFromRecinct{id, "Votos nulos", "", totalVotosNulos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosNulos)

	rowVotosTotal := ReportResultsFromRecinct{id, "Votos en total", "", totalVotosTotal, "100,00"}
	result = append(result, rowVotosTotal)

	return result, nil
}
