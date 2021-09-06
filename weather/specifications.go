package weather

import "fmt"

/*
	https://dev.netatmo.com/apidocumentation/weather#specifications
*/

const (
	// UnitWind is the unit used for the wind in API responses
	UnitWind = "kph"
	// UnitPressure is the unit used for pressure in API responses
	UnitPressure = "mbar"
	// UnitTemperature is the unit used for temperature in API responses
	UnitTemperature = "°C"
	// UnitCO2 is the unit used for the co2 concentration in API responses
	UnitCO2 = "ppm"
	// UnitHumidity is the unit used for humidity in API responses
	UnitHumidity = "%"
	// UnitNoise is the unit used for noise level in API responses
	UnitNoise = "dB"
)

// AnemometerBatteryStatus represents the battery status of the anemometer battery
type AnemometerBatteryStatus int

const (
	// AnemometerBatteryMax represents the maximum level of the anemometer battery
	AnemometerBatteryMax AnemometerBatteryStatus = 6000
	// AnemometerBatteryFull represents the full level of the anemometer battery
	AnemometerBatteryFull AnemometerBatteryStatus = 5590
	// AnemometerBatteryHigh represents the high level of the anemometer battery
	AnemometerBatteryHigh AnemometerBatteryStatus = 5180
	// AnemometerBatteryMedium represents the medium level of the anemometer battery
	AnemometerBatteryMedium AnemometerBatteryStatus = 4770
	// AnemometerBatteryLow represents the low level of the anemometer battery
	AnemometerBatteryLow AnemometerBatteryStatus = 4360
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

// WiFiQuality represents the WiFi strength signal
type WiFiQuality int

const (
	// WiFiQualityBad represents a bad level for WiFi reception
	WiFiQualityBad WiFiQuality = 86
	// WiFiQualityAverage represents an average level for WiFi reception
	WiFiQualityAverage WiFiQuality = 71
	// WiFiQUalityGood represents a good level for WiFi reception
	WiFiQUalityGood WiFiQuality = 56
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

// ModulesBatteryStatus represents the battery status for additionnal modules
type ModulesBatteryStatus int

const (
	// ModulesBatteryMax represents the maximum level of a module battery
	ModulesBatteryMax ModulesBatteryStatus = 6000
	// ModulesBatteryFull represents the full level of a module battery
	ModulesBatteryFull ModulesBatteryStatus = 5500
	// ModulesBatteryHigh represents the high level of a module battery
	ModulesBatteryHigh ModulesBatteryStatus = 5000
	// ModulesBatteryMedium represents the high level of a module battery
	ModulesBatteryMedium ModulesBatteryStatus = 4550
	// ModulesBatteryLow represents the low level of a module battery
	ModulesBatteryLow ModulesBatteryStatus = 4000
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

// ModuleType represents the type of an additionnal module
type ModuleType string

const (
	// ModuleTypeStation represents the main status
	ModuleTypeStation ModuleType = "NAMain"
	// ModuleTypeOutdoor represents the outdoor module
	ModuleTypeOutdoor ModuleType = "NAModule1"
	// ModuleTypeAnemometer represents the anemometer module
	ModuleTypeAnemometer ModuleType = "NAModule2"
	// ModuleTypeRainGauge represents the rain gauge module
	ModuleTypeRainGauge ModuleType = "NAModule3"
	// ModuleTypeIndoor represents an indoor module
	ModuleTypeIndoor ModuleType = "NAModule4"
)

// ModuleDataType represents a given data type for a module
type ModuleDataType string

const (
	// ModuleDataTypeTemperature represents a temperature data type
	ModuleDataTypeTemperature ModuleDataType = "Temperature"
	// ModuleDataTypeCO2 represents a co2 type
	ModuleDataTypeCO2 ModuleDataType = "CO2"
	// ModuleDataTypeHumidity represents a humidity data type
	ModuleDataTypeHumidity ModuleDataType = "Humidity"
	// ModuleDataTypeNoise represents a noise data type
	ModuleDataTypeNoise ModuleDataType = "Noise"
	// ModuleDataTypePressure represents a pressure data type
	ModuleDataTypePressure ModuleDataType = "Pressure"
	// ModuleDataTypeWind represents a wind data type
	ModuleDataTypeWind ModuleDataType = "Wind"
	// ModuleDataTypeRain represents a rain data type
	ModuleDataTypeRain ModuleDataType = "Rain"
)

// RadioQuality represents the radio signal quality between a station and its modules
type RadioQuality int

const (
	// RadioQualityLow represents a low quality radio link
	RadioQualityLow RadioQuality = 90
	// RadioQualityMedium represents a medium quality radio link
	RadioQualityMedium RadioQuality = 80 // custom
	// RadioQualityHigh represents a high quality radio link
	RadioQualityHigh RadioQuality = 70 // custom
	// RadioQUalityHighest represents the highest quality radio link
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

// Trend represents a current trend
type Trend string

const (
	// TrendUp represents a trend going up
	TrendUp Trend = "up"
	// TrendDown represents a trend going down
	TrendDown Trend = "down"
	// TrendStable represents a stable trend
	TrendStable Trend = "stable"
)
