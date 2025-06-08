package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Temperature float64 `json:"temperature"`
	Location    string  `json:"location"`
	SensorID    string  `json:"sensor_id"`
	Unit        string  `json:"unit"`
	Timestamp   string  `json:"timestamp"`
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Temperature API server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	// Generate random temperature between 15 and 30 degrees Celsius
	temperature := 15.0 + rand.Float64()*15.0
	
	// Round to 2 decimal places
	temperature = float64(int(temperature*100)) / 100.0

	response := TemperatureResponse{
		Temperature: temperature,
		Location:    location,
		SensorID:    sensorID,
		Unit:        "°C",
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "healthy", "service": "temperature-api"}`)
} 