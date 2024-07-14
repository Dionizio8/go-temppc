package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Dionizio8/go-temppc/internal/entity"
)

type AddressViaCepDTO struct {
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type AddressRepository struct {
	ViaCEPClientURL string
}

func NewAddressRepository(viaCEPClientURL string) *AddressRepository {
	return &AddressRepository{
		ViaCEPClientURL: viaCEPClientURL,
	}
}

func (r *AddressRepository) GetAddress(zipCode string) (entity.Address, error) {
	url := fmt.Sprintf("%s/ws/%s/json", r.ViaCEPClientURL, zipCode)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return entity.Address{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.Address{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return entity.Address{}, errors.New(entity.ErrAddressNotFoundMsg)
		}
		return entity.Address{}, err
	}

	var addressViaCep AddressViaCepDTO
	err = json.NewDecoder(resp.Body).Decode(&addressViaCep)
	if err != nil {
		return entity.Address{}, err
	}

	return entity.Address{
		City:  addressViaCep.Localidade,
		State: addressViaCep.Uf,
	}, nil
}
