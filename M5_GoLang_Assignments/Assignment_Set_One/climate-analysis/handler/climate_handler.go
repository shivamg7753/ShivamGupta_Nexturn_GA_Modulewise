package handler

import (
	"fmt"
	"climate-analysis/service"
)

func StartClimateAnalysis() {
	climateService := service.NewClimateService()

	for {
		fmt.Println("\nClimate Data Analysis")
		fmt.Println("1. Find City with Highest Temperature")
		fmt.Println("2. Find City with Lowest Temperature")
		fmt.Println("3. Calculate Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall Threshold")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			highest := climateService.HighestTemperature()
			fmt.Printf("City with Highest Temperature: %s (%.2f°C)\n", highest.City, highest.Temperature)
		case 2:
			lowest := climateService.LowestTemperature()
			fmt.Printf("City with Lowest Temperature: %s (%.2f°C)\n", lowest.City, lowest.Temperature)
		case 3:
			avgRainfall := climateService.AverageRainfall()
			fmt.Printf("Average Rainfall: %.2f mm\n", avgRainfall)
		case 4:
			var threshold float64
			fmt.Print("Enter rainfall threshold (mm): ")
			fmt.Scan(&threshold)
			filtered := climateService.FilterByRainfall(threshold)
			if len(filtered) == 0 {
				fmt.Println("No cities found with rainfall above the threshold.")
			} else {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filtered {
					fmt.Printf("%s: %.2f mm\n", city.City, city.Rainfall)
				}
			}
		case 5:
			var cityName string
			fmt.Print("Enter city name: ")
			fmt.Scan(&cityName)
			city, err := climateService.SearchByCity(cityName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("City: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.City, city.Temperature, city.Rainfall)
			}
		case 6:
			fmt.Println("Exiting Climate Data Analysis.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
