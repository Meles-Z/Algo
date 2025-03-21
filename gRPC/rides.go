package rides

import "time"

type Location struct {
	Lat  float64
	Long float64
}

type RideType byte

const (
	Regular RideType = iota + 1
	Pool
)

type Ride struct {
	ID           string
	DriverID     string
	Location     Location
	PassengerIds []string
	Start        time.Time
	End          time.Time
	Distance     float64
	Type         RideType
}
