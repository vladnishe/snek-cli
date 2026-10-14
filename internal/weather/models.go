package weather

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp          float64 `json:"temp"`
		TempFeelsLike float64 `json:"feels_like"`
		TempMin       float64 `json:"temp_min"`
		TempMax       float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Clouds struct {
		Percentage uint32 `json:"all"`
	} `json:"clouds"`
	Rain struct {
		Precipation float64 `json:"1h"`
	} `json:"rain"`
}
