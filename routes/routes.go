package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restFullApi/controllers"
)

func Init() *echo.Echo  {
	e := echo.New()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, this is echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)
	//pegawai ke 3
	e.POST("/pegawai",controllers.StorePegawai )
	//ke update 3
	e.PUT("/pegawai",controllers.UpdatePegawai)
	//ke 3 delete
	e. DELETE("/pegawai", controllers.Deletepegawai)
	return e


}
