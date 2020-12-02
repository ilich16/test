package main

import (
	"api/controller"
	"api/controller/functionality"
	"api/controller/proceeding"
	"api/controller/report"
	"api/controller/result/member"
	"api/database/postgres/connection"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {

	if err := connection.OpenConnection(); err != nil {
		log.Println(err)
	}

}

func main() {

	r := gin.Default()
	r.Use(controller.CORSMiddleware())

	v1 := r.Group("/api/v1/sw1")
	{
		// Mobile
		v1.POST("/login", functionality.Login)
		v1.POST("/send-image", functionality.SendImage)
		v1.POST("/send-confirmation-data", functionality.SendConfirmationData)
		v1.POST("/send-update-data", functionality.SendUpdateData)
		// Vue
		v1.GET("/results-president-country", report.ResultsPresidentCountry)
		v1.GET("/results-president-city/:city", report.ResultsPresidentCity)
		v1.GET("/recincts/:city", report.GetAvailablePrecincts)
		v1.POST("/result-from-precinct", report.GetResultsFromPrecinct)
		v1.POST("/result-from-district", report.ResultsFromDistrict)
		// Proceedings
		v1.GET("/total-proceedings", proceeding.TotalProceedings)
		v1.GET("/proceeding-votes/:proceedingCode", proceeding.ProceedingVotes)
		v1.GET("/proceeding-ubication", report.ProceedingUbication)
		v1.GET("/available-proceedings", report.AvailableProceedings)
		// Members
		v1.POST("/member-result-from-district", member.ResultsFromDistrict)
		v1.POST("/member-result-from-precinct", member.ResultsFromDistrict)
		// Reports
		v1.POST("/generate-report-president-country", report.GeneratePresidentCountryReport)
		v1.POST("/generate-report-president-city/:city", report.GeneratePresidentCityReport)
		v1.POST("/generate-report-president-district/:district", report.GeneratePresidentDistrictReport)
		v1.POST("/generate-report-president-precinct/:precinct", report.GeneratePresidentPrecinctReport)
		v1.POST("/generate-report-member-district/:district", report.GenerateMemberDistrictReport)
		v1.POST("/generate-report-proceeding-city", report.GenerateProceedingReportCity)
		v1.POST("/generate-report-proceeding", report.GenerateProceedingReport)
		v1.POST("/download-proceeding/:proceeding", report.DownloadProceeding)
	}

	defer connection.CloseConnection()

	r.Run(":8000")
}
