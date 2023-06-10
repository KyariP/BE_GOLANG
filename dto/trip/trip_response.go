package countrydto

import "dumbmerch/models"

type TripResponse struct {
	Id             int                    `json:"id"`
	Title          string                 `json:"title"`
	IdCountry      int                    `json:"id_country"`
	Country        models.CountryResponse `json:"country" `
	Accomodation   string                 `json:"accomodation"`
	Transportation string                 `json:"transportation"`
	Eat            string                 `json:"eat"`
	Day            int                    `json:"day"`
	Night          int                    `json:"night"`
	DateTrip       string                 `json:"date_trip"`
	Price          int                    `json:"price"`
	Quota          int                    `json:"quota"`
	Current_Quota  int                    `json:"current_quota"`
	Description    string                 `json:"description"`
	Image          string                 `json:"image"`
}
