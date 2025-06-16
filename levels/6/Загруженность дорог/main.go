package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	mult := 1.0

	if weather.WindSpeed > 15 {
		mult += 0.1
	}

	switch weather.Condition {
	case Clear:
		return mult
	case Rain:
		mult += 0.125
	case HeavyRain:
		mult += 0.2
	case Snow:
		mult += 0.15
	}

	return mult
}
