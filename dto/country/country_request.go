package countrydto

type CreateCountry struct {
	Id   int    `json:"id_country"`
	Name string `json:"name_country"`
}

type UpdateCountry struct {
	Id   int    `json:"id_country"`
	Name string `json:"name_country"`
}
