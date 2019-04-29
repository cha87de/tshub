package datahub

import "time"

const (
	resolution1min  = time.Duration(1) * time.Minute
	resolution30min = time.Duration(30) * time.Minute
	resolution1h    = time.Duration(1) * time.Hour
	resolution12h   = time.Duration(12) * time.Hour
	resolution24h   = time.Duration(24) * time.Hour
)

var resolutions = [...]time.Duration{
	resolution1min,
	resolution30min,
	resolution1h,
	resolution12h,
	resolution24h,
}
