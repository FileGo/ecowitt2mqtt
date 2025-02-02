package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertUnits(t *testing.T) {
	assert := assert.New(t)

	data := msg{
		TempInF:         15,
		TempOutF:        15,
		BaromRelIn:      29.92,
		BaromAbsIn:      29.92,
		WindSpdMph:      10,
		WindGustMph:     15,
		MaxDailyGustMph: 20,
	}

	data.convertUnits()

	assert.NotNil(data.TempInC)
	assert.NotNil(data.TempOutC)
	assert.NotNil(data.BaromRelHpa)
	assert.NotNil(data.BaromAbsHpa)
	assert.NotNil(data.WindSpdMps)
	assert.NotNil(data.WindGustMps)
	assert.NotNil(data.MaxDailyGustMps)
}
