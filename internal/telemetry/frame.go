package telemetry

import (
	"time"
)

type Frame struct {
	At       time.Time
	TrackPos float64
	Speed    float64
	Throttle float64
	Brake    float64
	Steer    float64
	Gear     int
	Lap      int
	LapTime  time.Duration
}
