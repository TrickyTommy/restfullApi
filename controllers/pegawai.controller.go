package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restFullApi/models"
	"strconv"
)

func FetchAllPegawai(c echo.Context)error  {

	result, err := models.FetchAllPegawai()
	if err != nil {

		return  c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)

}
//ke 2 pegawai

func StorePegawai(c echo.Context)error  {
	//tampung parameter yang dikirim dari external
	nama:= c.FormValue("nama")
	alamat:= c.FormValue("alamat")
	telepon:= c.FormValue("telepon")

	//menampung apa yang dikembalikan oleh models
	result, err := models.StorePegawai(nama,alamat,telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
//ke 2 update pegawai

func UpdatePegawai(c echo.Context)error  {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")
	
	//konfersi string to int
	conv_id,err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegawai(conv_id, nama,alamat,telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
	
}
// ke 2

func Deletepegawai(c echo.Context)error  {
	id:= c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}

	return c.JSON(http.StatusOK, result)

}