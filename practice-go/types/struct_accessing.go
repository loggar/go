package main

import (
	"fmt"
	"time"
)

// Bootcamp2 is a struct ...
type Bootcamp2 struct {
	Lat, Lon float64
	Date     time.Time
}

func structAccessing() {
	event := Bootcamp2{
		Lat: 34.012836,
		Lon: -118.495338,
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)
}
