package server

type WeatherRequest struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
