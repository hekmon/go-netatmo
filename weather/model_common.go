package weather

import "fmt"

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

/*
	TODO organize
// */

// // BodyMeasureGet struct for BodyMeasureGet
// type BodyMeasureGet struct {
// 	Body []MeasureGet `json:"body,omitempty"`
// }

// // MeasureGet struct for MeasureGet
// type MeasureGet struct {
// 	BegTime  float32   `json:"beg_time,omitempty"`
// 	StepTime float32   `json:"step_time,omitempty"`
// 	Value    []float32 `json:"value,omitempty"`
// }

// // InvalidServerResponse struct for InvalidServerResponse
// type InvalidServerResponse struct {
// 	Error InvalidServerResponseError `json:"error"`
// }

// // InvalidServerResponseError struct for InvalidServerResponseError
// type InvalidServerResponseError struct {
// 	Code    float32 `json:"code"`
// 	Message string  `json:"message"`
// }

// // MeasureMacAddressNAMain struct for MeasureMacAddressNAMain
// type MeasureMacAddressNAMain struct {
// 	Res  *MeasureMacAddressNAMainRes `json:"res,omitempty"`
// 	Type *string                     `json:"type,omitempty"`
// }

// // MeasureMacAddressNAMainRes struct for MeasureMacAddressNAMainRes
// type MeasureMacAddressNAMainRes struct {
// 	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
// }

// // MeasureMacAddressNAModule1 struct for MeasureMacAddressNAModule1
// type MeasureMacAddressNAModule1 struct {
// 	Res  *MeasureMacAddressNAModule1Res `json:"res,omitempty"`
// 	Type *[]string                      `json:"type,omitempty"`
// }

// // MeasureMacAddressNAModule1Res struct for MeasureMacAddressNAModule1Res
// type MeasureMacAddressNAModule1Res struct {
// 	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
// }

// // PublicData struct for PublicData
// type PublicData struct {
// 	Body *[]PublicStationData `json:"body,omitempty"`
// }

// // ServerResponse struct for ServerResponse
// type ServerResponse struct {
// 	Status     *string `json:"status,omitempty"`
// 	TimeExec   *string `json:"time_exec,omitempty"`
// 	TimeServer *string `json:"time_server,omitempty"`
// }
