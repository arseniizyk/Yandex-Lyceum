package main

type TripParameters struct {
	Distance float64
	Duration float64
}

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

func CalculateBasePrice(trip TripParameters) float64 {
	return trip.Distance*pricePerKm + trip.Duration*pricePerMinute
}
