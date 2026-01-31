package utils

import (
	"encoding/base64"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/kpi-studio/go-strava-api/models"
)

// Distance conversion helpers

// MetersToMiles converts meters to miles
func MetersToMiles(meters float64) float64 {
	return meters * 0.000621371
}

// MetersToKilometers converts meters to kilometers
func MetersToKilometers(meters float64) float64 {
	return meters / 1000
}

// MilesToMeters converts miles to meters
func MilesToMeters(miles float64) float64 {
	return miles * 1609.34
}

// KilometersToMeters converts kilometers to meters
func KilometersToMeters(km float64) float64 {
	return km * 1000
}

// Speed conversion helpers

// MetersPerSecondToMilesPerHour converts m/s to mph
func MetersPerSecondToMilesPerHour(mps float64) float64 {
	return mps * 2.23694
}

// MetersPerSecondToKilometersPerHour converts m/s to km/h
func MetersPerSecondToKilometersPerHour(mps float64) float64 {
	return mps * 3.6
}

// Pace calculation helpers

// CalculatePacePerMile returns the pace per mile in seconds
func CalculatePacePerMile(distanceMeters float64, timeSeconds int) int {
	miles := MetersToMiles(distanceMeters)
	if miles == 0 {
		return 0
	}
	return int(float64(timeSeconds) / miles)
}

// CalculatePacePerKilometer returns the pace per kilometer in seconds
func CalculatePacePerKilometer(distanceMeters float64, timeSeconds int) int {
	km := MetersToKilometers(distanceMeters)
	if km == 0 {
		return 0
	}
	return int(float64(timeSeconds) / km)
}

// FormatPace formats a pace (seconds per distance unit) as MM:SS
func FormatPace(paceSeconds int) string {
	minutes := paceSeconds / 60
	seconds := paceSeconds % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

// FormatDuration formats a duration in seconds as HH:MM:SS
func FormatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, secs)
	}
	return fmt.Sprintf("%d:%02d", minutes, secs)
}

// ParseDuration parses a duration string (HH:MM:SS or MM:SS) to seconds
func ParseDuration(duration string) int {
	parts := strings.Split(duration, ":")
	switch len(parts) {
	case 2: // MM:SS
		var minutes, seconds int
		fmt.Sscanf(duration, "%d:%d", &minutes, &seconds)
		return minutes*60 + seconds
	case 3: // HH:MM:SS
		var hours, minutes, seconds int
		fmt.Sscanf(duration, "%d:%d:%d", &hours, &minutes, &seconds)
		return hours*3600 + minutes*60 + seconds
	default:
		return 0
	}
}

// Weight conversion helpers

// KilogramsToPounds converts kilograms to pounds
func KilogramsToPounds(kg float64) float64 {
	return kg * 2.20462
}

// PoundsToKilograms converts pounds to kilograms
func PoundsToKilograms(lbs float64) float64 {
	return lbs / 2.20462
}

// Polyline decoding

// DecodePolyline decodes a Google polyline string to coordinates
func DecodePolyline(encoded string) [][]float64 {
	var coordinates [][]float64
	index := 0
	lat := 0
	lng := 0

	for index < len(encoded) {
		var b, shift, result int

		// Decode latitude
		shift = 0
		result = 0
		for {
			b = int(encoded[index]) - 63
			index++
			result |= (b & 0x1f) << shift
			shift += 5
			if b < 0x20 {
				break
			}
		}
		dlat := result
		if (result & 1) != 0 {
			dlat = ^(result >> 1)
		} else {
			dlat = result >> 1
		}
		lat += dlat

		// Decode longitude
		shift = 0
		result = 0
		for {
			b = int(encoded[index]) - 63
			index++
			result |= (b & 0x1f) << shift
			shift += 5
			if b < 0x20 {
				break
			}
		}
		dlng := result
		if (result & 1) != 0 {
			dlng = ^(result >> 1)
		} else {
			dlng = result >> 1
		}
		lng += dlng

		coordinates = append(coordinates, []float64{
			float64(lat) / 1e5,
			float64(lng) / 1e5,
		})
	}

	return coordinates
}

// EncodePolyline encodes coordinates to a Google polyline string
func EncodePolyline(coordinates [][]float64) string {
	var encoded strings.Builder
	prevLat := 0
	prevLng := 0

	for _, coord := range coordinates {
		lat := int(math.Round(coord[0] * 1e5))
		lng := int(math.Round(coord[1] * 1e5))

		dlat := lat - prevLat
		dlng := lng - prevLng

		prevLat = lat
		prevLng = lng

		encodeValue(&encoded, dlat)
		encodeValue(&encoded, dlng)
	}

	return encoded.String()
}

func encodeValue(encoded *strings.Builder, value int) {
	if value < 0 {
		value = ^(value << 1)
	} else {
		value = value << 1
	}

	for value >= 0x20 {
		encoded.WriteByte(byte((0x20 | (value & 0x1f)) + 63))
		value >>= 5
	}
	encoded.WriteByte(byte(value + 63))
}

// Distance calculation

// CalculateDistance calculates the distance between two coordinates using the Haversine formula
func CalculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadius = 6371000 // meters

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLng := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLng/2)*math.Sin(deltaLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// Time helpers

// TimeToUnix converts a time to Unix timestamp
func TimeToUnix(t time.Time) int {
	return int(t.Unix())
}

// UnixToTime converts a Unix timestamp to time
func UnixToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

// Activity type helpers

// IsRideActivity checks if the activity type is a ride
func IsRideActivity(activityType models.ActivityType) bool {
	switch activityType {
	case models.ActivityTypeRide, models.ActivityTypeEBikeRide, models.ActivityTypeVirtualRide, models.ActivityTypeHandcycle:
		return true
	default:
		return false
	}
}

// IsRunActivity checks if the activity type is a run
func IsRunActivity(activityType models.ActivityType) bool {
	switch activityType {
	case models.ActivityTypeRun, models.ActivityTypeVirtualRun, models.ActivityTypeWalk, models.ActivityTypeHike:
		return true
	default:
		return false
	}
}

// IsSwimActivity checks if the activity type is a swim
func IsSwimActivity(activityType models.ActivityType) bool {
	return activityType == models.ActivityTypeSwim
}

// IsWinterActivity checks if the activity type is a winter sport
func IsWinterActivity(activityType models.ActivityType) bool {
	switch activityType {
	case models.ActivityTypeAlpineSki, models.ActivityTypeBackcountrySki, models.ActivityTypeNordicSki,
		 models.ActivityTypeSnowboard, models.ActivityTypeSnowshoe, models.ActivityTypeIceSkate:
		return true
	default:
		return false
	}
}

// IsWaterActivity checks if the activity type is a water sport
func IsWaterActivity(activityType models.ActivityType) bool {
	switch activityType {
	case models.ActivityTypeSwim, models.ActivityTypeCanoeing, models.ActivityTypeKayaking,
		 models.ActivityTypeKitesurf, models.ActivityTypeRowing, models.ActivityTypeSail,
		 models.ActivityTypeStandUpPaddling, models.ActivityTypeSurf, models.ActivityTypeWindsurf:
		return true
	default:
		return false
	}
}

// Grade helpers

// CalculateGrade calculates the grade percentage between two points
func CalculateGrade(elevationGain float64, distance float64) float64 {
	if distance == 0 {
		return 0
	}
	return (elevationGain / distance) * 100
}

// Power helpers

// CalculateNormalizedPower calculates normalized power from a power stream
func CalculateNormalizedPower(powers []int) float64 {
	if len(powers) == 0 {
		return 0
	}

	// Calculate 30-second rolling average
	windowSize := 30
	var rollingAvgs []float64

	for i := 0; i <= len(powers)-windowSize; i++ {
		sum := 0
		for j := i; j < i+windowSize; j++ {
			sum += powers[j]
		}
		rollingAvgs = append(rollingAvgs, float64(sum)/float64(windowSize))
	}

	// Calculate fourth power mean
	var sum float64
	for _, avg := range rollingAvgs {
		sum += math.Pow(avg, 4)
	}

	if len(rollingAvgs) == 0 {
		return 0
	}

	return math.Pow(sum/float64(len(rollingAvgs)), 0.25)
}

// CalculateIntensityFactor calculates the intensity factor
func CalculateIntensityFactor(normalizedPower float64, ftp int) float64 {
	if ftp == 0 {
		return 0
	}
	return normalizedPower / float64(ftp)
}

// CalculateTSS calculates Training Stress Score
func CalculateTSS(normalizedPower float64, intensityFactor float64, durationSeconds int) float64 {
	hours := float64(durationSeconds) / 3600
	return (hours * normalizedPower * intensityFactor * 100) / float64(ftp)
}

var ftp int // This should be set from athlete's FTP

// Base64 encoding for file uploads

// EncodeFileToBase64 encodes file content to base64
func EncodeFileToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64ToFile decodes base64 to file content
func DecodeBase64ToFile(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}