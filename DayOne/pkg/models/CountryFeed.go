package models

//CountryFeed hold all the seeds
type CountryFeed struct {
	Name      string   `json:"name"`
	Iso       string   `json:"iso"`
	Iso3      string   `json:"iso3"`
	Numcode   int      `json:"numcode"`
	Phonecode int      `json:"phonecode"`
	Sources   []Source `json:"sources"`
}
