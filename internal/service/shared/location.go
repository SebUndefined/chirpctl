package shared

import (
	"fmt"
	"strings"
)

type LocationSource string

const (
	LocationSourceUnknown         LocationSource = "UNKNOWN"
	LocationSourceGPS             LocationSource = "GPS"
	LocationSourceConfig          LocationSource = "CONFIG"
	LocationSourceGeoResolverTDOA LocationSource = "GEO_RESOLVER_TDOA"
	LocationSourceGeoResolverRSSI LocationSource = "GEO_RESOLVER_RSSI"
	LocationSourceGeoResolverGNSS LocationSource = "GEO_RESOLVER_GNSS"
	LocationSourceGeoResolverWIFI LocationSource = "GEO_RESOLVER_WIFI"
)

func (s LocationSource) String() string {
	if len(strings.TrimSpace(string(s))) == 0 {
		return LocationSourceUnknown.String()
	}
	return string(s)
}

func ParseLocationSourceFromInt32(l int32) LocationSource {
	switch l {
	case 1:
		return LocationSourceGPS
	case 2:
		return LocationSourceConfig
	case 3:
		return LocationSourceGeoResolverTDOA
	case 4:
		return LocationSourceGeoResolverRSSI
	case 5:
		return LocationSourceGeoResolverGNSS
	case 6:
		return LocationSourceGeoResolverWIFI
	default:
		return LocationSourceUnknown
	}
}

type Location struct {
	Latitude  float64        `json:"latitude,omitempty"`
	Longitude float64        `json:"longitude,omitempty"`
	Altitude  float64        `json:"altitude,omitempty"`
	Source    LocationSource `json:"source,omitempty"`
	Accuracy  float32        `json:"accuracy,omitempty"`
}

func (l Location) String() string {
	return fmt.Sprintf("%.6f, %.6f (alt: %.1f m, source: %s)",
		l.Latitude, l.Longitude, l.Altitude, l.Source)
}
