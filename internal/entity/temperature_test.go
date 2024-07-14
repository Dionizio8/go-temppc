package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTemperature_Ok(t *testing.T) {
	temp := NewTemperature(10, 15)
	assert.Equal(t, 10.0, temp.Celsius)
	assert.Equal(t, 15.0, temp.Fahrenheit)
	assert.Equal(t, 283.0, temp.Kelvin)
}