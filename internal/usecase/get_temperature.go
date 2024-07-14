package usecase

import (
	"errors"
	"regexp"

	"github.com/Dionizio8/go-temppc/internal/entity"
)

type TemperatureOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type GetTemperatureUseCase struct {
	AddressRepository     entity.AddressRepositoryInterface
	TemperatureRepository entity.TemperatureRepositoryInterface
}

func NewGetTemperatureUseCase(addressRepository entity.AddressRepositoryInterface, temperatureRepository entity.TemperatureRepositoryInterface) *GetTemperatureUseCase {
	return &GetTemperatureUseCase{
		AddressRepository:     addressRepository,
		TemperatureRepository: temperatureRepository,
	}
}

func (uc *GetTemperatureUseCase) Execute(zipCode string) (TemperatureOutputDTO, error) {
	regex := regexp.MustCompile(`^[0-9]{8}$`)
	if !regex.MatchString(zipCode) {
		return TemperatureOutputDTO{}, errors.New(entity.ErrInvalidZipCodeMsg)
	}
	address, err := uc.AddressRepository.GetAddress(zipCode)
	if err != nil {
		return TemperatureOutputDTO{}, err
	}

	if address.City == "" {
		return TemperatureOutputDTO{}, errors.New(entity.ErrAddressNotFoundMsg)
	}

	temperature, err := uc.TemperatureRepository.GetTemperature(address.City)
	if err != nil {
		return TemperatureOutputDTO{}, err
	}

	return TemperatureOutputDTO{
		TempC: temperature.Celsius,
		TempF: temperature.Fahrenheit,
		TempK: temperature.Kelvin,
	}, nil
}
