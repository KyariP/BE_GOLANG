package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handler.HandlerCountry(CountryRepository)
	e.GET("/country", middleware.Auth(h.FindCountry))
	e.GET("/country/:id", middleware.Auth(h.FindCountryId))
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
	e.POST("/country", middleware.Auth(h.CreateCountry))
	e.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))
}
