package main

import (
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	Name() string
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	estimatedTime := make(map[string]float64, len(vehicles))

	for _, v := range vehicles {
		name := reflect.TypeOf(v).String()
		estimatedTime[name] = v.CalculateTravelTime(distance)
	}

	return estimatedTime
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	if c.Speed <= 0 {
		return 0
	}
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	if m.Speed <= 0 {
		return 0
	}
	return distance / m.Speed
}

func (c Car) Name() string {
	return c.Type
}

func (m Motorcycle) Name() string {
	return m.Type
}
