package utils

import (
	"fmt"
	"math"
	"strings"
)

func CalculateDistance(lat1, long1, lat2, long2 float64) float64 {
	if (lat1 != 0 && long1 != 0) && (lat2 != 0 && long2 != 0) {
		return CalLatitudeLongitudeDistance(lat1, long1, lat2, long2)
	}
	return 0
}

// Haversine function to calculate the distance between two points given their latitude and longitude
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Radius of the Earth in kilometers

	// Convert latitude and longitude from degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Differences in coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := R * c

	return distance
}

func CalLatitudeLongitudeDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius of the Earth in kilometers

	// Convert degrees to radians
	lat1Rad := degreesToRadians(lat1)
	lon1Rad := degreesToRadians(lon1)
	lat2Rad := degreesToRadians(lat2)
	lon2Rad := degreesToRadians(lon2)

	// Haversine formula
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := R * c
	// Check if the number of decimal places is more than 3
	if numDecimals(distance) > 3 {
		// Round to 3 decimal places
		distance = RoundToThreeDecimals(distance)
	}

	return distance
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// RoundToThreeDecimals rounds a float64 to 3 decimal places
func RoundToThreeDecimals(num float64) float64 {
	return math.Round(num*1000) / 1000
}

// numDecimals counts the number of decimal places in a float64
func numDecimals(num float64) int {
	str := fmt.Sprintf("%f", num)
	parts := strings.Split(str, ".")
	if len(parts) < 2 {
		return 0
	}
	return len(strings.TrimRight(parts[1], "0"))
}
