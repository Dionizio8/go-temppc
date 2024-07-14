package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dionizio8/go-temppc/internal/entity"
)

type TemperatureWeatherDTO struct {
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

type TemperatureRepository struct {
	WeatherAPIClientURL    string
	WeatherAPIClientAPIKey string
}

func NewTemperatureRepository(weatherAPIClientURL string, weatherAPIClientAPIKey string) *TemperatureRepository {
	return &TemperatureRepository{
		WeatherAPIClientURL:    weatherAPIClientURL,
		WeatherAPIClientAPIKey: weatherAPIClientAPIKey,
	}
}

func (r *TemperatureRepository) GetTemperature(city string) (entity.Temperature, error) {
	url := fmt.Sprintf("%s/v1/current.json?q=%s&lang=pt&key=%s", r.WeatherAPIClientURL, city, r.WeatherAPIClientAPIKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return entity.Temperature{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.Temperature{}, err
	}
	defer resp.Body.Close()

	var temperature TemperatureWeatherDTO
	err = json.NewDecoder(resp.Body).Decode(&temperature)
	if err != nil {
		return entity.Temperature{}, err
	}

	return *entity.NewTemperature(temperature.Current.TempC, temperature.Current.TempF), nil
}
