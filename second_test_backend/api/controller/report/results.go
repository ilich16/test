package report

import (
	"log"
	"net/http"
	"strconv"

	"api/database/postgres/connection"
	"api/database/postgres/query/president"
	"api/database/postgres/query/proceeding"

	"github.com/gin-gonic/gin"
)

func ResultsPresidentCountry(c *gin.Context) {

	log.Println("Hit: ResultsPresidentCountry")

	result := make([]president.ResPresidencial, 0)

	result, err := president.GetResults(connection.DB)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func ResultsPresidentCity(c *gin.Context) {

	log.Println("Hit: ResultsPresidentCity")

	result := make([]president.ResPresidencial, 0)

	city := getCityFormatted(c.Param("city"))

	result, err := president.GetResultsCity(connection.DB, city)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func GetAvailablePrecincts(c *gin.Context) {

	log.Println("Hit: GetAvailablePrecincts")

	result := make([]president.Precinct, 0)

	city := getCityFormatted(c.Param("city"))

	result, err := president.GetPrecincts(connection.DB, city)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

type PrecinctPost struct {
	Name string `json:"name"`
}

func GetResultsFromPrecinct(c *gin.Context) {

	log.Println("Hit: GetResultsFromPrecinct")

	result := make([]president.ReportResultsFromRecinct, 0)

	var precinct PrecinctPost

	if err := c.ShouldBindJSON(&precinct); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	result, err := president.GetResultsFromRecinct(connection.DB, precinct.Name)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

type DistrictPost struct {
	Circunscripcion string `json:"circunscripcion"`
}

func ResultsFromDistrict(c *gin.Context) {

	log.Println("Hit: ResultsFromDistrict")

	var district DistrictPost

	if err := c.ShouldBindJSON(&district); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	number, err := strconv.Atoi(district.Circunscripcion)
	if err != nil {
		log.Println(err)
	}

	result, err := president.GetResultsFromDistrict(connection.DB, number)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func ReportFromProceeding(c *gin.Context) {

	log.Println("Hit: ReportFromProceeding")

	result, err := proceeding.GetTotalProceedings(connection.DB)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func ProceedingUbication(c *gin.Context) {

	log.Println("Hit: ProceedingUbication")

	proceedingCode := 15348

	result, err := proceeding.GetUbicationProceeding(connection.DB, proceedingCode)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func AvailableProceedings(c *gin.Context) {

	log.Println("Hit: AvailableProceedings")

	result, err := proceeding.ProceedingValidate(connection.DB)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func getCityFormatted(city string) string {

	switch city {
	case "santa-cruz":
		return "Santa Cruz"
	case "la-paz":
		return "La Paz"
	case "cochabamba":
		return "Cochabamba"
	case "tarija":
		return "Tarija"
	case "pando":
		return "Pando"
	case "beni":
		return "Beni"
	case "oruro":
		return "Oruro"
	case "sucre":
		return "Sucre"
	default:
		return "Potosí"
	}

}
