package entity

const (
	ErrInvalidZipCodeMsg  = "invalid zipcode"
	ErrAddressNotFoundMsg = "can not find zipcode"
)

type Address struct {
	City  string
	State string
}

func NewAddress(city, state string) *Address {
	return &Address{
		City:  city,
		State: state,
	}
}
