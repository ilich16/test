package report

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Proceeding struct {
	Codigo          int    `json:"codigo"`
	Recinto         string `json:"recinto"`
	Mesa            int    `json:"mesa"`
	Circunscripcion int    `json:"circunscripcion"`
	Departamento    string `json:"departamento"`
	Provincia       string `json:"provincia"`
	Municipio       string `json:"municipio"`
	Localidad       string `json:"localidad"`
	Partido         string `json:"partido"`
	VotosPresidente int    `json:"votosPresidente"`
	VotosDiputado   int    `json:"votosDiputado"`
	Ubicacion       string `json:"ubicacion"`
	Dispositivo     string `json:"dispositivo"`
}

func ReportProceeding(data []Proceeding) bytes.Buffer {

	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "REGISTRO DE VOTOS")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "C1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "C1"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetRowHeight("Sheet1", 1, 31); err != nil {
		fmt.Println(err)
	}

	if err := f.SetRowHeight("Sheet1", 2, 60); err != nil {
		fmt.Println(err)
	}

	if err := f.SetColWidth("Sheet1", "A", "A", 35); err != nil {
		fmt.Println(err)
	}

	if err := f.SetColWidth("Sheet1", "B", "B", 15); err != nil {
		fmt.Println(err)
	}

	if err := f.SetColWidth("Sheet1", "C", "C", 35); err != nil {
		fmt.Println(err)
	}

	firstData := data[0]

	if err := f.SetCellRichText("Sheet1", "A2", []excelize.RichTextRun{
		{
			Text: "Código de acta: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: strconv.Itoa(firstData.Codigo),
		},
		{
			Text: "\r\nMesa: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: strconv.Itoa(firstData.Mesa),
		},
		{
			Text: "\r\nCir. Uninominal: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: strconv.Itoa(firstData.Circunscripcion),
		},
		{
			Text: "\r\nRecinto: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: firstData.Recinto,
		},
	}); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellRichText("Sheet1", "C2", []excelize.RichTextRun{
		{
			Text: "Departamento: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: firstData.Departamento,
		},
		{
			Text: "\r\nProvincia: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: firstData.Provincia,
		},
		{
			Text: "\r\nMunicipio: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: firstData.Municipio,
		},
		{
			Text: "\r\nLocalidad: ",
			Font: &excelize.Font{
				Bold: true,
			},
		},
		{
			Text: firstData.Localidad,
		},
	}); err != nil {
		fmt.Println(err)
	}

	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "A2", style); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C2", "C2", style); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "B2"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A3", "Partido político")
	f.SetCellValue("Sheet1", "B3", "Presidente/a")
	f.SetCellValue("Sheet1", "C3", "Diputado/a Cir. Uninominal")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A3", "A3", thirdRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B3", "C3", thirdRowStyle); err != nil {
		fmt.Println(err)
	}

	for i, row := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Partido)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.VotosPresidente)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), row.VotosDiputado)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A16", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C16", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B13", "B16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C16", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	link := "https://www.google.com/maps/search/?api=1&query=" + firstData.Ubicacion

	f.SetCellValue("Sheet1", "A17", "Obtener ubicación de donde fue registrada esta acta")

	if err := f.SetCellHyperLink("Sheet1", "A17", link, "External"); err != nil {
		fmt.Println(err)
	}

	linkStyle, err := f.NewStyle(`{"font":{"color":"#1265BE","underline":"single"}, "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A17", "A17", linkStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A17", "C17"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A18", "Esta acta fue registrada desde el siguiente dispositivo: "+firstData.Dispositivo)

	deviceStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A18", "A18", deviceStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A18", "C18"); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A19", GetTime())

	if err := f.MergeCell("Sheet1", "A19", "C19"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A19", "C19", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}

type ProceedingCity struct {
	Ciudad     string `json:"ciudad"`
	Actas      int    `json:"actas"`
	Porcentaje string `json:"porcentaje"`
}

func ReportProceedingCity(data []ProceedingCity) bytes.Buffer {
	f := excelize.NewFile()

	// Set value of a cell
	f.SetCellValue("Sheet1", "A1", "ELECCIONES GENERALES 2019")
	// Create style for title
	titleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":24}}`)
	if err != nil {
		fmt.Println(err)
	}
	// Apply style for title
	if err := f.SetCellStyle("Sheet1", "A1", "C1", titleStyle); err != nil {
		fmt.Println(err)
	}
	// Merge cells
	if err := f.MergeCell("Sheet1", "A1", "C1"); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A2", "ACTAS COMPUTADAS EN TODO EL PAÍS")

	subtitleStyle, err := f.NewStyle(`{"alignment": {"horizontal": "center"}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A2", "C2", subtitleStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.MergeCell("Sheet1", "A2", "C2"); err != nil {
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

	f.SetCellValue("Sheet1", "A3", "Departamento")
	f.SetCellValue("Sheet1", "B3", "Cantidad de actas")
	f.SetCellValue("Sheet1", "C3", "%")

	thirdRowFirstColStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": { "bold":true }, "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	thirdRowCells := []string{"A3", "B3", "C3"}
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
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), row.Ciudad)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+4), row.Actas)
		porcentaje := strings.Replace(row.Porcentaje, ",", ".", -1)
		porcentajeF, err := strconv.ParseFloat(porcentaje, 64)
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+4), porcentajeF)
	}

	fourthRowFirstColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A4", "A14", fourthRowFirstColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowSecondColStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "right"}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B4", "B14", fourthRowSecondColStyle); err != nil {
		fmt.Println(err)
	}

	fourthRowLastColsStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C4", "C14", fourthRowLastColsStyle); err != nil {
		fmt.Println(err)
	}

	votosTextStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "font": {"bold":true, "size":11}}`)
	if err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A13", "A14", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "C13", "C14", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "B13", "B14", votosTextStyle); err != nil {
		fmt.Println(err)
	}

	// For date
	dateStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#D9D9D9"],"pattern":1}, "border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}], "alignment": {"horizontal": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A15", GetTime())

	if err := f.MergeCell("Sheet1", "A15", "C15"); err != nil {
		fmt.Println(err)
	}

	if err := f.SetCellStyle("Sheet1", "A15", "C15", dateStyle); err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	if err := f.Write(&result); err != nil {
		fmt.Println(err)
	}

	return result

}

func GetTime() string {
	date := time.Now()
	year, month, day := date.Date()
	hour, minute, second := date.Clock()

	result := strconv.Itoa(day) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)

	return result
}
