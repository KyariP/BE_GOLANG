package countrydto

type CreateTrip struct {
	Title          string `json:"title" form:"title" validate:"required"`
	IdCountry      int    `json:"id_country" form:"id_country" validate:"required"`
	Accomodation   string `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string `json:"transportation" form:"transportation" validate:"required"`
	Eat            string `json:"eat" form:"eat" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	Current_Quota  int    `json:"current_quota" form:"current_quota"`
	Description    string `json:"description" form:"description" validate:"required"`
	Image          string `json:"image" form:"image"`
}

type UpdateTrip struct {
	Title          string `json:"title" form:"title"`
	IdCountry      int    `json:"id_country" form:"id_country"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	DateTrip       string `json:"date_trip" form:"date_trip"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Current_Quota  int    `json:"current_quota" form:"current_quota"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}
