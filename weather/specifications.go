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
	return fmt.Sprintf("%s (%d)", abs, abs)
}

type WiFiQuality int

const (
	WiFiQualityBad     WiFiQuality = 86
	WiFiQualityAverage WiFiQuality = 71
	WiFiQUalityGood    WiFiQuality = 56
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (wq WiFiQuality) String() string {
	switch {
	case wq <= WiFiQUalityGood:
		return "good"
	case wq <= WiFiQualityAverage:
		return "average"
	case wq <= WiFiQualityBad:
		return "bad"
	default:
		return "very bad"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (wq WiFiQuality) GoString() string {
	return fmt.Sprintf("%s (%d)", wq, wq)
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
	return fmt.Sprintf("%s (%d)", abs, abs)
}

/*
	Others modules specifications
*/

type ModuleType string

const (
	ModuleTypeStation    ModuleType = "NAMain"
	ModuleTypeOutdoor    ModuleType = "NAModule1"
	ModuleTypeAnemometer ModuleType = "NAModule2"
	ModuleTypeRainGauge  ModuleType = "NAModule3"
	ModuleTypeIndoor     ModuleType = "NAModule4"
)

type ModuleDataType string

const (
	ModuleDataTypeTemperature ModuleDataType = "Temperature"
	ModuleDataTypeCO2         ModuleDataType = "CO2"
	ModuleDataTypeHumidity    ModuleDataType = "Humidity"
	ModuleDataTypeNoise       ModuleDataType = "Noise"
	ModuleDataTypePressure    ModuleDataType = "Pressure"
	ModuleDataTypeWind        ModuleDataType = "Wind"
	ModuleDataTypeWRain       ModuleDataType = "Rain"
)

type RadioQuality int

const (
	RadioQualityLow     RadioQuality = 90
	RadioQualityMedium  RadioQuality = 80 // custom
	RadioQualityHigh    RadioQuality = 70 // custom
	RadioQUalityHighest RadioQuality = 60
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (rq RadioQuality) String() string {
	switch {
	case rq <= RadioQUalityHighest:
		return "highest"
	case rq <= RadioQualityHigh:
		return "high"
	case rq <= RadioQualityMedium:
		return "medium"
	case rq <= RadioQualityLow:
		return "low"
	default:
		return "very low"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (rq RadioQuality) GoString() string {
	return fmt.Sprintf("%s (%d)", rq, rq)
}

type Trend string

const (
	TrendUp     Trend = "up"
	TrendDown   Trend = "down"
	TrendStable Trend = "stable"
)
