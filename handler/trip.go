package handler

import (
	resultdto "dumbmerch/dto/result"
	tripdto "dumbmerch/dto/trip"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

// var path_file = os.Getenv("PATH_FILE")

type HandlerTrips struct {
	TripRepository repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *HandlerTrips {
	return &HandlerTrips{TripRepository}
}

func (h *HandlerTrips) FindTrip(c echo.Context) error {
	trip, err := h.TripRepository.FindTrip()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	for i, p := range trip {
		trip[i].Image = path_file + p.Image
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: trip})

}

func (h *HandlerTrips) FindTripId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, _ := h.TripRepository.FindTripId(id)

	if trip.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	trip.Image = path_file + trip.Image
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: trip})

}

func (h *HandlerTrips) DeleteTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, _ := h.TripRepository.FindTripId(id)

	if trip.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TripRepository.DeleteTrip(id, trip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *HandlerTrips) CreateTrip(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)
	request := new(tripdto.CreateTrip)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// trip, _ := h.TripRepository.FindTripId(Id)

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

	countries, err := h.TripRepository.GetCountryId(request.IdCountry)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	trip := models.Trip{
		Title:          request.Title,
		IdCountry:      request.IdCountry,
		Country:        countries,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Current_Quota:  request.Current_Quota,
		Description:    request.Description,
		Image:          dataFile,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *HandlerTrips) UpdateTrip(c echo.Context) error {
	request := new(tripdto.UpdateTrip)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetUpdateId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	// countries, err := h.TripRepository.GetCountryId(request.IdCountry)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error()})
	// }

	if request.Title != "" {
		trip.Title = request.Title
	}

	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}

	// trip.Country = countries

	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}
	if request.Eat != "" {
		trip.Eat = request.Eat
	}
	if request.Day != 0 {
		trip.Day = request.Day
	}
	if request.Night != 0 {
		trip.Night = request.Night
	}
	if request.DateTrip != "" {
		trip.DateTrip = request.DateTrip
	}
	if request.Price != 0 {
		trip.Price = request.Price
	}
	if request.Quota != 0 {
		trip.Quota = request.Quota
	}
	if request.Current_Quota != 0 {
		trip.Current_Quota = request.Current_Quota
	}
	if request.Description != "" {
		trip.Description = request.Description
	}
	if request.Image != "" {
		trip.Image = request.Image
	}

	trip.UpdatedAt = time.Now()

	data, err := h.TripRepository.UpdateTrip(id, trip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func convertResponseTrip(Trip models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		Id:             Trip.Id,
		Title:          Trip.Title,
		IdCountry:      Trip.IdCountry,
		Country:        Trip.Country,
		Accomodation:   Trip.Accomodation,
		Transportation: Trip.Transportation,
		Eat:            Trip.Eat,
		Day:            Trip.Day,
		Night:          Trip.Night,
		DateTrip:       Trip.DateTrip,
		Price:          Trip.Price,
		Quota:          Trip.Quota,
		Current_Quota:  Trip.Current_Quota,
		Description:    Trip.Description,
		Image:          Trip.Image,
	}
}
