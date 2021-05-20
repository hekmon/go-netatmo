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

type UserUnitSystem int

const (
	UserUnitSystemMetric   UserUnitSystem = 0
	UserUnitSystemImperial UserUnitSystem = 1
)

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

func (uus UserUnitSystem) GoString() string {
	return fmt.Sprintf("%s (%d)", uus.String(), uus)
}

type UserUnitWind int

const (
	UserUnitWindKph      UserUnitWind = 0
	UserUnitWindMph      UserUnitWind = 1
	UserUnitWindMs       UserUnitWind = 2
	UserUnitWindBeaufort UserUnitWind = 3
	UserUnitWindKnot     UserUnitWind = 4
)

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

func (uuw UserUnitWind) GoString() string {
	return fmt.Sprintf("%s (%d)", uuw.String(), uuw)
}

type UserUnitPressure int

const (
	UserUnitPressureMbar UserUnitPressure = 0
	UserUnitPressureInHg UserUnitPressure = 1
	UserUnitPressureMmHg UserUnitPressure = 2
)

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

func (uup UserUnitPressure) GoString() string {
	return fmt.Sprintf("%s (%d)", uup.String(), uup)
}

type UserUnitFeelLike int

const (
	UserUnitFeelLikeHumidex   UserUnitFeelLike = 0
	UserUnitFeelLikeHeatIndex UserUnitFeelLike = 1
)

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

func (uufl UserUnitFeelLike) GoString() string {
	return fmt.Sprintf("%s (%d)", uufl.String(), uufl)
}
