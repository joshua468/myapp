package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Visitor"
	}

	clientIP := getClientIP(r)

	city := "New York"
	temperature := 11.0

	greeting := fmt.Sprintf("Hello, %s! The temperature is %.1f degrees Celsius in %s", visitorName, temperature, city)
	response := Response{
		ClientIP: clientIP,
		Location: city,
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
		return "unknown"
	}

	return ip
}
