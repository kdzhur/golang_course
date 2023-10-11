package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rest/cmd/task1/internal/data"
)

func GetWeather(city string) (*data.WeatherResponse, error) {

	apiKey, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		return nil, fmt.Errorf("no API key provided. WEATHER_API_KEY env variable should be set")
	}

	apiUrl := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		weather := new(data.WeatherResponse)
		err := json.Unmarshal(body, weather)
		if err != nil {
			return nil, err
		}
		return weather, nil
	}

	return nil, fmt.Errorf("error making request to weather API: %s", resp.Status)
}
