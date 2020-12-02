package president

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetResultsCity(db *pgxpool.Pool, departamento string) ([]ResPresidencial, error) {
	ctx := context.Background()

	rowsFirstQuery, err := db.Query(ctx, "SELECT votos_validos, votos_blancos, votos_nulos, votos_total FROM votos_por_candidatura t1, mesa t2, recinto t3, ubicacion t4 WHERE t1.id_mesa = t2.id AND t2.id_recinto = t3.id AND t3.id_ubicacion = t4.id AND t1.id_candidatura = 1 AND t4.departamento = $1", departamento)
	if err != nil {
		fmt.Println(err)
	}

	defer rowsFirstQuery.Close()

	var totalVotosValidos, totalVotosBlancos, totalVotosNulos, totalVotosTotal int

	for rowsFirstQuery.Next() {
		var votosValidos, votosBlancos, votosNulos, votosTotal int
		if err := rowsFirstQuery.Scan(&votosValidos, &votosBlancos, &votosNulos, &votosTotal); err != nil {
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

		rows, err := db.Query(ctx, "SELECT cantidad FROM votos_por_candidatura t1, mesa t2, recinto t3, ubicacion t4, votos_por_partido t5 WHERE t1.id_mesa = t2.id AND t2.id_recinto = t3.id AND t3.id_ubicacion = t4.id AND t5.id_mesa = t1.id_mesa AND t5.id_candidatura = 1 AND t1.id_candidatura = 1 AND t5.id_partido = (SELECT id FROM partido WHERE sigla = $1) AND t4.departamento = $2", partido, departamento)
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

	result := make([]ResPresidencial, 0)

	for _, partido := range partidos {
		row := ResPresidencial{candidatosPorPartido[partido], nombrePartidos[partido], votosPorPartido[partido], porcentajePorPartido[partido]}
		result = append(result, row)
	}

	porcentaje := float64(totalVotosValidos) / float64(totalVotosTotal) * 100
	porcentajeS := fmt.Sprintf("%.2f", porcentaje)
	rowVotosValidos := ResPresidencial{"Votos válidos", "", totalVotosValidos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosValidos)

	porcentaje = float64(totalVotosBlancos) / float64(totalVotosTotal) * 100
	porcentajeS = fmt.Sprintf("%.2f", porcentaje)
	rowVotosBlancos := ResPresidencial{"Votos blancos", "", totalVotosBlancos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosBlancos)

	porcentaje = float64(totalVotosNulos) / float64(totalVotosTotal) * 100
	porcentajeS = fmt.Sprintf("%.2f", porcentaje)
	rowVotosNulos := ResPresidencial{"Votos nulos", "", totalVotosNulos, strings.Replace(porcentajeS, ".", ",", -1)}
	result = append(result, rowVotosNulos)

	rowVotosTotal := ResPresidencial{"Votos en total", "", totalVotosTotal, "100,00"}
	result = append(result, rowVotosTotal)

	return result, nil
}
