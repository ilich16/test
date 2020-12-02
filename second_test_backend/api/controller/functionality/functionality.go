package functionality

import (
	"api/controller/functionality/visionai"
	"api/database/postgres/connection"
	"api/database/postgres/query/login"
	"api/database/postgres/query/precinct"
	"api/database/postgres/query/proceeding"
	"api/database/postgres/query/ubication"
	"api/database/postgres/query/votes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	log.Println("Hit: Login")

	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	if !login.VerifyUser(connection.DB, credentials.Username, credentials.Password) {
		c.String(http.StatusBadRequest, "bad credentials")
		return
	}

	c.String(http.StatusOK, "Ok!")

}

func SendImage(c *gin.Context) {

	log.Println("Hit: SendImage")

	// Extract file from the Form
	file, err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form:= %s", err.Error()))
		return
	}

	location := c.PostForm("location")

	device := c.PostForm("device")

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err:= %s", err.Error()))
		return
	}

	proceeding := visionai.DetectText("IMG_20201001_200016.jpg", device, location)

	os.Rename(filename, strconv.Itoa(proceeding.ID)+".jpg")

	c.JSON(http.StatusOK, proceeding)

}

func SendConfirmationData(c *gin.Context) {

	log.Println("Hit: SendConfirmationData")

	var proceedingO visionai.Proceeding

	if err := c.ShouldBindJSON(&proceedingO); err != nil {
		log.Println(err)
	}

	fmt.Println(proceedingO.ID)

	if proceeding.ExistsProceeding(connection.DB, proceedingO.ID) {
		c.JSON(http.StatusAccepted, proceedingO)
		return
	}

	if !ubication.ExistsUbication(connection.DB, proceedingO.Departamento, proceedingO.Provincia, proceedingO.Municipio, proceedingO.Localidad) {
		if err := ubication.InsertUbication(connection.DB, proceedingO.Departamento, proceedingO.Provincia, proceedingO.Municipio, proceedingO.Localidad); err != nil {
			fmt.Println(err)
		}
	}

	ubicationID := ubication.GetUbicationID(connection.DB, proceedingO.Departamento, proceedingO.Provincia, proceedingO.Municipio, proceedingO.Localidad)

	if !precinct.ExistsPrecinct(connection.DB, proceedingO.Recinto, ubicationID) {
		if err := precinct.InsertPrecinct(connection.DB, proceedingO.Recinto, ubicationID); err != nil {
			fmt.Println(err)
		}
	}

	precinctID := precinct.GetPrecinctID(connection.DB, proceedingO.Recinto, ubicationID)

	if err := proceeding.InsertProceeding(connection.DB, proceedingO.ID, proceedingO.Numero, proceedingO.Circunscripcion, proceedingO.Ubicacion, proceedingO.Dispositivo, precinctID); err != nil {
		fmt.Println(err)
	}

	if err := votes.InsertVotesPresident(connection.DB, proceedingO); err != nil {
		fmt.Println(err)
	}

	if err := votes.InsertVotesMember(connection.DB, proceedingO); err != nil {
		fmt.Println(err)
	}

}

func SendUpdateData(c *gin.Context) {

	log.Println("Hit: SendUpdateData")

	var proceedingO visionai.Proceeding

	if err := c.ShouldBindJSON(&proceedingO); err != nil {
		fmt.Println(err)
	}

	if err := votes.UpdateVotesPresident(connection.DB, proceedingO); err != nil {
		fmt.Println(err)
	}

	if err := votes.UpdateVotesMember(connection.DB, proceedingO); err != nil {
		fmt.Println(err)
	}

	log.Println(proceedingO)

}
