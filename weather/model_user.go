package weather

// UserWeather struct for UserWeather
type UserWeather struct {
	Mail           string                    `json:"mail"`
	Administrative UserWeatherAdministrative `json:"administrative"`
}

// UserWeatherAdministrative struct for UserWeatherAdministrative
type UserWeatherAdministrative struct {
	RegLocale    string  `json:"reg_locale"`     // user regional preferences (used for displaying date)
	Lang         string  `json:"lang"`           // user locale
	Country      string  `json:"country"`        // country
	Unit         float32 `json:"unit"`           // 0 -> metric system, 1 -> imperial system
	Windunit     float32 `json:"windunit"`       // 0 -> kph, 1 -> mph, 2 -> ms, 3 -> beaufort, 4 -> knot
	Pressureunit float32 `json:"pressureunit"`   // 0 -> mbar, 1 -> inHg, 2 -> mmHg
	FeelLikeAlgo float32 `json:"feel_like_algo"` // algorithm used to compute feel like temperature, 0 -> humidex, 1 -> heat-index
}
