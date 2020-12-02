package report

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Report struct {
	Candidato  string `json:"candidato"`
	Partido    string `json:"partido"`
	Votos      int    `json:"votos"`
	Porcentaje string `json:"porcentaje"`
}

func ReportPresidentCountry(data []Report) bytes.Buffer {
	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "ELECCIONES GENERALES 2019")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "D1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A2", "RESULTADOS PRESIDENCIALES EN TODO EL PAÍS")

	subtitleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "D2", subtitleStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
	}

	// Set col width
	if err := f.SetColWidth("Sheet1", "A", "A", 60); err != nil {
		fmt.Println(err)
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 40); err != nil {
		fmt.Println(err)
	}
	// Set row height
	if err := f.SetRowHeight("Sheet1", 1, 32); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A3", "Candidato a presidente y vicepresidente")
	f.SetCellValue("Sheet1", "B3", "Partido político")
	f.SetCellValue("Sheet1", "C3", "Votos")
	f.SetCellValue("Sheet1", "D3", "%")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowCells := []string{"A3", "B3", "C3", "D3"}
	for _, cell := range thirdRowCells {
		if cell != "A3" {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowStyle); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowFirstColStyle); err != nil {
				fmt.Println(err)
			}
		}
	}

	for i, row := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Candidato)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.Partido)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), row.Votos)
		porcentaje := strings.Replace(row.Porcentaje, ",", ".", -1)
		porcentajeF, err := strconv.ParseFloat(porcentaje, 64)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+4), porcentajeF)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A16", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowSecondColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B16", fourthRowSecondColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D4", "D16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A13", "B13"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A14", "B14"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A15", "B15"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A16", "B16"); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D13", "D16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A17", GetTime())

	if err := f.MergeCell("Sheet1", "A17", "D17"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A17", "D17", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}

func ReportPresidentCity(data []Report, city string) bytes.Buffer {
	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "ELECCIONES GENERALES 2019")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "D1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A2", "RESULTADOS PRESIDENCIALES EN LA CIUDAD DE "+strings.ToUpper(city))

	subtitleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "D2", subtitleStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
	}

	// Set col width
	if err := f.SetColWidth("Sheet1", "A", "A", 60); err != nil {
		fmt.Println(err)
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 40); err != nil {
		fmt.Println(err)
	}
	// Set row height
	if err := f.SetRowHeight("Sheet1", 1, 32); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A3", "Candidato a presidente y vicepresidente")
	f.SetCellValue("Sheet1", "B3", "Partido político")
	f.SetCellValue("Sheet1", "C3", "Votos")
	f.SetCellValue("Sheet1", "D3", "%")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowCells := []string{"A3", "B3", "C3", "D3"}
	for _, cell := range thirdRowCells {
		if cell != "A3" {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowStyle); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowFirstColStyle); err != nil {
				fmt.Println(err)
			}
		}
	}

	for i, row := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Candidato)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.Partido)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), row.Votos)
		porcentaje := strings.Replace(row.Porcentaje, ",", ".", -1)
		porcentajeF, err := strconv.ParseFloat(porcentaje, 64)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+4), porcentajeF)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A16", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowSecondColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B16", fourthRowSecondColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D4", "D16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A13", "B13"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A14", "B14"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A15", "B15"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A16", "B16"); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D13", "D16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A17", GetTime())

	if err := f.MergeCell("Sheet1", "A17", "D17"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A17", "D17", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}

func ReportPresidentDistrict(data []Report, district string) bytes.Buffer {
	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "ELECCIONES GENERALES 2019")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "D1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A2", "RESULTADOS PRESIDENCIALES EN LA CIRCUNSCRIPCIÓN "+district)

	subtitleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "D2", subtitleStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
	}

	// Set col width
	if err := f.SetColWidth("Sheet1", "A", "A", 60); err != nil {
		fmt.Println(err)
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 40); err != nil {
		fmt.Println(err)
	}
	// Set row height
	if err := f.SetRowHeight("Sheet1", 1, 32); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A3", "Candidato a presidente y vicepresidente")
	f.SetCellValue("Sheet1", "B3", "Partido político")
	f.SetCellValue("Sheet1", "C3", "Votos")
	f.SetCellValue("Sheet1", "D3", "%")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowCells := []string{"A3", "B3", "C3", "D3"}
	for _, cell := range thirdRowCells {
		if cell != "A3" {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowStyle); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowFirstColStyle); err != nil {
				fmt.Println(err)
			}
		}
	}

	for i, row := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Candidato)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.Partido)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), row.Votos)
		porcentaje := strings.Replace(row.Porcentaje, ",", ".", -1)
		porcentajeF, err := strconv.ParseFloat(porcentaje, 64)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+4), porcentajeF)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A16", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowSecondColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B16", fourthRowSecondColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D4", "D16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A13", "B13"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A14", "B14"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A15", "B15"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A16", "B16"); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D13", "D16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A17", GetTime())

	if err := f.MergeCell("Sheet1", "A17", "D17"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A17", "D17", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}

func ReportPresidentPrecinct(data []Report, precinct string) bytes.Buffer {
	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "ELECCIONES GENERALES 2019")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "D1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A2", "RESULTADOS PRESIDENCIALES EN EL RECINTO: "+strings.ToUpper(precinct))

	subtitleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "D2", subtitleStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
	}

	// Set col width
	if err := f.SetColWidth("Sheet1", "A", "A", 60); err != nil {
		fmt.Println(err)
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 40); err != nil {
		fmt.Println(err)
	}
	// Set row height
	if err := f.SetRowHeight("Sheet1", 1, 32); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A3", "Candidato a presidente y vicepresidente")
	f.SetCellValue("Sheet1", "B3", "Partido político")
	f.SetCellValue("Sheet1", "C3", "Votos")
	f.SetCellValue("Sheet1", "D3", "%")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowCells := []string{"A3", "B3", "C3", "D3"}
	for _, cell := range thirdRowCells {
		if cell != "A3" {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowStyle); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SetCellStyle("Sheet1", cell, cell, thirdRowFirstColStyle); err != nil {
				fmt.Println(err)
			}
		}
	}

	for i, row := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Candidato)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.Partido)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), row.Votos)
		porcentaje := strings.Replace(row.Porcentaje, ",", ".", -1)
		porcentajeF, err := strconv.ParseFloat(porcentaje, 64)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+4), porcentajeF)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A16", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowSecondColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B16", fourthRowSecondColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D4", "D16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A13", "B13"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A14", "B14"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A15", "B15"); err != nil {
		fmt.Println(err)
	}

	// Merge cells
	if err := f.MergeCell("Sheet1", "A16", "B16"); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "D13", "D16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A17", GetTime())

	if err := f.MergeCell("Sheet1", "A17", "D17"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A17", "D17", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}
