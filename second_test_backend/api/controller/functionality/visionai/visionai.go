package visionai

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	vision "cloud.google.com/go/vision/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

type Proceeding struct {
	ID               int    `json:"id"`
	Numero           int    `json:"numero"`
	Circunscripcion  int    `json:"circunscripcion"`
	Departamento     string `json:"departamento"`
	Provincia        string `json:"provincia"`
	Municipio        string `json:"municipio"`
	Localidad        string `json:"localidad"`
	Recinto          string `json:"recinto"`
	PresidenteCC     int    `json:"presidenteCC"`
	PresidenteFPV    int    `json:"presidenteFPV"`
	PresidenteMTS    int    `json:"presidenteMTS"`
	PresidenteUCS    int    `json:"presidenteUCS"`
	PresidenteMAS    int    `json:"presidenteMAS"`
	PresidenteVeF    int    `json:"presidenteVeF"`
	PresidentePDC    int    `json:"presidentePDC"`
	PresidenteMNR    int    `json:"presidenteMNR"`
	PresidentePAN    int    `json:"presidentePAN"`
	PresidenteValido int    `json:"presidenteValido"`
	PresidenteBlanco int    `json:"presidenteBlanco"`
	PresidenteNulo   int    `json:"presidenteNulo"`
	DiputadoCC       int    `json:"diputadoCC"`
	DiputadoFPV      int    `json:"diputadoFPV"`
	DiputadoMTS      int    `json:"diputadoMTS"`
	DiputadoUCS      int    `json:"diputadoUCS"`
	DiputadoMAS      int    `json:"diputadoMAS"`
	DiputadoVeF      int    `json:"diputadoVeF"`
	DiputadoPDC      int    `json:"diputadoPDC"`
	DiputadoMNR      int    `json:"diputadoMNR"`
	DiputadoPAN      int    `json:"diputadoPAN"`
	DiputadoValido   int    `json:"diputadoValido"`
	DiputadoBlanco   int    `json:"diputadoBlanco"`
	DiputadoNulo     int    `json:"diputadoNulo"`
	Dispositivo      string `json:"dispositivo"`
	Ubicacion        string `json:"ubicacion"`
	VotosTotal       int    `json:"votosTotal"`
}

func DetectText(file, dispositivo, ubicacion string) Proceeding {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		fmt.Println(err)
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		fmt.Println(err)
	}

	var (
		id              int
		numero          int
		circunscripcion int
		departamento    string
		provincia       string
		municipio       string
		localidad       string
		recinto         string
	)

	var ubicationLimitFirstX int32
	var ubicationLimitFirstY int32
	var ubicationLimitSecondX int32
	var ubicationLimitSecondY int32
	var votesLimitFirstX int32
	var votesLimitSecondX int32
	var memberLimitFirstX int32
	var memberLimitFirstY int32
	var memberLimitSecondX int32
	var ccLimitY int32
	var panLimitY int32
	var observacionLimitY int32
	//var votesLimitSecondY int32
	if len(annotations) == 0 {
		fmt.Println("No text found.")
	} else {
		// Encontrar el código de mesa
		for i, annotation := range annotations {
			if i != 0 {
				if isCodeOfTable(annotation.Description) {
					id, err = strconv.Atoi(annotation.Description)
					if err != nil {
						fmt.Println(err)
					}
					break
				}
			}
		}

		// Econtrar el número de mesa
		for i, annotation := range annotations {
			if i != 0 {
				if annotation.Description == "MESA:" && isNumber(annotations[i+1].Description) {
					numero, err = strconv.Atoi(annotations[i+1].Description)
					if err != nil {
						fmt.Println(err)
					}
					break
				}
			}
		}

		// Encontrar circunscripción uninominal
		for i, annotation := range annotations {
			if i != 0 {
				if annotation.Description == "UNINOMINAL" && isNumber(annotations[i+1].Description) {
					circunscripcion, err = strconv.Atoi(annotations[i+1].Description)
					if err != nil {
						fmt.Println(err)
					}
					break
				}
			}
		}

		// Encontrar votos en total
		/*var firstLimitX int32
		for i, annotation := range annotations {
			if i != 0 {
				if annotation.Description == "Departamento:" {
					firstLimitX = annotation.BoundingPoly.Vertices[0].X
				}
				if annotation.BoundingPoly.Vertices[2].X < firstLimitX {
					fmt.Println(annotation.Description)
				}
			}
		}*/

		// Encontrar ubicación de la mesa
		// Limites
		for i, annotation := range annotations {
			if i != 0 {
				if annotation.Description == "Departamento:" {
					ubicationLimitFirstX = annotation.BoundingPoly.Vertices[1].X
				}
				if annotation.Description == "PARA" && annotations[i+1].Description == "LLENAR" {
					ubicationLimitSecondX = annotation.BoundingPoly.Vertices[0].X
				}
				if annotation.Description == "UBICACIÓN" {
					ubicationLimitFirstY = annotation.BoundingPoly.Vertices[2].Y
				}
				if annotation.Description == "CÓMPUTO" {
					if annotations[i+1].Description == "DE" {
						if annotations[i+2].Description == "VOTOS" {
							ubicationLimitSecondY = annotation.BoundingPoly.Vertices[0].Y
							if annotations[i+3].Description == "OBTENIDOS" {
								votesLimitSecondX = annotations[i+3].BoundingPoly.Vertices[1].X
							}
						}
					}
				}
				if annotation.Description == "PRESIDENTE/A" {
					votesLimitFirstX = annotation.BoundingPoly.Vertices[1].X
					ccLimitY = annotation.BoundingPoly.Vertices[2].Y
					memberLimitFirstY = ccLimitY
				}
				if annotation.Description == "PAN-BOL" {
					panLimitY = annotation.BoundingPoly.Vertices[2].Y
				}
				if annotation.Description == "OBSERVACIONES" {
					observacionLimitY = annotation.BoundingPoly.Vertices[0].Y
				}
				if annotation.Description == "DIPUTADO/A" {
					memberLimitFirstX = annotation.BoundingPoly.Vertices[1].X
					memberLimitSecondX = memberLimitFirstX + ((annotation.BoundingPoly.Vertices[1].X - annotation.BoundingPoly.Vertices[0].X) * 3)
				}
			}
		}

		var ubicationSlice []*pb.EntityAnnotation
		ubicationIndex := make([]int, 0)
		for i, annotation := range annotations {
			if i != 0 {
				if annotation.BoundingPoly.Vertices[0].X > ubicationLimitFirstX {
					if annotation.BoundingPoly.Vertices[0].Y > ubicationLimitFirstY {
						if annotation.BoundingPoly.Vertices[2].Y < ubicationLimitSecondY {
							if annotation.BoundingPoly.Vertices[1].X < ubicationLimitSecondX {
								ubicationSlice = append(ubicationSlice, annotation)
								ubicationIndex = append(ubicationIndex, i)
							}
						}
					}
				}
			}
		}

		i := 0
		for {
			if ubicationIndex[i+1] != ubicationIndex[i]+1 {
				departamento += ubicationSlice[i].Description
				i++
				break
			} else {
				departamento += ubicationSlice[i].Description + " "
			}
			i++
		}
		for {
			if ubicationIndex[i+1] != ubicationIndex[i]+1 {
				provincia += ubicationSlice[i].Description
				i++
				break
			} else {
				provincia += ubicationSlice[i].Description + " "
			}
			i++
		}
		for {
			if ubicationIndex[i+1] != ubicationIndex[i]+1 {
				municipio += ubicationSlice[i].Description
				i++
				break
			} else {
				municipio += ubicationSlice[i].Description + " "
			}
			i++
		}
		for {
			if ubicationIndex[i+1] != ubicationIndex[i]+1 {
				localidad += ubicationSlice[i].Description
				i++
				break
			} else {
				localidad += ubicationSlice[i].Description + " "
			}
			i++
		}
		for {
			if i == len(ubicationIndex)-1 {
				recinto += ubicationSlice[i].Description
				break
			}
			recinto += ubicationSlice[i].Description + " "
			if ubicationIndex[i+1] != ubicationIndex[i]+1 {
				i++
				break
			}
			i++
		}

	}

	fmt.Println(id)
	fmt.Println(numero)
	fmt.Println(circunscripcion)
	fmt.Println(departamento)
	fmt.Println(provincia)
	fmt.Println(municipio)
	fmt.Println(localidad)
	fmt.Println(recinto)

	// Secooooooond partttttt
	annotation, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("CC limite")
	fmt.Println(ccLimitY)
	fmt.Println("Pan limite")
	fmt.Println(panLimitY)

	salto := ((panLimitY - ccLimitY) / 9) + 2
	limitForVotes := ccLimitY
	memberSalto := salto
	fmt.Println(salto)

	// Resultados para presidente

	var presidenteCC, presidenteFPV, presidenteMTS, presidenteUCS, presidenteMAS, presidenteVeF, presidentePDC, presidenteMNR, presidentePAN, presidenteValido, presidenteBlanco, presidenteNulo int

	if annotation == nil {
		fmt.Println("No text found.")
	} else {
		var result string
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteCC, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteFPV, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteMTS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteUCS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteMAS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteVeF, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidentePDC, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidenteMNR, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < panLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += salto
		presidentePAN, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		var aux int32
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[2].Y < limitForVotes+salto+salto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
									aux = word.BoundingBox.Vertices[2].Y
								}
							}
						}
					}
				}
			}
		}
		limitForVotes = aux
		newSaltoSecond := (observacionLimitY - limitForVotes) / 2
		newSaltoThird := (observacionLimitY - limitForVotes) / 3
		presidenteValido, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[0].Y < observacionLimitY && word.BoundingBox.Vertices[2].Y < limitForVotes+newSaltoThird+newSaltoThird {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		limitForVotes += newSaltoSecond
		presidenteBlanco, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > votesLimitFirstX {
							if word.BoundingBox.Vertices[0].X < votesLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > limitForVotes && word.BoundingBox.Vertices[2].Y < observacionLimitY {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		presidenteNulo, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
	}

	var diputadoCC, diputadoFPV, diputadoMTS, diputadoUCS, diputadoMAS, diputadoVeF, diputadoPDC, diputadoMNR, diputadoPAN, diputadoValido, diputadoBlanco, diputadoNulo int

	if annotation == nil {
		fmt.Println("No text found.")
	} else {
		var result string
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoCC, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoFPV, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoMTS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoUCS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoMAS, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoVeF, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoPDC, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoMNR, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += memberSalto
		diputadoPAN, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		var aux int32
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[1].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+memberSalto+memberSalto {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
									aux = word.BoundingBox.Vertices[2].Y
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY = aux
		newSaltoSecond := (observacionLimitY - memberLimitFirstY) / 2
		newSaltoThird := (observacionLimitY - memberLimitFirstY) / 3
		diputadoValido, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[0].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < memberLimitFirstY+newSaltoThird+newSaltoThird {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		memberLimitFirstY += newSaltoSecond
		diputadoBlanco, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
		result = ""
		for _, page := range annotation.Pages {
			for _, block := range page.Blocks {
				for _, paragraph := range block.Paragraphs {
					for _, word := range paragraph.Words {
						if word.BoundingBox.Vertices[0].X > memberLimitFirstX {
							if word.BoundingBox.Vertices[0].X < memberLimitSecondX {
								if word.BoundingBox.Vertices[0].Y > memberLimitFirstY && word.BoundingBox.Vertices[2].Y < observacionLimitY {
									symbols := make([]string, len(word.Symbols))
									for i, s := range word.Symbols {
										symbols[i] = s.Text
									}
									wordText := strings.Join(symbols, "")
									formatString := formatText2Number(wordText)
									fmt.Println(wordText)
									result += formatString
								}
							}
						}
					}
				}
			}
		}
		diputadoNulo, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(presidenteCC)
	fmt.Println(presidenteFPV)
	fmt.Println(presidenteMTS)
	fmt.Println(presidenteUCS)
	fmt.Println(presidenteMAS)
	fmt.Println(presidenteVeF)
	fmt.Println(presidentePDC)
	fmt.Println(presidenteMNR)
	fmt.Println(presidentePAN)
	fmt.Println(presidenteValido)
	fmt.Println(presidenteBlanco)
	fmt.Println(presidenteNulo)
	fmt.Println("Diputado")
	fmt.Println(diputadoCC)
	fmt.Println(diputadoFPV)
	fmt.Println(diputadoMTS)
	fmt.Println(diputadoUCS)
	fmt.Println(diputadoMAS)
	fmt.Println(diputadoVeF)
	fmt.Println(diputadoPDC)
	fmt.Println(diputadoMNR)
	fmt.Println(diputadoPAN)
	fmt.Println(diputadoValido)
	fmt.Println(diputadoBlanco)
	fmt.Println(diputadoNulo)
	votosTotal := presidenteValido + presidenteBlanco + presidenteNulo
	fmt.Println(votosTotal)

	result := Proceeding{id, numero, circunscripcion, departamento, provincia, municipio, localidad, recinto, presidenteCC, presidenteFPV, presidenteMTS, presidenteUCS, presidenteMAS, presidenteVeF, presidentePDC, presidenteMNR, presidentePAN, presidenteValido, presidenteBlanco, presidenteNulo, diputadoCC, diputadoFPV, diputadoMTS, diputadoUCS, diputadoMAS, diputadoVeF, diputadoPDC, diputadoMNR, diputadoPAN, diputadoValido, diputadoBlanco, diputadoNulo, dispositivo, ubicacion, votosTotal}

	return result
}

func isCodeOfTable(word string) bool {

	if len(word) != 5 {
		return false
	}

	for _, char := range word {
		if !unicode.IsNumber(char) {
			return false
		}
	}

	return true

}

func isNumber(word string) bool {

	for _, char := range word {
		if !unicode.IsNumber(char) {
			return false
		}
	}

	return true

}

func formatText2Number(word string) string {
	var result string

	for _, char := range word {
		if unicode.IsNumber(char) {
			result += string(char)
		}
	}

	return result

}
