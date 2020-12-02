package member

import (
	"log"
	"net/http"
	"strconv"

	"api/database/postgres/connection"
	"api/database/postgres/query/member"

	"github.com/gin-gonic/gin"
)

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

	result, err := member.GetResultsFromDistrict(connection.DB, number)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}
