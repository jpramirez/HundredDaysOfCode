package models

type Source struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Feedlink string `json:"feedlink"`
	Valid    string `json:"valid"`
}
