package models

type Movies struct {
	Title    string  `json:"title"`
	Runtime  int     `json:"runtime"`
	Rate     float32 `json:"rate"`
	Date     string  `json:"date"`
	Overview string  `json:"overview"`
}
