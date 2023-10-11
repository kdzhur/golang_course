package data

type WeatherResponse struct {
	Location `json:"location"`
	Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdated string  `json:"last_updated"`
	TempC       float64 `json:"temp_c"`
	IsDay       int     `json:"is_day"`
	Condition   `json:"condition"`
	WindKph     float64 `json:"wind_kph"`
	WindDegree  int     `json:"wind_degree"`
	WindDir     string  `json:"wind_dir"`
	Humidity    int     `json:"humidity"`
	Cloud       int     `json:"cloud"`
	FeelslikeC  float64 `json:"feelslike_c"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
