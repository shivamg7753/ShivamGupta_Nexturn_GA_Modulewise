package service

import (
	"errors"
	"climate-analysis/model"
	// "fmt"
)

type ClimateService struct {
	Data []model.ClimateData
}

func NewClimateService() *ClimateService {
	return &ClimateService{
		Data: []model.ClimateData{
			{"New York", 16.0, 1200.5},
			{"London", 12.3, 800.7},
			{"Tokyo", 18.5, 1500.3},
			{"Sydney", 20.1, 900.4},
			{"Mumbai", 27.0, 2000.1},
		},
	}
}

// Find the city with the highest temperature
func (cs *ClimateService) HighestTemperature() model.ClimateData {
	highest := cs.Data[0]
	for _, city := range cs.Data {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// Find the city with the lowest temperature
func (cs *ClimateService) LowestTemperature() model.ClimateData {
	lowest := cs.Data[0]
	for _, city := range cs.Data {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// Calculate the average rainfall
func (cs *ClimateService) AverageRainfall() float64 {
	totalRainfall := 0.0
	for _, city := range cs.Data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cs.Data))
}

// Filter cities by rainfall threshold
func (cs *ClimateService) FilterByRainfall(threshold float64) []model.ClimateData {
	filtered := []model.ClimateData{}
	for _, city := range cs.Data {
		if city.Rainfall > threshold {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

// Search for a city by name
func (cs *ClimateService) SearchByCity(name string) (model.ClimateData, error) {
	for _, city := range cs.Data {
		if city.City == name {
			return city, nil
		}
	}
	return model.ClimateData{}, errors.New("city not found")
}
