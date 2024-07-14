package web

import (
	"encoding/json"
	"net/http"

	"github.com/Dionizio8/go-temppc/internal/entity"
	"github.com/Dionizio8/go-temppc/internal/usecase"
	"github.com/go-chi/chi"
)

type WebTemperatureHandler struct {
	AddressRepository     entity.AddressRepositoryInterface
	TemperatureRepository entity.TemperatureRepositoryInterface
}

func NewWebTemperatureHandler(addressRepository entity.AddressRepositoryInterface, temperatureRepository entity.TemperatureRepositoryInterface) *WebTemperatureHandler {
	return &WebTemperatureHandler{
		AddressRepository:     addressRepository,
		TemperatureRepository: temperatureRepository,
	}
}

func (t *WebTemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getTemperatureUserCase := usecase.NewGetTemperatureUseCase(t.AddressRepository, t.TemperatureRepository)
	temperature, err := getTemperatureUserCase.Execute(zipCode)
	if err != nil {
		msgErr := err.Error()
		if msgErr == entity.ErrAddressNotFoundMsg {
			w.WriteHeader(http.StatusNotFound)
		} else if msgErr == entity.ErrInvalidZipCodeMsg {
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(msgErr))
		return
	}

	err = json.NewEncoder(w).Encode(temperature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
