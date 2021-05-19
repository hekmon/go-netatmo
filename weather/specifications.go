package weather

import "fmt"

/*
	https://dev.netatmo.com/apidocumentation/weather#specifications
*/

const (
	UnitWind        = "kph"
	UnitPressure    = "mbar"
	UnitTemperature = "°C"
	UnitCO3         = "ppm"
	UnitHumidity    = "%"
	UnitNoise       = "dB"
)

type AnemometerBatteryStatus int

const (
	AnemometerBatteryMax    AnemometerBatteryStatus = 6000
	AnemometerBatteryFull   AnemometerBatteryStatus = 5590
	AnemometerBatteryHigh   AnemometerBatteryStatus = 5180
	AnemometerBatteryMedium AnemometerBatteryStatus = 4770
	AnemometerBatteryLow    AnemometerBatteryStatus = 4360
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (abs AnemometerBatteryStatus) String() string {
	switch {
	case abs >= AnemometerBatteryMax:
		return "max"
	case abs >= AnemometerBatteryFull:
		return "full"
	case abs >= AnemometerBatteryHigh:
		return "high"
	case abs >= AnemometerBatteryMedium:
		return "medium"
	case abs >= AnemometerBatteryLow:
		return "low"
	default:
		return "very low"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (abs AnemometerBatteryStatus) GoString() string {
	return fmt.Sprintf("%s (%d)", abs.String(), abs)
}

type WiFiStatus int

const (
	WiFiQualityBad     WiFiStatus = 86
	WiFiQualityAverage WiFiStatus = 71
	WiFiQUalityGood    WiFiStatus = 56
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (ws WiFiStatus) String() string {
	switch {
	case ws <= WiFiQUalityGood:
		return "good"
	case ws <= WiFiQualityAverage:
		return "average"
	case ws <= WiFiQualityBad:
		return "bad"
	default:
		return "very bad"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (ws WiFiStatus) GoString() string {
	return fmt.Sprintf("%s (%d)", ws.String(), ws)
}

type ModulesBatteryStatus int

const (
	ModulesBatteryMax    ModulesBatteryStatus = 6000
	ModulesBatteryFull   ModulesBatteryStatus = 5500
	ModulesBatteryHigh   ModulesBatteryStatus = 5000
	ModulesBatteryMedium ModulesBatteryStatus = 4550
	ModulesBatteryLow    ModulesBatteryStatus = 4000
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (abs ModulesBatteryStatus) String() string {
	switch {
	case abs >= ModulesBatteryMax:
		return "max"
	case abs >= ModulesBatteryFull:
		return "full"
	case abs >= ModulesBatteryHigh:
		return "high"
	case abs >= ModulesBatteryMedium:
		return "medium"
	case abs >= ModulesBatteryLow:
		return "low"
	default:
		return "very low"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (abs ModulesBatteryStatus) GoString() string {
	return fmt.Sprintf("%s (%d)", abs.String(), abs)
}
