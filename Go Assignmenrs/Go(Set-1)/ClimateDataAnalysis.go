package main

import "fmt"

//To store city data:
type CData struct {
	CityName    string
	AverageTemp float64
	Rainfall    float64
}

func main() {
	// Data for the cities:
	cities := []CData{
		{"New York", 15.3, 120.5},
		{"Los Angeles", 20.1, 45.0},
		{"Chicago", 10.5, 80.2},
		{"Miami", 25.4, 150.1},
		{"Seattle", 13.2, 120.0},
	}

	for {
		//Menu options
		fmt.Println("\n1. Highest Temperature")
		fmt.Println("2. Lowest Temperature")
		fmt.Println("3. Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		if option == 1 {
			// For high temperature:
			highest := cities[0]
			for _, city := range cities {
				if city.AverageTemp > highest.AverageTemp {
					highest = city
				}
			}
			fmt.Printf("Highest Temperature: %s (%.2f°C)\n", highest.CityName, highest.AverageTemp)

		} else if option == 2 {
	        //For low temperature:
			lowest := cities[0]
			for _, city := range cities {
				if city.AverageTemp < lowest.AverageTemp {
					lowest = city
				}
			}
			fmt.Printf("Lowest Temperature: %s (%.2f°C)\n", lowest.CityName, lowest.AverageTemp)

		} else if option == 3 {
	        //For total rainfall:
			var totalRainfall float64
			for _, city := range cities {
				totalRainfall += city.Rainfall
			}
			avgRainfall := totalRainfall / float64(len(cities))
			fmt.Printf("Average Rainfall: %.2f mm\n", avgRainfall)

		} else if option == 4 {
            //For rainfall threshold:
			fmt.Print("Enter rainfall threshold: ")
			var threshold float64
			_, err := fmt.Scanf("%f", &threshold)
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			for _, city := range cities {
				if city.Rainfall > threshold {
					fmt.Printf("%s: %.2f mm\n", city.CityName, city.Rainfall)
				}
			}

		} else if option == 5 {
	        //To get the name of city entered by user:
			fmt.Print("Enter city name to search: ")
			var cityName string
			fmt.Scanln(&cityName)
			found := false
			for _, city := range cities {
				if city.CityName == cityName {
					fmt.Printf("City: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.CityName, city.AverageTemp, city.Rainfall)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("City not found.")
			}

		} else if option == 6 {

			fmt.Println("Exiting. Goodbye!")
			break

		} else {
			fmt.Println("Invalid option. Try again.")
		}
	}
}
