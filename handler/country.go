package handler

import (
	countrydto "dumbmerch/dto/country"
	resultdto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerCountries struct {
	CountryRepository repositories.CountryRepository
}

func HandlerCountry(CountryRepository repositories.CountryRepository) *HandlerCountries {
	return &HandlerCountries{CountryRepository}
}

func (h *HandlerCountries) FindCountry(c echo.Context) error {
	country, err := h.CountryRepository.FindCountry()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: country})

}

func (h *HandlerCountries) FindCountryId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	country, _ := h.CountryRepository.FindCountryId(id)

	if country.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: country})
}

func (h *HandlerCountries) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	country, _ := h.CountryRepository.FindCountryId(id)

	if country.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.CountryRepository.DeleteCountry(id, country)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func (h *HandlerCountries) CreateCountry(c echo.Context) error {
	request := new(countrydto.CreateCountry)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// country, _ := h.CountryRepository.FindCountryId(Id)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	country := models.Country{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := h.CountryRepository.CreateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func (h *HandlerCountries) UpdateCountry(c echo.Context) error {
	request := new(countrydto.UpdateCountry)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.FindCountryId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Name != "" {
		country.Name = request.Name
	}

	country.UpdatedAt = time.Now()

	data, err := h.CountryRepository.UpdateCountry(id, country)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func convertResponseCountry(Country models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		Id:   Country.Id,
		Name: Country.Name,
	}
}
