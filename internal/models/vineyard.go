package models

import "time"

type Vineyard struct {
	ID             int          `json:"id"`
	XCoord         float64      `json:"xCoord"`
	YCoord         float64      `json:"yCoord"`
	WorkHistory    []WorkObject `json:"workHistory"`
	WeatherHistory string       `json:"weatherHistory"`
	GroundHistory  []GroundAnalysis
}

type WorkObject struct {
	Date          time.Time
	WorkType      string
	VineYardGrade int
}

type WeatherObject struct {
	Sunny       int
	Wet         int
	Temperature float32
}

type GroundAnalysis struct {
	Date       time.Time
	Organic    float32
	Acidity    float32
	NO3        float32
	Phosphorus float32
}
