package proceeding

import (
	"log"
	"strconv"

	"api/database/postgres/connection"
	"api/database/postgres/query/proceeding"

	"github.com/gin-gonic/gin"
)

func TotalProceedings(c *gin.Context) {

	log.Println("Hit: ReportFromProceeding")

	result, err := proceeding.GetTotalProceedings(connection.DB)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}

func ProceedingVotes(c *gin.Context) {

	log.Println("Hit AvailableProceedings")

	proceedingCode, err := strconv.Atoi(c.Param("proceedingCode"))
	if err != nil {
		log.Println(err)
	}

	result, err := proceeding.ProceedingVotes(connection.DB, proceedingCode)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)

}
