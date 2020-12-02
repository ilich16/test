package report

import (
	"api/controller/functionality/report"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GeneratePresidentCountryReport(c *gin.Context) {

	log.Println("Hit: GeneratePresidentCountryReport")

	data := make([]report.Report, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportPresidentCountry(data)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GeneratePresidentCityReport(c *gin.Context) {

	log.Println("Hit: GeneratePresidentCityReport")

	city := getCityFormatted(c.Param("city"))

	data := make([]report.Report, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportPresidentCity(data, city)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GeneratePresidentDistrictReport(c *gin.Context) {

	log.Println("Hit: GeneratePresidentCityReport")

	district := c.Param("district")

	data := make([]report.Report, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportPresidentDistrict(data, district)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GeneratePresidentPrecinctReport(c *gin.Context) {

	log.Println("Hit: GeneratePresidentPrecinctReport")

	precinct := c.Param("precinct")

	log.Println(precinct)

	data := make([]report.Report, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportPresidentPrecinct(data, precinct)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GenerateMemberDistrictReport(c *gin.Context) {

	log.Println("Hit: GeneratePresidentCityReport")

	district := c.Param("district")

	data := make([]report.Report, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportMemberDistrict(data, district)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GenerateProceedingReportCity(c *gin.Context) {

	log.Println("Hit: GenerateProceedingReportCity")

	data := make([]report.ProceedingCity, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportProceedingCity(data)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func GenerateProceedingReport(c *gin.Context) {

	log.Println("Hit: GenerateProceedingReport")

	data := make([]report.Proceeding, 0)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	file := report.ReportProceeding(data)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.xlsx"`)
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}

func DownloadProceeding(c *gin.Context) {

	log.Println("Hit: DownloadProceeding")

	proceedingCode := c.Param("proceeding")

	file, err := ioutil.ReadFile(proceedingCode + ".jpg")
	if err != nil {
		log.Println(err)
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", `attachment; filename="file.jpg"`)
	c.Data(http.StatusOK, "application/octet-stream", file)

}
