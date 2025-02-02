package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFtoC(t *testing.T) {
	assert := assert.New(t)
	testDelta := 0.05

	testdata := []struct {
		Input  float64
		Expect float64
	}{
		{-50.0, -45.56},
		{-40.0, -40.00},
		{0.0, -17.78},
		{32.0, 0.0},
		{98.6, 37.0},
		{212.0, 100.0},
	}

	for _, val := range testdata {
		assert.InDelta(val.Expect, FtoC(val.Input), testDelta)
	}
}

func TestInToHpa(t *testing.T) {
	assert := assert.New(t)
	testDelta := 0.2

	testdata := []struct {
		Input  float64
		Expect float64
	}{
		{0, 0},
		{29.92, 1013.25},
		{29.38, 995},
		{29.56, 1001},
		{30.00, 1016},
		{30.70, 1039.7},
		{28.40, 961.8},
	}

	for _, val := range testdata {
		assert.InDelta(val.Expect, InToHpa(val.Input), testDelta)
	}
}

func TestMphToMps(t *testing.T) {
	assert := assert.New(t)
	testDelta := 0.1

	testdata := []struct {
		Input  float64
		Expect float64
	}{
		{0, 0},
		{1, 0.4470},
		{5, 2.2352},
		{9, 4.0234},
		{10, 4.4704},
		{50, 22.352},
	}

	for _, val := range testdata {
		assert.InDelta(val.Expect, MphToMps(val.Input), testDelta)
	}
}
