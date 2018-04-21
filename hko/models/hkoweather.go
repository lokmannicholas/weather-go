package models

import "time"

type HKOWeather struct {
	Temperature float64   `json:"temp"`
	Humidity    float64   `json:"humidity"`
	Direction   string    `json:"direction"`
	Location    string    `json:"location"`
	Speed       float64   `json:"speed"`
	Gust        float64   `json:"gust"`
	Time        time.Time `json:"time"`
}
