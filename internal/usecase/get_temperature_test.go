package usecase

import (
	"errors"
	"testing"

	"github.com/Dionizio8/go-temppc/internal/entity"
	"github.com/Dionizio8/go-temppc/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureUseCase_Ok(t *testing.T) {
	addressRepository := mocks.NewMockAddressRepository(t)
	temperatureRepository := mocks.NewMockTemperatureRepository(t)
	usecase := NewGetTemperatureUseCase(addressRepository, temperatureRepository)

	addressRepository.On("GetAddress", "11111111").Return(entity.Address{City: "Limeira", State: "SP"}, nil).Once()
	temperatureRepository.On("GetTemperature", "Limeira").Return(entity.Temperature{Celsius: 20, Fahrenheit: 68, Kelvin: 300}, nil).Once()

	temp, err := usecase.Execute("11111111")

	assert.Nil(t, err)
	assert.Equal(t, 20.0, temp.TempC)
	assert.Equal(t, 68.0, temp.TempF)
	assert.Equal(t, 300.0, temp.TempK)
}

func TestGetTemperatureUseCase_InvalidZipCode(t *testing.T) {
	usecase := NewGetTemperatureUseCase(nil, nil)

	_, err := usecase.Execute("11111111ABC")

	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrInvalidZipCodeMsg, err.Error())
}

func TestGetTemperatureUseCase_AddressNotFound(t *testing.T) {
	addressRepository := mocks.NewMockAddressRepository(t)
	usecase := NewGetTemperatureUseCase(addressRepository, nil)

	addressRepository.On("GetAddress", "11111111").Return(entity.Address{}, errors.New(entity.ErrAddressNotFoundMsg)).Once()

	_, err := usecase.Execute("11111111")

	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrAddressNotFoundMsg, err.Error())
}

func TestGetTemperatureUseCase_TemperatureNotFound(t *testing.T) {
	addressRepository := mocks.NewMockAddressRepository(t)
	temperatureRepository := mocks.NewMockTemperatureRepository(t)
	usecase := NewGetTemperatureUseCase(addressRepository, temperatureRepository)

	addressRepository.On("GetAddress", "11111111").Return(entity.Address{City: "Limeira", State: "SP"}, nil).Once()
	temperatureRepository.On("GetTemperature", "Limeira").Return(entity.Temperature{}, errors.New("error")).Once()

	_, err := usecase.Execute("11111111")

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
}
