package entity

type AddressRepositoryInterface interface {
	GetAddress(zipCode string) (Address, error)
}

type TemperatureRepositoryInterface interface {
	GetTemperature(city string) (Temperature, error)
}
