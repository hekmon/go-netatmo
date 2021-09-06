package weather

import "fmt"

// UserWeather struct for UserWeather
type UserWeather struct {
	Mail           string                    `json:"mail"`
	Administrative UserWeatherAdministrative `json:"administrative"`
}

// UserWeatherAdministrative struct for UserWeatherAdministrative
type UserWeatherAdministrative struct {
	RegLocale    string           `json:"reg_locale"`     // user regional preferences (used for displaying date)
	Lang         string           `json:"lang"`           // user locale
	Country      string           `json:"country"`        // country
	Unit         UserUnitSystem   `json:"unit"`           // metric system, imperial system
	Windunit     UserUnitWind     `json:"windunit"`       // kph, mph, ms, beaufort, knot
	Pressureunit UserUnitPressure `json:"pressureunit"`   // mbar, inHg, mmHg
	FeelLikeAlgo UserUnitFeelLike `json:"feel_like_algo"` // algorithm used to compute feel like temperature: humidex, heat-index
}

// UserUnitSystem represents the unit system the user is using
type UserUnitSystem int

const (
	// UserUnitSystemMetric represents the international Metric system
	UserUnitSystemMetric UserUnitSystem = 0
	// UserUnitSystemImperial represents the Imperial metric
	UserUnitSystemImperial UserUnitSystem = 1
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (uus UserUnitSystem) String() string {
	switch uus {
	case UserUnitSystemMetric:
		return "metric"
	case UserUnitSystemImperial:
		return "imperial"
	default:
		return "<unknown>"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (uus UserUnitSystem) GoString() string {
	return fmt.Sprintf("%s (%d)", uus.String(), uus)
}

// UserUnitWind represents the unit chosen by the user to represents the wind speed
type UserUnitWind int

const (
	// UserUnitWindKph represents the Kilometer per Hour unit
	UserUnitWindKph UserUnitWind = 0
	// UserUnitWindMph represents the Miles per Hour unit
	UserUnitWindMph UserUnitWind = 1
	// UserUnitWindMs represents the Meter per Hour unit
	UserUnitWindMs UserUnitWind = 2
	// UserUnitWindBeaufort represents the Beaufort unit
	UserUnitWindBeaufort UserUnitWind = 3
	// UserUnitWindKnot represents the Knot unit
	UserUnitWindKnot UserUnitWind = 4
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (uuw UserUnitWind) String() string {
	switch uuw {
	case UserUnitWindKph:
		return "kph"
	case UserUnitWindMph:
		return "mph"
	case UserUnitWindMs:
		return "ms"
	case UserUnitWindBeaufort:
		return "beaufort"
	case UserUnitWindKnot:
		return "knot"
	default:
		return "<unknown>"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (uuw UserUnitWind) GoString() string {
	return fmt.Sprintf("%s (%d)", uuw.String(), uuw)
}

type UserUnitPressure int

const (
	UserUnitPressureMbar UserUnitPressure = 0
	UserUnitPressureInHg UserUnitPressure = 1
	UserUnitPressureMmHg UserUnitPressure = 2
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (uup UserUnitPressure) String() string {
	switch uup {
	case UserUnitPressureMbar:
		return "mbar"
	case UserUnitPressureInHg:
		return "inHg"
	case UserUnitPressureMmHg:
		return "mmHg"
	default:
		return "<unknown>"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (uup UserUnitPressure) GoString() string {
	return fmt.Sprintf("%s (%d)", uup.String(), uup)
}

type UserUnitFeelLike int

const (
	UserUnitFeelLikeHumidex   UserUnitFeelLike = 0
	UserUnitFeelLikeHeatIndex UserUnitFeelLike = 1
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (uufl UserUnitFeelLike) String() string {
	switch uufl {
	case UserUnitFeelLikeHumidex:
		return "humidex"
	case UserUnitFeelLikeHeatIndex:
		return "heat-index"
	default:
		return "<unknown>"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (uufl UserUnitFeelLike) GoString() string {
	return fmt.Sprintf("%s (%d)", uufl.String(), uufl)
}
