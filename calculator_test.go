package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCircleArea(t *testing.T) {
	err := errors.New("radius value cannot be negative or zero")

	tests := []struct {
		name      string
		input     float64
		expected  float64
		areaError error
	}{
		{"Valid value", 3.0, 28.27, nil},
		{"With decimal value", 2.5, 19.63, nil},
		{"Negative radius", -3.0, 0, err},
		{"Zero radius value", 0, 0, err},
	}

	assert := assert.New(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateCircleArea(tc.input)
			assert.Equal(tc.expected, got)
			assert.Equal(tc.areaError, err)
		})
	}
}

func TestCalculateCylinderVolume(t *testing.T) {
	err := errors.New("radius or height value cannot be negative or zero")

	tests := []struct {
		name        string
		radius      float64
		height      float64
		expected    float64
		volumeError error
	}{
		{"Valid value", 3.0, 5.0, 141.37, nil},
		{"With decimal value", 3.5, 5.5, 211.66, nil},
		{"Negative radius value", -3.0, 5.0, 0, err},
		{"Zero radius value", 0, 5.0, 0, err},
		{"Negative height value", 3.0, -5.0, 0, err},
		{"Zero height value", 3.0, 0, 0, err},
		{"Both negative values", -3.0, -5.0, 0, err},
		{"Both zero values", 0, 0, 0, err},
	}

	assert := assert.New(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateCylinderVolume(tc.radius, tc.height)
			assert.Equal(tc.expected, got)
			assert.Equal(tc.volumeError, err)
		})
	}
}

func TestGetNameThatStartsWithA(t *testing.T) {

	tests := []struct {
		name     string
		person   Person
		expected bool
	}{
		{"Starts with uppercase A",
			Person{name: "Angel", city: "Pasig"},
			true},
		{"Starts with lowercase A",
			Person{name: "anna", city: "Cavite"},
			true},
		{"Doesn't start with A",
			Person{name: "Baby", city: "Taguig"},
			false},
	}

	assert := assert.New(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetNameThatStartsWithA(tc.person)
			assert.Equal(tc.expected, got)
		})
	}
}

func TestGetNamesByCity(t *testing.T) {
	err := errors.New("no matching city found")

	tests := []struct {
		name      string
		persons   []Person
		city      string
		expected  []string
		cityError error
	}{
		{"Has returned names",
			[]Person{{name: "Anna", city: "Taguig"}, {name: "Elsa", city: "Taguig"}, {name: "Kristoff", city: "Pasig"}},
			"Taguig",
			[]string{"Anna", "Elsa"},
			nil},
		{"Has returned names - discarding letter case",
			[]Person{{name: "Anna", city: "TAGUIG"}, {name: "Elsa", city: "Taguig"}, {name: "Kristoff", city: "Pasig"}},
			"taguig",
			[]string{"Anna", "Elsa"},
			nil},
		{"No returned name",
			[]Person{{name: "Anna", city: "Taguig"}, {name: "Elsa", city: "Taguig"}, {name: "Kristoff", city: "Pasig"}},
			"Pateros",
			nil,
			err},
	}

	assert := assert.New(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetNamesByCity(tc.persons, tc.city)
			assert.Equal(tc.expected, got)
			assert.Equal(tc.cityError, err)
		})
	}
}
