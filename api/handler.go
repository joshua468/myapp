package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Visitor"
	}

	clientIP := getClientIP(r)
	log.Println("Client IP:", clientIP)

	location, err := getLocationFromIP(clientIP)
	if err != nil {
		log.Println("Error fetching location:", err)
		http.Error(w, "Error fetching location", http.StatusInternalServerError)
		return
	}

	temperature, err := getTemperature(location)
	if err != nil {
		log.Println("Error fetching temperature:", err)
		http.Error(w, "Error fetching temperature", http.StatusInternalServerError)
		return
	}

	greeting := fmt.Sprintf("Hello, %s! The temperature is %.1f degrees Celsius in %s", visitorName, temperature, location)
	response := Response{
		ClientIP: clientIP,
		Location: location,
		Greeting: greeting,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getClientIP(r *http.Request) string {
	headersToCheck := []string{"X-Forwarded-For", "X-Real-IP"}

	for _, header := range headersToCheck {
		if ip := r.Header.Get(header); ip != "" {
			if header == "X-Forwarded-For" {
				ips := strings.Split(ip, ",")
				return strings.TrimSpace(ips[0])
			}
			return ip
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("Error splitting host port:", err)
		return "unknown"
	}

	if ip == "::1" || ip == "127.0.0.1" {
		ip = "8.8.8.8" // Example public IP (Google DNS)
	}

	return ip
}

func getLocationFromIP(ip string) (string, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request to IP API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		City string `json:"city"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response from IP API: %v", err)
	}

	if result.City == "" {
		return "", fmt.Errorf("city not found for IP: %s", ip)
	}

	return result.City, nil
}

func getTemperature(location string) (float64, error) {
	apiKey := os.Getenv("OPENWEATHERMAP_API")
	if apiKey == "" {
		return 0, fmt.Errorf("missing OpenWeatherMap API key")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", location, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error making request to OpenWeatherMap API: %v", err)
	}
	defer resp.Body.Close()

	var weatherData struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return 0, fmt.Errorf("error decoding response from OpenWeatherMap API: %v", err)
	}

	return weatherData.Main.Temp, nil
}
