package main

import (
	"net/http"

	"github.com/Dionizio8/go-temppc/configs"
	"github.com/Dionizio8/go-temppc/internal/infra/client"
	"github.com/Dionizio8/go-temppc/internal/infra/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	addressRepository := client.NewAddressRepository(cfg.ViaCEPClientURL)
	temperatureRepository := client.NewTemperatureRepository(cfg.WeatherAPIClientURL, cfg.WeatherAPIClientAPIKey)
	temperatureHandler := web.NewWebTemperatureHandler(addressRepository, temperatureRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/temperature", func(r chi.Router) {
		r.Get("/{zipCode}", temperatureHandler.GetTemperature)
	})

	http.ListenAndServe(cfg.WebServerPort, r)
}
