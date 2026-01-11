// Package weather
// provides tools to build application on the weather.
package weather

var (
	// CurrentCondition express the condition.
	CurrentCondition string
	// CurrentLocation express the location.
	CurrentLocation string
)

// Forecast returns a string with the forecast of the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
