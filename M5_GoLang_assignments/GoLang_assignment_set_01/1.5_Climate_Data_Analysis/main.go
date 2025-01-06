package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// ClimateData represents weather data for a city
type ClimateData struct {
    CityName     string
    Temperature  float64 // in Celsius
    Rainfall     float64 // in millimeters
}

// ClimateAnalyzer handles climate data analysis operations
type ClimateAnalyzer struct {
    cities   []ClimateData
    scanner  *bufio.Scanner
}

// NewClimateAnalyzer creates a new instance with sample data
func NewClimateAnalyzer() *ClimateAnalyzer {
    // Sample climate data for Indian cities
    sampleData := []ClimateData{
        {"Mumbai", 27.2, 2166.0},
        {"Delhi", 25.0, 797.0},
        {"Bangalore", 24.1, 970.0},
        {"Chennai", 28.6, 1400.0},
        {"Kolkata", 26.8, 1582.0},
        {"Hyderabad", 26.7, 812.0},
        {"Pune", 24.9, 722.0},
        {"Jaipur", 25.8, 650.0},
    }

    return &ClimateAnalyzer{
        cities:  sampleData,
        scanner: bufio.NewScanner(os.Stdin),
    }
}

// AddCity adds a new city to the dataset
func (ca *ClimateAnalyzer) AddCity(name string, temp, rainfall float64) error {
    // Validate input
    if name == "" {
        return errors.New("city name cannot be empty")
    }
    if temp < -50 || temp > 50 {
        return fmt.Errorf("invalid temperature: %.2f°C (must be between -50°C and 50°C)", temp)
    }
    if rainfall < 0 {
        return fmt.Errorf("invalid rainfall: %.2f mm (must be non-negative)", rainfall)
    }

    // Check for duplicate city
    for _, city := range ca.cities {
        if strings.EqualFold(city.CityName, name) {
            return fmt.Errorf("city '%s' already exists in the dataset", name)
        }
    }

    newCity := ClimateData{
        CityName:    name,
        Temperature: temp,
        Rainfall:    rainfall,
    }
    ca.cities = append(ca.cities, newCity)
    return nil
}

// FindHighestTemperature returns the city with the highest temperature
func (ca *ClimateAnalyzer) FindHighestTemperature() (*ClimateData, error) {
    if len(ca.cities) == 0 {
        return nil, errors.New("no cities in dataset")
    }

    highest := &ca.cities[0]
    for i := range ca.cities {
        if ca.cities[i].Temperature > highest.Temperature {
            highest = &ca.cities[i]
        }
    }
    return highest, nil
}

// FindLowestTemperature returns the city with the lowest temperature
func (ca *ClimateAnalyzer) FindLowestTemperature() (*ClimateData, error) {
    if len(ca.cities) == 0 {
        return nil, errors.New("no cities in dataset")
    }

    lowest := &ca.cities[0]
    for i := range ca.cities {
        if ca.cities[i].Temperature < lowest.Temperature {
            lowest = &ca.cities[i]
        }
    }
    return lowest, nil
}

// CalculateAverageRainfall computes the mean rainfall across all cities
func (ca *ClimateAnalyzer) CalculateAverageRainfall() (float64, error) {
    if len(ca.cities) == 0 {
        return 0, errors.New("no cities in dataset")
    }

    var total float64
    for _, city := range ca.cities {
        total += city.Rainfall
    }
    return total / float64(len(ca.cities)), nil
}

// FilterCitiesByRainfall returns cities with rainfall above the threshold
func (ca *ClimateAnalyzer) FilterCitiesByRainfall(threshold float64) []ClimateData {
    var filteredCities []ClimateData
    for _, city := range ca.cities {
        if city.Rainfall > threshold {
            filteredCities = append(filteredCities, city)
        }
    }
    
    // Sort by rainfall in descending order
    sort.Slice(filteredCities, func(i, j int) bool {
        return filteredCities[i].Rainfall > filteredCities[j].Rainfall
    })
    
    return filteredCities
}

// SearchCity finds a city by name (case-insensitive)
func (ca *ClimateAnalyzer) SearchCity(name string) (*ClimateData, error) {
    if name == "" {
        return nil, errors.New("city name cannot be empty")
    }

    for i := range ca.cities {
        if strings.EqualFold(ca.cities[i].CityName, name) {
            return &ca.cities[i], nil
        }
    }
    return nil, fmt.Errorf("city '%s' not found", name)
}

// DisplayAllCities shows all cities in a formatted table
func (ca *ClimateAnalyzer) DisplayAllCities() {
    fmt.Printf("\n%-15s | %12s | %12s\n", "City", "Temperature", "Rainfall")
    fmt.Println(strings.Repeat("-", 45))
    
    for _, city := range ca.cities {
        fmt.Printf("%-15s | %9.1f°C | %8.1f mm\n",
            city.CityName, city.Temperature, city.Rainfall)
    }
    fmt.Println()
}

func main() {
    analyzer := NewClimateAnalyzer()
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nClimate Data Analysis System")
        fmt.Println("1. Display All Cities")
        fmt.Println("2. Add New City")
        fmt.Println("3. Find Temperature Extremes")
        fmt.Println("4. Calculate Average Rainfall")
        fmt.Println("5. Filter Cities by Rainfall")
        fmt.Println("6. Search City")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice (1-7): ")

        scanner.Scan()
        choice := strings.TrimSpace(scanner.Text())

        switch choice {
        case "1":
            analyzer.DisplayAllCities()

        case "2":
            fmt.Print("Enter city name: ")
            scanner.Scan()
            name := strings.TrimSpace(scanner.Text())

            fmt.Print("Enter temperature (°C): ")
            scanner.Scan()
            temp, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid temperature value")
                continue
            }

            fmt.Print("Enter rainfall (mm): ")
            scanner.Scan()
            rainfall, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid rainfall value")
                continue
            }

            if err := analyzer.AddCity(name, temp, rainfall); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("City added successfully!")
            }

        case "3":
            highest, err := analyzer.FindHighestTemperature()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nHighest Temperature: %s (%.1f°C)\n",
                    highest.CityName, highest.Temperature)
            }

            lowest, err := analyzer.FindLowestTemperature()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Lowest Temperature: %s (%.1f°C)\n",
                    lowest.CityName, lowest.Temperature)
            }

        case "4":
            avg, err := analyzer.CalculateAverageRainfall()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nAverage Rainfall across all cities: %.1f mm\n", avg)
            }

        case "5":
            fmt.Print("Enter rainfall threshold (mm): ")
            scanner.Scan()
            threshold, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid threshold value")
                continue
            }

            cities := analyzer.FilterCitiesByRainfall(threshold)
            if len(cities) == 0 {
                fmt.Printf("No cities found with rainfall above %.1f mm\n", threshold)
            } else {
                fmt.Printf("\nCities with rainfall above %.1f mm:\n", threshold)
                for _, city := range cities {
                    fmt.Printf("%-15s: %.1f mm\n", city.CityName, city.Rainfall)
                }
            }

        case "6":
            fmt.Print("Enter city name to search: ")
            scanner.Scan()
            name := strings.TrimSpace(scanner.Text())

            city, err := analyzer.SearchCity(name)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nCity: %s\n", city.CityName)
                fmt.Printf("Temperature: %.1f°C\n", city.Temperature)
                fmt.Printf("Rainfall: %.1f mm\n", city.Rainfall)
            }

        case "7":
            fmt.Println("Thank you for using the Climate Data Analysis System!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}