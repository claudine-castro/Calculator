package main

import (
	"errors"
	"math"
	"strings"
)

func CalculateCircleArea(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, errors.New("radius value cannot be negative or zero")
	}
	return math.Round((math.Pi*radius*radius)*100) / 100, nil
}

func CalculateCylinderVolume(radius, height float64) (float64, error) {
	if radius <= 0 || height <= 0 {
		return 0, errors.New("radius or height value cannot be negative or zero")
	}
	return math.Round((math.Pi*radius*radius*height)*100) / 100, nil
}

//

type Person struct {
	name string
	city string
}

func GetNameThatStartsWithA(person Person) bool {
	return person.name != "" && strings.ToUpper(person.name)[0] == 'A'
}

func GetNamesByCity(persons []Person, city string) ([]string, error) {
	var filteredNamesByCity []string

	for _, person := range persons {
		if strings.ToUpper(person.city) == strings.ToUpper(city) {
			filteredNamesByCity = append(filteredNamesByCity, person.name)
		}
	}

	if len(filteredNamesByCity) == 0 {
		return nil, errors.New("no matching city found")
	}

	return filteredNamesByCity, nil
}
